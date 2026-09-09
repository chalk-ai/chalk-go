package envfs

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"path/filepath"
	"strings"
	"time"
)

// MapEnvironmentGetter provides a mock implementation for testing
type MapEnvironmentGetter struct {
	EnvironmentGetter
	Env map[string]string
	Wd  string
	FS  map[string]MockFile
}

// MockFile represents a mocked file with content and optional file info
type MockFile struct {
	Content []byte
	Info    fs.FileInfo
}

type mockFileInfo struct {
	name    string
	size    int64
	mode    fs.FileMode
	modTime time.Time
	isDir   bool
}

func (m *mockFileInfo) Name() string       { return m.name }
func (m *mockFileInfo) Size() int64        { return m.size }
func (m *mockFileInfo) Mode() fs.FileMode  { return m.mode }
func (m *mockFileInfo) ModTime() time.Time { return m.modTime }
func (m *mockFileInfo) IsDir() bool        { return m.isDir }
func (m *mockFileInfo) Sys() any           { return nil }

func (m MapEnvironmentGetter) Getenv(key string) string {
	return m.Env[key]
}

func (m MapEnvironmentGetter) Getwd() (string, error) {
	return m.Wd, nil
}

func (m MapEnvironmentGetter) Abs(path string) (string, error) {
	return defaultAbs(&m, path)
}

func (m MapEnvironmentGetter) Stat(name string) (fs.FileInfo, error) {
	if m.FS == nil {
		return nil, fs.ErrNotExist
	}

	if file, ok := m.FS[name]; ok {
		if file.Info != nil {
			return file.Info, nil
		}
		// Create default file info
		return &mockFileInfo{
			name:    filepath.Base(name),
			size:    int64(len(file.Content)),
			mode:    0644,
			modTime: time.Now(),
			isDir:   false,
		}, nil
	}

	return nil, fs.ErrNotExist
}

func (m MapEnvironmentGetter) ReadFile(name string) ([]byte, error) {
	if m.FS == nil {
		return nil, fs.ErrNotExist
	}

	if file, ok := m.FS[name]; ok {
		return file.Content, nil
	}

	return nil, fs.ErrNotExist
}

func (m *MapEnvironmentGetter) Open(name string) (File, error) {
	if m.FS == nil {
		return nil, fs.ErrNotExist
	}
	if _, ok := m.FS[name]; !ok {
		return nil, fs.ErrNotExist
	}
	return &mapEnvironmentFile{
		environment: m,
		name:        name,
		readable:    true,
	}, nil
}

func (m *MapEnvironmentGetter) CreateTemp(dir string, pattern string) (WritableFile, error) {
	if filepath.Base(pattern) != pattern {
		return nil, fs.ErrInvalid
	}
	if m.FS == nil {
		m.FS = make(map[string]MockFile)
	}
	if dir == "" {
		dir = m.Getenv("TMPDIR")
		if dir == "" {
			dir = string(filepath.Separator) + "tmp"
		}
	}

	for index := 1; ; index++ {
		suffix := fmt.Sprintf("%06d", index)
		name := pattern + suffix
		if wildcard := strings.LastIndex(pattern, "*"); wildcard >= 0 {
			name = pattern[:wildcard] + suffix + pattern[wildcard+1:]
		}
		path := filepath.Join(dir, name)
		if _, exists := m.FS[path]; exists {
			continue
		}
		m.FS[path] = MockFile{
			Info: &mockFileInfo{
				name:    name,
				mode:    0o600,
				modTime: time.Now(),
			},
		}
		return &mapEnvironmentFile{
			environment: m,
			name:        path,
			readable:    true,
			writable:    true,
		}, nil
	}
}

func (m *MapEnvironmentGetter) Remove(name string) error {
	if m.FS == nil {
		return fs.ErrNotExist
	}
	if _, ok := m.FS[name]; !ok {
		return fs.ErrNotExist
	}
	delete(m.FS, name)
	return nil
}

type mapEnvironmentFile struct {
	environment *MapEnvironmentGetter
	name        string
	offset      int64
	readable    bool
	writable    bool
	closed      bool
}

func (f *mapEnvironmentFile) Read(data []byte) (int, error) {
	if f.closed {
		return 0, fs.ErrClosed
	}
	if !f.readable {
		return 0, fs.ErrPermission
	}
	file, ok := f.environment.FS[f.name]
	if !ok {
		return 0, fs.ErrNotExist
	}
	reader := bytes.NewReader(file.Content)
	if _, err := reader.Seek(f.offset, io.SeekStart); err != nil {
		return 0, err
	}
	read, err := reader.Read(data)
	f.offset += int64(read)
	return read, err
}

func (f *mapEnvironmentFile) ReadAt(data []byte, offset int64) (int, error) {
	if f.closed {
		return 0, fs.ErrClosed
	}
	if !f.readable {
		return 0, fs.ErrPermission
	}
	file, ok := f.environment.FS[f.name]
	if !ok {
		return 0, fs.ErrNotExist
	}
	return bytes.NewReader(file.Content).ReadAt(data, offset)
}

func (f *mapEnvironmentFile) Seek(offset int64, whence int) (int64, error) {
	if f.closed {
		return 0, fs.ErrClosed
	}
	file, ok := f.environment.FS[f.name]
	if !ok {
		return 0, fs.ErrNotExist
	}
	position := offset
	switch whence {
	case io.SeekStart:
	case io.SeekCurrent:
		position += f.offset
	case io.SeekEnd:
		position += int64(len(file.Content))
	default:
		return 0, fs.ErrInvalid
	}
	if position < 0 {
		return 0, fs.ErrInvalid
	}
	f.offset = position
	return position, nil
}

func (f *mapEnvironmentFile) Write(data []byte) (int, error) {
	if f.closed {
		return 0, fs.ErrClosed
	}
	if !f.writable {
		return 0, fs.ErrPermission
	}
	file, ok := f.environment.FS[f.name]
	if !ok {
		return 0, fs.ErrNotExist
	}
	end := f.offset + int64(len(data))
	if end < f.offset || end > int64(int(^uint(0)>>1)) {
		return 0, fs.ErrInvalid
	}
	content := append([]byte(nil), file.Content...)
	if end > int64(len(content)) {
		content = append(content, make([]byte, int(end)-len(content))...)
	}
	written := copy(content[int(f.offset):int(end)], data)
	f.offset += int64(written)
	mode := fs.FileMode(0o644)
	if file.Info != nil {
		mode = file.Info.Mode()
	}
	f.environment.FS[f.name] = MockFile{
		Content: content,
		Info: &mockFileInfo{
			name:    filepath.Base(f.name),
			size:    int64(len(content)),
			mode:    mode,
			modTime: time.Now(),
		},
	}
	return written, nil
}

func (f *mapEnvironmentFile) Close() error {
	if f.closed {
		return fs.ErrClosed
	}
	f.closed = true
	return nil
}

func (f *mapEnvironmentFile) Name() string {
	return f.name
}

func (f *mapEnvironmentFile) Stat() (fs.FileInfo, error) {
	if f.closed {
		return nil, fs.ErrClosed
	}
	return f.environment.Stat(f.name)
}

// MockEnvironmentOption is a functional option for configuring MapEnvironmentGetter
type MockEnvironmentOption func(*MapEnvironmentGetter)

// WithEnv sets environment variables for the mock
func WithEnv(env map[string]string) MockEnvironmentOption {
	return func(m *MapEnvironmentGetter) {
		m.Env = env
	}
}

// WithWorkingDirectory sets the working directory for the mock
func WithWorkingDirectory(wd string) MockEnvironmentOption {
	return func(m *MapEnvironmentGetter) {
		m.Wd = wd
	}
}

// WithFS provides a complete mock filesystem
func WithFS(fs map[string]MockFile) MockEnvironmentOption {
	return func(m *MapEnvironmentGetter) {
		m.FS = fs
	}
}

// WithFile adds a single file to the mock filesystem
func WithFile(path string, content []byte) MockEnvironmentOption {
	return func(m *MapEnvironmentGetter) {
		if m.FS == nil {
			m.FS = make(map[string]MockFile)
		}
		m.FS[path] = MockFile{Content: content}
	}
}

// WithFileString adds a single file to the mock filesystem with string content
func WithFileString(path string, content string) MockEnvironmentOption {
	return WithFile(path, []byte(content))
}

// WithFileInfo adds a file with custom FileInfo to the mock filesystem
func WithFileInfo(path string, content []byte, info fs.FileInfo) MockEnvironmentOption {
	return func(m *MapEnvironmentGetter) {
		if m.FS == nil {
			m.FS = make(map[string]MockFile)
		}
		m.FS[path] = MockFile{Content: content, Info: info}
	}
}

// NewMockEnvironmentGetter creates a new mock environment getter with options
func NewMockEnvironmentGetter(opts ...MockEnvironmentOption) *MapEnvironmentGetter {
	m := &MapEnvironmentGetter{
		Env: make(map[string]string),
		Wd:  "/",
		FS:  make(map[string]MockFile),
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

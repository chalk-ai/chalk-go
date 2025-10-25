package envfs

import (
	"io/fs"
	"path/filepath"
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
func (m *mockFileInfo) Sys() interface{}   { return nil }

func (m MapEnvironmentGetter) Getenv(key string) string {
	return m.Env[key]
}

func (m MapEnvironmentGetter) Getwd() (string, error) {
	return m.Wd, nil
}

func (m MapEnvironmentGetter) Abs(path string) (string, error) {
	return defaultAbs(m, path)
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
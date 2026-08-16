package envfs

import (
	"io"
	"io/fs"
)

// File is a readable filesystem handle returned by EnvironmentGetter. It
// supports streaming and section readers without coupling callers to *os.File.
type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
	Name() string
	Stat() (fs.FileInfo, error)
}

// WritableFile is a File that can also be written to, as returned by
// EnvironmentGetter.CreateTemp.
type WritableFile interface {
	File
	io.Writer
}

// EnvironmentGetter provides environment and filesystem access. This can help
// with testing by making otherwise impure functions pure.
type EnvironmentGetter interface {
	Getenv(key string) string
	Getwd() (string, error)
	Abs(path string) (string, error)
	Stat(name string) (fs.FileInfo, error)
	ReadFile(name string) ([]byte, error)
	Open(name string) (File, error)
	CreateTemp(dir string, pattern string) (WritableFile, error)
	Remove(name string) error
}

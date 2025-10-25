package envfs

import "io/fs"

// EnvironmentGetter is an interface for getting environment variables
// and the current working directory. This can help with testing by making
// otherwise impure functions pure.
type EnvironmentGetter interface {
	Getenv(key string) string
	Getwd() (string, error)
	Abs(path string) (string, error)
	Stat(name string) (fs.FileInfo, error)
	ReadFile(name string) ([]byte, error)
}
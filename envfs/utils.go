package envfs

import "path/filepath"

// defaultAbs is a helper function to convert relative paths to absolute paths
func defaultAbs(getter EnvironmentGetter, path string) (string, error) {
	wd, err := getter.Getwd()
	if err != nil {
		return "", err
	}
	if filepath.IsAbs(path) {
		return path, nil
	}
	return filepath.Join(wd, path), nil
}
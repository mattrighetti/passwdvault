package utils

import "os"

// FileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func FileExists(filepath string) bool {
	info, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

// FolderExists checks if a dir exists
func FolderExists(dirpath string) bool {
	info, err := os.Stat(dirpath)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}

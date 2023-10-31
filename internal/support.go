package internal

import (
	"path/filepath"
)

// Return absolute path of a relative path
func GetAbsPath (relativePath string) (string, error) {
    return filepath.Abs(relativePath)
}

// Return if a file is an absolute path
func IsAbsPath (path string) bool {
    return filepath.IsAbs(path)
}

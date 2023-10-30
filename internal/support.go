package internal

import (
	"path/filepath"
)

func GetAbsPath (relativePath string) (string, error) {
    return filepath.Abs(relativePath)
}

func IsAbsPath (path string) bool {
    return filepath.IsAbs(path)
}

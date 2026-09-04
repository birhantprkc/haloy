package helpers

import (
	"os"

	"github.com/haloydev/haloy/internal/constants"
)

// EnsureDir creates the directory and any necessary parents with secure permissions.
func EnsureDir(dirPath string) error {
	return os.MkdirAll(dirPath, constants.ModeDirPrivate)
}

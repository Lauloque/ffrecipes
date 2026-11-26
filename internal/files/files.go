package files

import (
	"fmt"
	"os"
)

// Require checks that path exists and is not a directory.
func Require(path, label string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("%s %q: %w", label, path, err)
	}
	if info.IsDir() {
		return fmt.Errorf("%s %q is a directory", label, path)
	}
	return nil
}

package files

import (
	"fmt"
	"os"
	"path/filepath"
)

// WriteToFile writes data to file. Creates all directories in path and file if something does not exist.
func WriteToFile(path string, content []byte) error {
	err := os.MkdirAll(filepath.Dir(path), 0744)
	if err != nil {
		return fmt.Errorf("failed to create dirs: %w", err)
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("cannot generate output file: %w", err)
	}
	defer f.Close()

	_, err = f.Write(content)
	if err != nil {
		return fmt.Errorf("cannot write content to file: %w", err)
	}

	return nil
}

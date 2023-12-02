package internal

import (
	"io/fs"
	"path/filepath"
)

func WriteTree() {
	filepath.Walk(".", func(path string, entry fs.FileInfo, err error) error {
		if entry.IsDir() && entry.Name() == GOGIT_DIR_NAME {
			return filepath.SkipDir
		}
		return err
	})
}

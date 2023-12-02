package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitializeRepository() {
	// check whether directory exist or not
	path, err := os.Getwd()
	if err != nil {
		println("Unable to get current directory.")
		return
	}
	fullPath := filepath.Join(path, GOGIT_DIR_NAME)
	if err := os.Mkdir(fullPath, os.ModePerm); err != nil {
		println(err.Error())
		return
	}
	for _, internal_dir := range INTERNAL_DIRS {
		if err := os.Mkdir(filepath.Join(fullPath, internal_dir), os.ModePerm); err != nil {
			println(err.Error())
			return
		}
	}
	fmt.Printf("Initialized %s in %s.\n", GOGIT_DIR_NAME, path)
}

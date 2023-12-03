package internal

import (
	"os"
	"path/filepath"
)

type node struct {
	name string
	oid  string
	tag  ObjTag
}

func writeTreeObject(path string) string {
	dir, err := os.Open(path)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	defer dir.Close()
	files, err := dir.ReadDir(-1)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	for _, file := range files {
		fullPath := filepath.Join(path, file.Name())
		if file.Name() == GOGIT_DIR_NAME || file.Name() == ".git" {
			continue
		}
		if file.IsDir() {
			writeTreeObject(fullPath)
		}
		println(fullPath)
	}
	return "Hello"
}
func WriteTree() {
	writeTreeObject(".")
	// filepath.Walk(".", func(path string, entry fs.FileInfo, err error) error {
	// 	if err != nil {
	// 		println(err)
	// 		return err
	// 	}
	// 	// entries := make([]node, 0)
	// 	if entry.IsDir() && (entry.Name() == GOGIT_DIR_NAME || entry.Name() == ".git") {
	// 		return filepath.SkipDir
	// 	}
	// 	println(createHashObject([]byte(entry.Name()), BLOB))
	// 	return nil
	// })
}

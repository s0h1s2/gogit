package internal

import (
	"fmt"
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
	entries := make([]node, 0, 4)
	for _, file := range files {
		fullPath := filepath.Join(path, file.Name())
		if file.Name() == GOGIT_DIR_NAME || file.Name() == ".git" {
			continue
		}
		if !file.IsDir() {
			// content, err := os.ReadFile(fullPath)
			if err != nil {
				println(err.Error())
				os.Exit(1)
			}
			oid := createHashObjectFile(fullPath, BLOB)
			entries = append(entries, node{oid: oid, name: file.Name(), tag: BLOB})
		} else {
			oid := writeTreeObject(fullPath)
			entries = append(entries, node{oid: oid, name: file.Name(), tag: TREE})
		}
	}
	tree := ""
	for _, entry := range entries {
		tree += fmt.Sprintf("%s %s %s\n", entry.tag, entry.oid, entry.name)
	}
	return createHashObject([]byte(tree), TREE)

}
func WriteTree() {
	writeTreeObject(".")
}

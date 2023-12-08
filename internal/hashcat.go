package internal

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func HashCat(hashcat *HashCatCmd, expectedKind ObjTag) string {
	// check for hash existance in file system.
	fullPath := path.Join(getRepoPath(), OBJECTS_DIR_NAME, hashcat.Hash)
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Hash '%s' doesn't exist.\n", hashcat.Hash)
		return ""
	}
	content, err := os.ReadFile(fullPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	contentResult := strings.Split(string(content), "\x00")
	objectType := contentResult[0]
	if objectType != string(expectedKind) {
		fmt.Printf("Expected '%s' object but got '%s' object\n", expectedKind, objectType)
		return ""
	}
	fileContent := contentResult[1]
	return fileContent
}

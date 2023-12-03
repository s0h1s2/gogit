package internal

import (
	"log"
	"os"
	"path"
)

const GOGIT_DIR_NAME = ".gogit"
const OBJECTS_DIR_NAME = "objects"

var INTERNAL_DIRS = []string{OBJECTS_DIR_NAME}

type ObjTag string

const (
	BLOB ObjTag = "blob"
	TREE        = "TREE"
)

func getRepoPath() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return path.Join(wd, GOGIT_DIR_NAME)
}
func isRepoExist() {

}

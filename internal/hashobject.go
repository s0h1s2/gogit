package internal

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"path"
)

func createHashObject(content []byte, tag ObjTag) string {
	h := sha1.New()
	content = append([]byte(tag+"\x00"), content...)
	h.Write(content)
	oid := hex.EncodeToString(h.Sum(nil))
	f, err := os.Create(path.Join(getRepoPath(), OBJECTS_DIR_NAME, oid))
	defer f.Close()
	if err != nil {
		panic(err.Error())
	}
	f.Write(content)
	return oid
}
func createHashObjectFile(filePath string, tag ObjTag) string {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "File '%s' doesn't exist\n", filePath)
		return ""
	}
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	oid := createHashObject(content, tag)
	return oid
}
func CreateBlobHashObject(hashobj *HashObjectCmd) {
	createHashObjectFile(hashobj.FilePath, BLOB)
}

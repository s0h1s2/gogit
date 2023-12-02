package internal

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path"
)

func createHashObject(hashObjectCmd *HashObjectCmd, kind ObjTag) {
	if _, err := os.Stat(hashObjectCmd.FilePath); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "File '%s' doesn't exist\n", hashObjectCmd.FilePath)
		return
	}
	content, err := os.ReadFile(hashObjectCmd.FilePath)
	if err != nil {
		panic(err)
	}
	h := sha1.New()
	content = append([]byte(kind+"\x00"), content...)
	h.Write(content)
	fileName := hex.EncodeToString(h.Sum(nil))
	f, err := os.Create(path.Join(getRepoPath(), OBJECTS_DIR_NAME, fileName))
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(string(content))
}
func CreateBlobHashObject(hashobj *HashObjectCmd) {
	createHashObject(hashobj, BLOB)
}

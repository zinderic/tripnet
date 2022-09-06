package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

type Intel struct {
	FileName string `json:"file_name"`
	FileHash string `json:"file_hash"`
}

func main() {
	files := []string{"/etc/hosts", "/etc/fstab"} // TODO make discovery for all files eventually
	fmt.Println(GetFilesWithHashes(files))
}

func GetFilesWithHashes(files []string) []Intel {
	var FileHashes []Intel
	hasher := sha256.New()
	for _, f := range files {
		s, err := os.ReadFile(string(f))
		hasher.Write(s)
		if err != nil {
			log.Fatal(err)
		}
		hash := hex.EncodeToString(hasher.Sum(nil))
		FileHashes = append(FileHashes, Intel{
			FileName: f,
			FileHash: hash,
		})
		hasher.Reset()
	}
	return FileHashes
}

package agent

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Intel struct {
	FileName string `json:"file_name"`
	FileHash string `json:"file_hash"`
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

func GetFilesFromDirectory(directory string) []string {

	var allFiles []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() {
			allFiles = append(allFiles, path)
		}
		return nil
	})

	if err != nil {
		log.Fatalln(err)
	}

	return allFiles
}

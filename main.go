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
	hash(files)
}

func hash(files []string) {
	hasher := sha256.New()
	for _, f := range files {
		fmt.Println("Reading file " + f)
		s, err := os.ReadFile(string(f))
		hasher.Write(s)
		if err != nil {
			log.Fatal(err)
		}
		os.Stdout.WriteString(hex.EncodeToString(hasher.Sum(nil)))
		hasher.Reset()
	}
}

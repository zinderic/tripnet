package main

import (
	"flag"
	"fmt"
	"tripnet/pkg/agent"
)

var (
	directory *string
)

func main() {
	directory = flag.String("directory", ".", "directory to watch")
	flag.Parse()

	fmt.Printf("Directory: %s", *directory)

	files := agent.GetFilesFromDirectory(*directory)

	collectedIntel := agent.GetFilesWithHashes(files)
	// TODO do something with those instead of just returning
	for _, i := range collectedIntel {
		fmt.Printf("%s: %s\n", i.FileName, i.FileHash)
	}
}

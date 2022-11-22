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
	// TODO send those to the server using gRPC
	for _, i := range collectedIntel {
		fmt.Printf("%s: %s\n", i.FileName, i.FileHash)
	}
}

package main

import (
	"fmt"
	"tripnet/pkg/agent"
)

func main() {
	files := []string{"/etc/hosts", "/etc/fstab"} // TODO make discovery for all files eventually
	fmt.Println(agent.GetFilesWithHashes(files))
}

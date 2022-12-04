package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/bufbuild/connect-go"
	"log"
	"net/http"
	tripservv1 "tripnet/gen/tripserv/v1"
	"tripnet/gen/tripserv/v1/tripservv1connect"
	"tripnet/pkg/agent"
)

var (
	directory      *string
	serverEndpoint string = "http://localhost:8080" // TODO pass this as argument
)

func main() {
	directory = flag.String("directory", ".", "directory to watch")
	flag.Parse()

	fmt.Printf("Directory: %s", *directory)
	fmt.Printf("Server: %s", serverEndpoint)

	client := tripservv1connect.NewTripnetServiceClient(
		http.DefaultClient,
		serverEndpoint,
	)

	files := agent.GetFilesFromDirectory(*directory)

	collectedIntel := agent.GetFilesWithHashes(files)
	for _, i := range collectedIntel {
		res, err := client.FileHash(
			context.Background(),
			connect.NewRequest(&tripservv1.FileHashRequest{
				FilePath: i.FilePath,
				FileHash: i.FileHash,
			}),
		)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(res.Msg)
	}
}

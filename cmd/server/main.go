package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	tripservv1 "tripnet/gen/pkg/tripserv/v1"
	"tripnet/gen/pkg/tripserv/v1/tripservv1connect"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type TripServer struct{}

func (s *TripServer) FileHash(
	ctx context.Context,
	req *connect.Request[tripservv1.FileHashRequest],
) (*connect.Response[tripservv1.FileHashResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&tripservv1.FileHashResponse{
		Status: "received hash " + req.Msg.FileHash,
	})
	res.Header().Set("Tripserv-Version", "v1")
	// TODO persist to storage the received file/hash
	return res, nil
}

func main() {
	tripserv := &TripServer{}
	mux := http.NewServeMux()
	path, handler := tripservv1connect.NewTripnetServiceHandler(tripserv)
	mux.Handle(path, handler)
	fmt.Println("starting server..")
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

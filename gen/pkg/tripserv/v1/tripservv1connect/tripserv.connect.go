// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: pkg/tripserv/v1/tripserv.proto

package tripservv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
	v1 "tripnet/gen/pkg/tripserv/v1"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// TripnetServiceName is the fully-qualified name of the TripnetService service.
	TripnetServiceName = "tripserv.v1.TripnetService"
)

// TripnetServiceClient is a client for the tripserv.v1.TripnetService service.
type TripnetServiceClient interface {
	FileHash(context.Context, *connect_go.Request[v1.FileHashRequest]) (*connect_go.Response[v1.FileHashResponse], error)
}

// NewTripnetServiceClient constructs a client for the tripserv.v1.TripnetService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTripnetServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TripnetServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &tripnetServiceClient{
		fileHash: connect_go.NewClient[v1.FileHashRequest, v1.FileHashResponse](
			httpClient,
			baseURL+"/tripserv.v1.TripnetService/FileHash",
			opts...,
		),
	}
}

// tripnetServiceClient implements TripnetServiceClient.
type tripnetServiceClient struct {
	fileHash *connect_go.Client[v1.FileHashRequest, v1.FileHashResponse]
}

// FileHash calls tripserv.v1.TripnetService.FileHash.
func (c *tripnetServiceClient) FileHash(ctx context.Context, req *connect_go.Request[v1.FileHashRequest]) (*connect_go.Response[v1.FileHashResponse], error) {
	return c.fileHash.CallUnary(ctx, req)
}

// TripnetServiceHandler is an implementation of the tripserv.v1.TripnetService service.
type TripnetServiceHandler interface {
	FileHash(context.Context, *connect_go.Request[v1.FileHashRequest]) (*connect_go.Response[v1.FileHashResponse], error)
}

// NewTripnetServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTripnetServiceHandler(svc TripnetServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/tripserv.v1.TripnetService/FileHash", connect_go.NewUnaryHandler(
		"/tripserv.v1.TripnetService/FileHash",
		svc.FileHash,
		opts...,
	))
	return "/tripserv.v1.TripnetService/", mux
}

// UnimplementedTripnetServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTripnetServiceHandler struct{}

func (UnimplementedTripnetServiceHandler) FileHash(context.Context, *connect_go.Request[v1.FileHashRequest]) (*connect_go.Response[v1.FileHashResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("tripserv.v1.TripnetService.FileHash is not implemented"))
}

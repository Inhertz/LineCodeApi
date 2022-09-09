package rpc

import (
	"LineCodeApi/internal/adapters/rpc/pb"
	"LineCodeApi/internal/application"
	"LineCodeApi/internal/core/models"
	"context"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

//protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     internal/adapters/rpc/pb/linecode_svc.proto

// Adapter implements the GRPCPort interface
type Adapter struct {
	api application.APIPort
	pb.UnimplementedLineCoderServer
}

// NewAdapter creates a new Adapter
func NewAdapter(api application.APIPort) *Adapter {
	return &Adapter{api: api}
}

// RunAsync runs the server with a wait group for concurrency
func (grpca Adapter) RunAsync(wg *sync.WaitGroup) {

	grpca.Run()

	wg.Done()
}

// Run registers the LineCoderServiceServer to a grpcServer and serves on
// the specified port
func (grpca Adapter) Run() {
	var err error

	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterLineCoderServer(grpcServer, grpca)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve gRPC server over port 9000: %v", err)
	}
}

// ManchesterEncode returns a completely encoded model to the grpc client
func (grpca Adapter) ManchesterEncode(ctx context.Context, req *pb.ManchesterEncoderIn) (*pb.ManchesterOut, error) {

	out := &pb.ManchesterOut{}
	manchester := models.Manchester{
		Decoded:           req.Decoded,
		DecodedPulseWidth: req.DecodedPulseWidth,
		Unit:              req.Unit,
	}

	err := grpca.api.GenerateEncodedManchester(&manchester)
	if err != nil {
		return out, status.Error(codes.Internal, "unexpected error")
	}

	out = &pb.ManchesterOut{
		Id:                int32(manchester.ID),
		Decoded:           manchester.Decoded,
		Encoded:           manchester.Encoded,
		DecodedPulseWidth: manchester.DecodedPulseWidth,
		EncodedPulseWidth: manchester.EncodedPulseWidth,
		Unit:              manchester.Unit,
	}

	return out, nil
}

// ManchesterDecode returns a completely decoded model to the grpc client
func (grpca Adapter) ManchesterDecode(ctx context.Context, req *pb.ManchesterDecoderIn) (*pb.ManchesterOut, error) {

	out := &pb.ManchesterOut{}
	manchester := &models.Manchester{
		Encoded:           req.Encoded,
		EncodedPulseWidth: req.EncodedPulseWidth,
		Unit:              req.Unit,
	}

	err := grpca.api.GenerateDecodedManchester(manchester)
	if err != nil {
		return out, status.Error(codes.Internal, "unexpected error")
	}

	out = &pb.ManchesterOut{
		Id:                int32(manchester.ID),
		Decoded:           manchester.Decoded,
		Encoded:           manchester.Encoded,
		DecodedPulseWidth: manchester.DecodedPulseWidth,
		EncodedPulseWidth: manchester.EncodedPulseWidth,
		Unit:              manchester.Unit,
	}

	return out, nil
}

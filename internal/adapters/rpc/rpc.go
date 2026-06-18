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
	api  application.APIPort
	port string
	pb.UnimplementedLineCoderServer
}

// NewAdapter creates a new Adapter listening on the given port
func NewAdapter(api application.APIPort, port string) *Adapter {
	return &Adapter{api: api, port: port}
}

// RunAsync runs the server with a wait group for concurrency
func (grpca Adapter) RunAsync(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	grpca.Run(ctx)
}

// Run registers the LineCoderServiceServer to a grpcServer and serves on
// the configured port. Run blocks until ctx is cancelled, then stops the
// server gracefully.
func (grpca Adapter) Run(ctx context.Context) {
	listen, err := net.Listen("tcp", ":"+grpca.port)
	if err != nil {
		log.Fatalf("failed to listen on port %s: %v", grpca.port, err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterLineCoderServer(grpcServer, grpca)

	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve gRPC server on port %s: %v", grpca.port, err)
		}
	}()
	log.Printf("gRPC server listening on port %s", grpca.port)

	<-ctx.Done()

	grpcServer.GracefulStop()
	log.Println("gRPC server stopped")
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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: internal/adapters/rpc/pb/linecode_svc.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LineCoderClient is the client API for LineCoder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LineCoderClient interface {
	ManchesterEncode(ctx context.Context, in *ManchesterEncoderIn, opts ...grpc.CallOption) (*ManchesterOut, error)
	ManchesterDecode(ctx context.Context, in *ManchesterDecoderIn, opts ...grpc.CallOption) (*ManchesterOut, error)
}

type lineCoderClient struct {
	cc grpc.ClientConnInterface
}

func NewLineCoderClient(cc grpc.ClientConnInterface) LineCoderClient {
	return &lineCoderClient{cc}
}

func (c *lineCoderClient) ManchesterEncode(ctx context.Context, in *ManchesterEncoderIn, opts ...grpc.CallOption) (*ManchesterOut, error) {
	out := new(ManchesterOut)
	err := c.cc.Invoke(ctx, "/pb.LineCoder/ManchesterEncode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lineCoderClient) ManchesterDecode(ctx context.Context, in *ManchesterDecoderIn, opts ...grpc.CallOption) (*ManchesterOut, error) {
	out := new(ManchesterOut)
	err := c.cc.Invoke(ctx, "/pb.LineCoder/ManchesterDecode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LineCoderServer is the server API for LineCoder service.
// All implementations must embed UnimplementedLineCoderServer
// for forward compatibility
type LineCoderServer interface {
	ManchesterEncode(context.Context, *ManchesterEncoderIn) (*ManchesterOut, error)
	ManchesterDecode(context.Context, *ManchesterDecoderIn) (*ManchesterOut, error)
	mustEmbedUnimplementedLineCoderServer()
}

// UnimplementedLineCoderServer must be embedded to have forward compatible implementations.
type UnimplementedLineCoderServer struct {
}

func (UnimplementedLineCoderServer) ManchesterEncode(context.Context, *ManchesterEncoderIn) (*ManchesterOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManchesterEncode not implemented")
}
func (UnimplementedLineCoderServer) ManchesterDecode(context.Context, *ManchesterDecoderIn) (*ManchesterOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ManchesterDecode not implemented")
}
func (UnimplementedLineCoderServer) mustEmbedUnimplementedLineCoderServer() {}

// UnsafeLineCoderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LineCoderServer will
// result in compilation errors.
type UnsafeLineCoderServer interface {
	mustEmbedUnimplementedLineCoderServer()
}

func RegisterLineCoderServer(s grpc.ServiceRegistrar, srv LineCoderServer) {
	s.RegisterService(&LineCoder_ServiceDesc, srv)
}

func _LineCoder_ManchesterEncode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManchesterEncoderIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LineCoderServer).ManchesterEncode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LineCoder/ManchesterEncode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LineCoderServer).ManchesterEncode(ctx, req.(*ManchesterEncoderIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _LineCoder_ManchesterDecode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ManchesterDecoderIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LineCoderServer).ManchesterDecode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LineCoder/ManchesterDecode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LineCoderServer).ManchesterDecode(ctx, req.(*ManchesterDecoderIn))
	}
	return interceptor(ctx, in, info, handler)
}

// LineCoder_ServiceDesc is the grpc.ServiceDesc for LineCoder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LineCoder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LineCoder",
	HandlerType: (*LineCoderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ManchesterEncode",
			Handler:    _LineCoder_ManchesterEncode_Handler,
		},
		{
			MethodName: "ManchesterDecode",
			Handler:    _LineCoder_ManchesterDecode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/adapters/rpc/pb/linecode_svc.proto",
}

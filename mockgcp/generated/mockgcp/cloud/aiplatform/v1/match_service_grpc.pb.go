// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: mockgcp/cloud/aiplatform/v1/match_service.proto

package aiplatformpb

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

// MatchServiceClient is the client API for MatchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MatchServiceClient interface {
	// Finds the nearest neighbors of each vector within the request.
	FindNeighbors(ctx context.Context, in *FindNeighborsRequest, opts ...grpc.CallOption) (*FindNeighborsResponse, error)
	// Reads the datapoints/vectors of the given IDs.
	// A maximum of 1000 datapoints can be retrieved in a batch.
	ReadIndexDatapoints(ctx context.Context, in *ReadIndexDatapointsRequest, opts ...grpc.CallOption) (*ReadIndexDatapointsResponse, error)
}

type matchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMatchServiceClient(cc grpc.ClientConnInterface) MatchServiceClient {
	return &matchServiceClient{cc}
}

func (c *matchServiceClient) FindNeighbors(ctx context.Context, in *FindNeighborsRequest, opts ...grpc.CallOption) (*FindNeighborsResponse, error) {
	out := new(FindNeighborsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1.MatchService/FindNeighbors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *matchServiceClient) ReadIndexDatapoints(ctx context.Context, in *ReadIndexDatapointsRequest, opts ...grpc.CallOption) (*ReadIndexDatapointsResponse, error) {
	out := new(ReadIndexDatapointsResponse)
	err := c.cc.Invoke(ctx, "/mockgcp.cloud.aiplatform.v1.MatchService/ReadIndexDatapoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MatchServiceServer is the server API for MatchService service.
// All implementations must embed UnimplementedMatchServiceServer
// for forward compatibility
type MatchServiceServer interface {
	// Finds the nearest neighbors of each vector within the request.
	FindNeighbors(context.Context, *FindNeighborsRequest) (*FindNeighborsResponse, error)
	// Reads the datapoints/vectors of the given IDs.
	// A maximum of 1000 datapoints can be retrieved in a batch.
	ReadIndexDatapoints(context.Context, *ReadIndexDatapointsRequest) (*ReadIndexDatapointsResponse, error)
	mustEmbedUnimplementedMatchServiceServer()
}

// UnimplementedMatchServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMatchServiceServer struct {
}

func (UnimplementedMatchServiceServer) FindNeighbors(context.Context, *FindNeighborsRequest) (*FindNeighborsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNeighbors not implemented")
}
func (UnimplementedMatchServiceServer) ReadIndexDatapoints(context.Context, *ReadIndexDatapointsRequest) (*ReadIndexDatapointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadIndexDatapoints not implemented")
}
func (UnimplementedMatchServiceServer) mustEmbedUnimplementedMatchServiceServer() {}

// UnsafeMatchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MatchServiceServer will
// result in compilation errors.
type UnsafeMatchServiceServer interface {
	mustEmbedUnimplementedMatchServiceServer()
}

func RegisterMatchServiceServer(s grpc.ServiceRegistrar, srv MatchServiceServer) {
	s.RegisterService(&MatchService_ServiceDesc, srv)
}

func _MatchService_FindNeighbors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindNeighborsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).FindNeighbors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1.MatchService/FindNeighbors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).FindNeighbors(ctx, req.(*FindNeighborsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MatchService_ReadIndexDatapoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadIndexDatapointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MatchServiceServer).ReadIndexDatapoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mockgcp.cloud.aiplatform.v1.MatchService/ReadIndexDatapoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MatchServiceServer).ReadIndexDatapoints(ctx, req.(*ReadIndexDatapointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MatchService_ServiceDesc is the grpc.ServiceDesc for MatchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MatchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mockgcp.cloud.aiplatform.v1.MatchService",
	HandlerType: (*MatchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindNeighbors",
			Handler:    _MatchService_FindNeighbors_Handler,
		},
		{
			MethodName: "ReadIndexDatapoints",
			Handler:    _MatchService_ReadIndexDatapoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mockgcp/cloud/aiplatform/v1/match_service.proto",
}

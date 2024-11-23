// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: rating/rating.proto

package rating

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RatingService_SaveRating_FullMethodName = "/rating.RatingService/SaveRating"
	RatingService_GetRatings_FullMethodName = "/rating.RatingService/GetRatings"
)

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	// RPC to save a rating for a specific book
	SaveRating(ctx context.Context, in *SaveRatingRequest, opts ...grpc.CallOption) (*SaveRatingResponse, error)
	// RPC to retrieve all ratings for a specific book
	GetRatings(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetRatingsResponse, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) SaveRating(ctx context.Context, in *SaveRatingRequest, opts ...grpc.CallOption) (*SaveRatingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveRatingResponse)
	err := c.cc.Invoke(ctx, RatingService_SaveRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingServiceClient) GetRatings(ctx context.Context, in *GetRatingsRequest, opts ...grpc.CallOption) (*GetRatingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRatingsResponse)
	err := c.cc.Invoke(ctx, RatingService_GetRatings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility.
type RatingServiceServer interface {
	// RPC to save a rating for a specific book
	SaveRating(context.Context, *SaveRatingRequest) (*SaveRatingResponse, error)
	// RPC to retrieve all ratings for a specific book
	GetRatings(context.Context, *GetRatingsRequest) (*GetRatingsResponse, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRatingServiceServer struct{}

func (UnimplementedRatingServiceServer) SaveRating(context.Context, *SaveRatingRequest) (*SaveRatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveRating not implemented")
}
func (UnimplementedRatingServiceServer) GetRatings(context.Context, *GetRatingsRequest) (*GetRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRatings not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}
func (UnimplementedRatingServiceServer) testEmbeddedByValue()                       {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	// If the following call pancis, it indicates UnimplementedRatingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_SaveRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).SaveRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_SaveRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).SaveRating(ctx, req.(*SaveRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingService_GetRatings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetRatings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingService_GetRatings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetRatings(ctx, req.(*GetRatingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rating.RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveRating",
			Handler:    _RatingService_SaveRating_Handler,
		},
		{
			MethodName: "GetRatings",
			Handler:    _RatingService_GetRatings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rating/rating.proto",
}
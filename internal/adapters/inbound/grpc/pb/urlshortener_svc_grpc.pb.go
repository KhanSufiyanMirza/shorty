// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: urlshortener_svc.proto

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

// UrlShortnerServiceClient is the client API for UrlShortnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlShortnerServiceClient interface {
	ShortenUrl(ctx context.Context, in *ShortnerRequest, opts ...grpc.CallOption) (*ShortnerResponse, error)
	ResolveUrl(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlResponse, error)
}

type urlShortnerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlShortnerServiceClient(cc grpc.ClientConnInterface) UrlShortnerServiceClient {
	return &urlShortnerServiceClient{cc}
}

func (c *urlShortnerServiceClient) ShortenUrl(ctx context.Context, in *ShortnerRequest, opts ...grpc.CallOption) (*ShortnerResponse, error) {
	out := new(ShortnerResponse)
	err := c.cc.Invoke(ctx, "/pb.UrlShortnerService/ShortenUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlShortnerServiceClient) ResolveUrl(ctx context.Context, in *UrlRequest, opts ...grpc.CallOption) (*UrlResponse, error) {
	out := new(UrlResponse)
	err := c.cc.Invoke(ctx, "/pb.UrlShortnerService/ResolveUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlShortnerServiceServer is the server API for UrlShortnerService service.
// All implementations must embed UnimplementedUrlShortnerServiceServer
// for forward compatibility
type UrlShortnerServiceServer interface {
	ShortenUrl(context.Context, *ShortnerRequest) (*ShortnerResponse, error)
	ResolveUrl(context.Context, *UrlRequest) (*UrlResponse, error)
	mustEmbedUnimplementedUrlShortnerServiceServer()
}

// UnimplementedUrlShortnerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUrlShortnerServiceServer struct {
}

func (UnimplementedUrlShortnerServiceServer) ShortenUrl(context.Context, *ShortnerRequest) (*ShortnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ShortenUrl not implemented")
}
func (UnimplementedUrlShortnerServiceServer) ResolveUrl(context.Context, *UrlRequest) (*UrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveUrl not implemented")
}
func (UnimplementedUrlShortnerServiceServer) mustEmbedUnimplementedUrlShortnerServiceServer() {}

// UnsafeUrlShortnerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlShortnerServiceServer will
// result in compilation errors.
type UnsafeUrlShortnerServiceServer interface {
	mustEmbedUnimplementedUrlShortnerServiceServer()
}

func RegisterUrlShortnerServiceServer(s grpc.ServiceRegistrar, srv UrlShortnerServiceServer) {
	s.RegisterService(&UrlShortnerService_ServiceDesc, srv)
}

func _UrlShortnerService_ShortenUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShortnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortnerServiceServer).ShortenUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UrlShortnerService/ShortenUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortnerServiceServer).ShortenUrl(ctx, req.(*ShortnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlShortnerService_ResolveUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortnerServiceServer).ResolveUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UrlShortnerService/ResolveUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortnerServiceServer).ResolveUrl(ctx, req.(*UrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UrlShortnerService_ServiceDesc is the grpc.ServiceDesc for UrlShortnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UrlShortnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UrlShortnerService",
	HandlerType: (*UrlShortnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShortenUrl",
			Handler:    _UrlShortnerService_ShortenUrl_Handler,
		},
		{
			MethodName: "ResolveUrl",
			Handler:    _UrlShortnerService_ResolveUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "urlshortener_svc.proto",
}

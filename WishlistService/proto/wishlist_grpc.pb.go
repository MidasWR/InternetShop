// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package wishlist

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

// WishlistClient is the client API for Wishlist service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WishlistClient interface {
	AddToWishlist(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Ok, error)
	DeleteFromWishlist(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Ok, error)
	GetWishlist(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
}

type wishlistClient struct {
	cc grpc.ClientConnInterface
}

func NewWishlistClient(cc grpc.ClientConnInterface) WishlistClient {
	return &wishlistClient{cc}
}

func (c *wishlistClient) AddToWishlist(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/wishlist.wishlist/AddToWishlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wishlistClient) DeleteFromWishlist(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Ok, error) {
	out := new(Ok)
	err := c.cc.Invoke(ctx, "/wishlist.wishlist/DeleteFromWishlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wishlistClient) GetWishlist(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/wishlist.wishlist/GetWishlist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WishlistServer is the server API for Wishlist service.
// All implementations must embed UnimplementedWishlistServer
// for forward compatibility
type WishlistServer interface {
	AddToWishlist(context.Context, *Id) (*Ok, error)
	DeleteFromWishlist(context.Context, *Id) (*Ok, error)
	GetWishlist(context.Context, *Empty) (*Response, error)
	mustEmbedUnimplementedWishlistServer()
}

// UnimplementedWishlistServer must be embedded to have forward compatible implementations.
type UnimplementedWishlistServer struct {
}

func (UnimplementedWishlistServer) AddToWishlist(context.Context, *Id) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToWishlist not implemented")
}
func (UnimplementedWishlistServer) DeleteFromWishlist(context.Context, *Id) (*Ok, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFromWishlist not implemented")
}
func (UnimplementedWishlistServer) GetWishlist(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWishlist not implemented")
}
func (UnimplementedWishlistServer) mustEmbedUnimplementedWishlistServer() {}

// UnsafeWishlistServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WishlistServer will
// result in compilation errors.
type UnsafeWishlistServer interface {
	mustEmbedUnimplementedWishlistServer()
}

func RegisterWishlistServer(s grpc.ServiceRegistrar, srv WishlistServer) {
	s.RegisterService(&Wishlist_ServiceDesc, srv)
}

func _Wishlist_AddToWishlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishlistServer).AddToWishlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wishlist.wishlist/AddToWishlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishlistServer).AddToWishlist(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wishlist_DeleteFromWishlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishlistServer).DeleteFromWishlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wishlist.wishlist/DeleteFromWishlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishlistServer).DeleteFromWishlist(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Wishlist_GetWishlist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WishlistServer).GetWishlist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wishlist.wishlist/GetWishlist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WishlistServer).GetWishlist(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Wishlist_ServiceDesc is the grpc.ServiceDesc for Wishlist service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Wishlist_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "wishlist.wishlist",
	HandlerType: (*WishlistServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddToWishlist",
			Handler:    _Wishlist_AddToWishlist_Handler,
		},
		{
			MethodName: "DeleteFromWishlist",
			Handler:    _Wishlist_DeleteFromWishlist_Handler,
		},
		{
			MethodName: "GetWishlist",
			Handler:    _Wishlist_GetWishlist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "WishlistService/proto/wishlist.proto",
}
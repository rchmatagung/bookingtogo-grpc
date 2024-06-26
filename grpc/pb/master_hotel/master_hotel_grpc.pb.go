// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0
// source: master_hotel/master_hotel.proto

package master_hotel

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

// MasterHotelClient is the client API for MasterHotel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterHotelClient interface {
	GetAutoComplete(ctx context.Context, in *AutoCompleteRequest, opts ...grpc.CallOption) (*AutoComplete, error)
}

type masterHotelClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterHotelClient(cc grpc.ClientConnInterface) MasterHotelClient {
	return &masterHotelClient{cc}
}

func (c *masterHotelClient) GetAutoComplete(ctx context.Context, in *AutoCompleteRequest, opts ...grpc.CallOption) (*AutoComplete, error) {
	out := new(AutoComplete)
	err := c.cc.Invoke(ctx, "/pos.library.master_hotel.MasterHotel/GetAutoComplete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterHotelServer is the server API for MasterHotel service.
// All implementations must embed UnimplementedMasterHotelServer
// for forward compatibility
type MasterHotelServer interface {
	GetAutoComplete(context.Context, *AutoCompleteRequest) (*AutoComplete, error)
	mustEmbedUnimplementedMasterHotelServer()
}

// UnimplementedMasterHotelServer must be embedded to have forward compatible implementations.
type UnimplementedMasterHotelServer struct {
}

func (UnimplementedMasterHotelServer) GetAutoComplete(context.Context, *AutoCompleteRequest) (*AutoComplete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAutoComplete not implemented")
}
func (UnimplementedMasterHotelServer) mustEmbedUnimplementedMasterHotelServer() {}

// UnsafeMasterHotelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterHotelServer will
// result in compilation errors.
type UnsafeMasterHotelServer interface {
	mustEmbedUnimplementedMasterHotelServer()
}

func RegisterMasterHotelServer(s grpc.ServiceRegistrar, srv MasterHotelServer) {
	s.RegisterService(&MasterHotel_ServiceDesc, srv)
}

func _MasterHotel_GetAutoComplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AutoCompleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterHotelServer).GetAutoComplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.library.master_hotel.MasterHotel/GetAutoComplete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterHotelServer).GetAutoComplete(ctx, req.(*AutoCompleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterHotel_ServiceDesc is the grpc.ServiceDesc for MasterHotel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterHotel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pos.library.master_hotel.MasterHotel",
	HandlerType: (*MasterHotelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAutoComplete",
			Handler:    _MasterHotel_GetAutoComplete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master_hotel/master_hotel.proto",
}

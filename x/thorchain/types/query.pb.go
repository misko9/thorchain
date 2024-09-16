// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("types/query.proto", fileDescriptor_3dd69a0a5972f391) }

var fileDescriptor_3dd69a0a5972f391 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xb1, 0x4e, 0x02, 0x41,
	0x10, 0x86, 0x41, 0xc5, 0x62, 0xb5, 0x90, 0xd1, 0x08, 0x5c, 0xcc, 0x42, 0x7c, 0x80, 0xdb, 0x44,
	0xdf, 0xc0, 0xd8, 0xd8, 0xa9, 0x85, 0x05, 0x8d, 0x59, 0xb8, 0xcd, 0x71, 0xc9, 0xb9, 0xb3, 0xdc,
	0x2c, 0x46, 0x62, 0x6c, 0x7c, 0x02, 0x13, 0x4b, 0x5f, 0xc8, 0x92, 0xc4, 0xc6, 0xd2, 0x80, 0x0f,
	0x62, 0x76, 0x39, 0xcc, 0x0a, 0xd2, 0x4d, 0xfe, 0x7f, 0xe7, 0xfb, 0x6e, 0x72, 0xac, 0x6e, 0xc7,
	0x46, 0x91, 0x18, 0x8e, 0x54, 0x31, 0x8e, 0x4d, 0x81, 0x16, 0xa1, 0xe6, 0xa3, 0xe8, 0x28, 0x45,
	0x4c, 0x73, 0x25, 0xa4, 0xc9, 0x84, 0xd4, 0x1a, 0xad, 0xb4, 0x19, 0x6a, 0x9a, 0x3f, 0x8a, 0x0e,
	0x52, 0x4c, 0xd1, 0x8f, 0xc2, 0x4d, 0x65, 0x7a, 0x18, 0xd0, 0x6e, 0x0d, 0x62, 0x5e, 0xe6, 0x3c,
	0xcc, 0x13, 0x55, 0x64, 0xf7, 0x2a, 0x09, 0xfa, 0x93, 0xb7, 0x4d, 0x56, 0xbb, 0x72, 0x25, 0x74,
	0xd9, 0xd6, 0x25, 0x62, 0x0e, 0x8d, 0xd8, 0xaf, 0xc4, 0x3e, 0x75, 0xc9, 0xb5, 0x1a, 0x8e, 0x14,
	0xd9, 0xa8, 0xb9, 0x5a, 0x90, 0x41, 0x4d, 0xea, 0xb8, 0xfd, 0xfc, 0xf1, 0xfd, 0xba, 0xd1, 0x82,
	0x86, 0xb0, 0x03, 0x2c, 0xfa, 0x03, 0x99, 0x69, 0xe1, 0x1c, 0xe2, 0x51, 0x12, 0x29, 0xfb, 0x04,
	0x37, 0xac, 0xe6, 0x16, 0x08, 0x56, 0x18, 0xb4, 0xa0, 0xb7, 0xfe, 0x69, 0x4a, 0x7c, 0xd3, 0xe3,
	0x01, 0xf6, 0x96, 0xf0, 0x04, 0x86, 0xed, 0x9c, 0xcf, 0x6f, 0xf2, 0x9f, 0xce, 0x43, 0x46, 0x50,
	0x2c, 0x1c, 0xed, 0xb5, 0x7d, 0x69, 0xea, 0x78, 0x53, 0x04, 0xcd, 0xc0, 0x94, 0xfc, 0xb9, 0x24,
	0x67, 0xbb, 0xc1, 0x22, 0xc1, 0x3a, 0xe4, 0xef, 0x5d, 0x9d, 0xf5, 0x0f, 0x4a, 0x69, 0xcb, 0x4b,
	0xf7, 0xa1, 0xbe, 0x2c, 0xa5, 0xb3, 0x8b, 0xf7, 0x29, 0xaf, 0x4e, 0xa6, 0xbc, 0xfa, 0x35, 0xe5,
	0xd5, 0x97, 0x19, 0xaf, 0x4c, 0x66, 0xbc, 0xf2, 0x39, 0xe3, 0x95, 0xae, 0x48, 0x33, 0x9b, 0xcb,
	0x5e, 0xdc, 0xc7, 0xbb, 0x60, 0xcd, 0x4d, 0x1a, 0x13, 0x25, 0x1e, 0xc2, 0xd0, 0xf9, 0x7b, 0xdb,
	0xfe, 0x7f, 0x9f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x67, 0xa7, 0xf5, 0x77, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Returns the extended pool information for the provided asset.
	Pool(ctx context.Context, in *QueryPoolRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error)
	// Pools returns all extended pools
	Pools(ctx context.Context, in *QueryPoolsRequest, opts ...grpc.CallOption) (*QueryPoolsResponse, error)
	DerivedPool(ctx context.Context, in *QueryDerivedPoolRequest, opts ...grpc.CallOption) (*QueryDerivedPoolResponse, error)
	DerivedPools(ctx context.Context, in *QueryDerivedPoolsRequest, opts ...grpc.CallOption) (*QueryDerivedPoolsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Pool(ctx context.Context, in *QueryPoolRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error) {
	out := new(QueryPoolResponse)
	err := c.cc.Invoke(ctx, "/types.Query/Pool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Pools(ctx context.Context, in *QueryPoolsRequest, opts ...grpc.CallOption) (*QueryPoolsResponse, error) {
	out := new(QueryPoolsResponse)
	err := c.cc.Invoke(ctx, "/types.Query/Pools", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DerivedPool(ctx context.Context, in *QueryDerivedPoolRequest, opts ...grpc.CallOption) (*QueryDerivedPoolResponse, error) {
	out := new(QueryDerivedPoolResponse)
	err := c.cc.Invoke(ctx, "/types.Query/DerivedPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DerivedPools(ctx context.Context, in *QueryDerivedPoolsRequest, opts ...grpc.CallOption) (*QueryDerivedPoolsResponse, error) {
	out := new(QueryDerivedPoolsResponse)
	err := c.cc.Invoke(ctx, "/types.Query/DerivedPools", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Returns the extended pool information for the provided asset.
	Pool(context.Context, *QueryPoolRequest) (*QueryPoolResponse, error)
	// Pools returns all extended pools
	Pools(context.Context, *QueryPoolsRequest) (*QueryPoolsResponse, error)
	DerivedPool(context.Context, *QueryDerivedPoolRequest) (*QueryDerivedPoolResponse, error)
	DerivedPools(context.Context, *QueryDerivedPoolsRequest) (*QueryDerivedPoolsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Pool(ctx context.Context, req *QueryPoolRequest) (*QueryPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pool not implemented")
}
func (*UnimplementedQueryServer) Pools(ctx context.Context, req *QueryPoolsRequest) (*QueryPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pools not implemented")
}
func (*UnimplementedQueryServer) DerivedPool(ctx context.Context, req *QueryDerivedPoolRequest) (*QueryDerivedPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DerivedPool not implemented")
}
func (*UnimplementedQueryServer) DerivedPools(ctx context.Context, req *QueryDerivedPoolsRequest) (*QueryDerivedPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DerivedPools not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Pool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Query/Pool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pool(ctx, req.(*QueryPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Pools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Pools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Query/Pools",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Pools(ctx, req.(*QueryPoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DerivedPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDerivedPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DerivedPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Query/DerivedPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DerivedPool(ctx, req.(*QueryDerivedPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DerivedPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDerivedPoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DerivedPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.Query/DerivedPools",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DerivedPools(ctx, req.(*QueryDerivedPoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Pool",
			Handler:    _Query_Pool_Handler,
		},
		{
			MethodName: "Pools",
			Handler:    _Query_Pools_Handler,
		},
		{
			MethodName: "DerivedPool",
			Handler:    _Query_DerivedPool_Handler,
		},
		{
			MethodName: "DerivedPools",
			Handler:    _Query_DerivedPools_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types/query.proto",
}

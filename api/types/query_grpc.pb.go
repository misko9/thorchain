// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: types/query.proto

package types

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

const (
	Query_Pool_FullMethodName               = "/types.Query/Pool"
	Query_Pools_FullMethodName              = "/types.Query/Pools"
	Query_DerivedPool_FullMethodName        = "/types.Query/DerivedPool"
	Query_DerivedPools_FullMethodName       = "/types.Query/DerivedPools"
	Query_LiquidityProvider_FullMethodName  = "/types.Query/LiquidityProvider"
	Query_LiquidityProviders_FullMethodName = "/types.Query/LiquidityProviders"
	Query_Saver_FullMethodName              = "/types.Query/Saver"
	Query_Savers_FullMethodName             = "/types.Query/Savers"
	Query_Borrower_FullMethodName           = "/types.Query/Borrower"
	Query_Borrowers_FullMethodName          = "/types.Query/Borrowers"
	Query_TradeUnit_FullMethodName          = "/types.Query/TradeUnit"
	Query_TradeUnits_FullMethodName         = "/types.Query/TradeUnits"
	Query_TradeAccount_FullMethodName       = "/types.Query/TradeAccount"
	Query_TradeAccounts_FullMethodName      = "/types.Query/TradeAccounts"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Returns the extended pool information for the provided asset.
	Pool(ctx context.Context, in *QueryPoolRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error)
	// Pools returns all extended pools
	Pools(ctx context.Context, in *QueryPoolsRequest, opts ...grpc.CallOption) (*QueryPoolsResponse, error)
	DerivedPool(ctx context.Context, in *QueryDerivedPoolRequest, opts ...grpc.CallOption) (*QueryDerivedPoolResponse, error)
	DerivedPools(ctx context.Context, in *QueryDerivedPoolsRequest, opts ...grpc.CallOption) (*QueryDerivedPoolsResponse, error)
	LiquidityProvider(ctx context.Context, in *QueryLiquidityProviderRequest, opts ...grpc.CallOption) (*QueryLiquidityProviderResponse, error)
	LiquidityProviders(ctx context.Context, in *QueryLiquidityProvidersRequest, opts ...grpc.CallOption) (*QueryLiquidityProvidersResponse, error)
	Saver(ctx context.Context, in *QuerySaverRequest, opts ...grpc.CallOption) (*QuerySaverResponse, error)
	Savers(ctx context.Context, in *QuerySaversRequest, opts ...grpc.CallOption) (*QuerySaversResponse, error)
	Borrower(ctx context.Context, in *QueryBorrowerRequest, opts ...grpc.CallOption) (*QueryBorrowerResponse, error)
	Borrowers(ctx context.Context, in *QueryBorrowersRequest, opts ...grpc.CallOption) (*QueryBorrowersResponse, error)
	TradeUnit(ctx context.Context, in *QueryTradeUnitRequest, opts ...grpc.CallOption) (*QueryTradeUnitResponse, error)
	TradeUnits(ctx context.Context, in *QueryTradeUnitsRequest, opts ...grpc.CallOption) (*QueryTradeUnitsResponse, error)
	TradeAccount(ctx context.Context, in *QueryTradeAccountRequest, opts ...grpc.CallOption) (*QueryTradeAccountsResponse, error)
	TradeAccounts(ctx context.Context, in *QueryTradeAccountsRequest, opts ...grpc.CallOption) (*QueryTradeAccountsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Pool(ctx context.Context, in *QueryPoolRequest, opts ...grpc.CallOption) (*QueryPoolResponse, error) {
	out := new(QueryPoolResponse)
	err := c.cc.Invoke(ctx, Query_Pool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Pools(ctx context.Context, in *QueryPoolsRequest, opts ...grpc.CallOption) (*QueryPoolsResponse, error) {
	out := new(QueryPoolsResponse)
	err := c.cc.Invoke(ctx, Query_Pools_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DerivedPool(ctx context.Context, in *QueryDerivedPoolRequest, opts ...grpc.CallOption) (*QueryDerivedPoolResponse, error) {
	out := new(QueryDerivedPoolResponse)
	err := c.cc.Invoke(ctx, Query_DerivedPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DerivedPools(ctx context.Context, in *QueryDerivedPoolsRequest, opts ...grpc.CallOption) (*QueryDerivedPoolsResponse, error) {
	out := new(QueryDerivedPoolsResponse)
	err := c.cc.Invoke(ctx, Query_DerivedPools_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LiquidityProvider(ctx context.Context, in *QueryLiquidityProviderRequest, opts ...grpc.CallOption) (*QueryLiquidityProviderResponse, error) {
	out := new(QueryLiquidityProviderResponse)
	err := c.cc.Invoke(ctx, Query_LiquidityProvider_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) LiquidityProviders(ctx context.Context, in *QueryLiquidityProvidersRequest, opts ...grpc.CallOption) (*QueryLiquidityProvidersResponse, error) {
	out := new(QueryLiquidityProvidersResponse)
	err := c.cc.Invoke(ctx, Query_LiquidityProviders_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Saver(ctx context.Context, in *QuerySaverRequest, opts ...grpc.CallOption) (*QuerySaverResponse, error) {
	out := new(QuerySaverResponse)
	err := c.cc.Invoke(ctx, Query_Saver_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Savers(ctx context.Context, in *QuerySaversRequest, opts ...grpc.CallOption) (*QuerySaversResponse, error) {
	out := new(QuerySaversResponse)
	err := c.cc.Invoke(ctx, Query_Savers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Borrower(ctx context.Context, in *QueryBorrowerRequest, opts ...grpc.CallOption) (*QueryBorrowerResponse, error) {
	out := new(QueryBorrowerResponse)
	err := c.cc.Invoke(ctx, Query_Borrower_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Borrowers(ctx context.Context, in *QueryBorrowersRequest, opts ...grpc.CallOption) (*QueryBorrowersResponse, error) {
	out := new(QueryBorrowersResponse)
	err := c.cc.Invoke(ctx, Query_Borrowers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TradeUnit(ctx context.Context, in *QueryTradeUnitRequest, opts ...grpc.CallOption) (*QueryTradeUnitResponse, error) {
	out := new(QueryTradeUnitResponse)
	err := c.cc.Invoke(ctx, Query_TradeUnit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TradeUnits(ctx context.Context, in *QueryTradeUnitsRequest, opts ...grpc.CallOption) (*QueryTradeUnitsResponse, error) {
	out := new(QueryTradeUnitsResponse)
	err := c.cc.Invoke(ctx, Query_TradeUnits_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TradeAccount(ctx context.Context, in *QueryTradeAccountRequest, opts ...grpc.CallOption) (*QueryTradeAccountsResponse, error) {
	out := new(QueryTradeAccountsResponse)
	err := c.cc.Invoke(ctx, Query_TradeAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TradeAccounts(ctx context.Context, in *QueryTradeAccountsRequest, opts ...grpc.CallOption) (*QueryTradeAccountsResponse, error) {
	out := new(QueryTradeAccountsResponse)
	err := c.cc.Invoke(ctx, Query_TradeAccounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Returns the extended pool information for the provided asset.
	Pool(context.Context, *QueryPoolRequest) (*QueryPoolResponse, error)
	// Pools returns all extended pools
	Pools(context.Context, *QueryPoolsRequest) (*QueryPoolsResponse, error)
	DerivedPool(context.Context, *QueryDerivedPoolRequest) (*QueryDerivedPoolResponse, error)
	DerivedPools(context.Context, *QueryDerivedPoolsRequest) (*QueryDerivedPoolsResponse, error)
	LiquidityProvider(context.Context, *QueryLiquidityProviderRequest) (*QueryLiquidityProviderResponse, error)
	LiquidityProviders(context.Context, *QueryLiquidityProvidersRequest) (*QueryLiquidityProvidersResponse, error)
	Saver(context.Context, *QuerySaverRequest) (*QuerySaverResponse, error)
	Savers(context.Context, *QuerySaversRequest) (*QuerySaversResponse, error)
	Borrower(context.Context, *QueryBorrowerRequest) (*QueryBorrowerResponse, error)
	Borrowers(context.Context, *QueryBorrowersRequest) (*QueryBorrowersResponse, error)
	TradeUnit(context.Context, *QueryTradeUnitRequest) (*QueryTradeUnitResponse, error)
	TradeUnits(context.Context, *QueryTradeUnitsRequest) (*QueryTradeUnitsResponse, error)
	TradeAccount(context.Context, *QueryTradeAccountRequest) (*QueryTradeAccountsResponse, error)
	TradeAccounts(context.Context, *QueryTradeAccountsRequest) (*QueryTradeAccountsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Pool(context.Context, *QueryPoolRequest) (*QueryPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pool not implemented")
}
func (UnimplementedQueryServer) Pools(context.Context, *QueryPoolsRequest) (*QueryPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pools not implemented")
}
func (UnimplementedQueryServer) DerivedPool(context.Context, *QueryDerivedPoolRequest) (*QueryDerivedPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DerivedPool not implemented")
}
func (UnimplementedQueryServer) DerivedPools(context.Context, *QueryDerivedPoolsRequest) (*QueryDerivedPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DerivedPools not implemented")
}
func (UnimplementedQueryServer) LiquidityProvider(context.Context, *QueryLiquidityProviderRequest) (*QueryLiquidityProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidityProvider not implemented")
}
func (UnimplementedQueryServer) LiquidityProviders(context.Context, *QueryLiquidityProvidersRequest) (*QueryLiquidityProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidityProviders not implemented")
}
func (UnimplementedQueryServer) Saver(context.Context, *QuerySaverRequest) (*QuerySaverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Saver not implemented")
}
func (UnimplementedQueryServer) Savers(context.Context, *QuerySaversRequest) (*QuerySaversResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Savers not implemented")
}
func (UnimplementedQueryServer) Borrower(context.Context, *QueryBorrowerRequest) (*QueryBorrowerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Borrower not implemented")
}
func (UnimplementedQueryServer) Borrowers(context.Context, *QueryBorrowersRequest) (*QueryBorrowersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Borrowers not implemented")
}
func (UnimplementedQueryServer) TradeUnit(context.Context, *QueryTradeUnitRequest) (*QueryTradeUnitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeUnit not implemented")
}
func (UnimplementedQueryServer) TradeUnits(context.Context, *QueryTradeUnitsRequest) (*QueryTradeUnitsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeUnits not implemented")
}
func (UnimplementedQueryServer) TradeAccount(context.Context, *QueryTradeAccountRequest) (*QueryTradeAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeAccount not implemented")
}
func (UnimplementedQueryServer) TradeAccounts(context.Context, *QueryTradeAccountsRequest) (*QueryTradeAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TradeAccounts not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
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
		FullMethod: Query_Pool_FullMethodName,
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
		FullMethod: Query_Pools_FullMethodName,
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
		FullMethod: Query_DerivedPool_FullMethodName,
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
		FullMethod: Query_DerivedPools_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DerivedPools(ctx, req.(*QueryDerivedPoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LiquidityProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLiquidityProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LiquidityProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_LiquidityProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LiquidityProvider(ctx, req.(*QueryLiquidityProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_LiquidityProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryLiquidityProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).LiquidityProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_LiquidityProviders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).LiquidityProviders(ctx, req.(*QueryLiquidityProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Saver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySaverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Saver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Saver_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Saver(ctx, req.(*QuerySaverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Savers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySaversRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Savers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Savers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Savers(ctx, req.(*QuerySaversRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Borrower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBorrowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Borrower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Borrower_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Borrower(ctx, req.(*QueryBorrowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Borrowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBorrowersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Borrowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Borrowers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Borrowers(ctx, req.(*QueryBorrowersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TradeUnit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTradeUnitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TradeUnit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TradeUnit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TradeUnit(ctx, req.(*QueryTradeUnitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TradeUnits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTradeUnitsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TradeUnits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TradeUnits_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TradeUnits(ctx, req.(*QueryTradeUnitsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TradeAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTradeAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TradeAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TradeAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TradeAccount(ctx, req.(*QueryTradeAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TradeAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTradeAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TradeAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_TradeAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TradeAccounts(ctx, req.(*QueryTradeAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
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
		{
			MethodName: "LiquidityProvider",
			Handler:    _Query_LiquidityProvider_Handler,
		},
		{
			MethodName: "LiquidityProviders",
			Handler:    _Query_LiquidityProviders_Handler,
		},
		{
			MethodName: "Saver",
			Handler:    _Query_Saver_Handler,
		},
		{
			MethodName: "Savers",
			Handler:    _Query_Savers_Handler,
		},
		{
			MethodName: "Borrower",
			Handler:    _Query_Borrower_Handler,
		},
		{
			MethodName: "Borrowers",
			Handler:    _Query_Borrowers_Handler,
		},
		{
			MethodName: "TradeUnit",
			Handler:    _Query_TradeUnit_Handler,
		},
		{
			MethodName: "TradeUnits",
			Handler:    _Query_TradeUnits_Handler,
		},
		{
			MethodName: "TradeAccount",
			Handler:    _Query_TradeAccount_Handler,
		},
		{
			MethodName: "TradeAccounts",
			Handler:    _Query_TradeAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types/query.proto",
}

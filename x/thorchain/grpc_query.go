package thorchain

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"gitlab.com/thorchain/thornode/x/thorchain/types"
)

type queryServer struct {
	mgr Manager
}

var _ types.QueryServer = queryServer{}

func NewQueryServerImpl(mgr Manager) types.QueryServer {
	return queryServer{mgr: mgr}
}

func (s queryServer) Pool(c context.Context, req *types.QueryPoolRequest) (*types.QueryPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryPool(ctx, req)
}

func (s queryServer) Pools(c context.Context, req *types.QueryPoolsRequest) (*types.QueryPoolsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryPools(ctx, req)
}

func (s queryServer) DerivedPool(c context.Context, req *types.QueryDerivedPoolRequest) (*types.QueryDerivedPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryDerivedPool(ctx, req)
}

func (s queryServer)DerivedPools(c context.Context, req *types.QueryDerivedPoolsRequest) (*types.QueryDerivedPoolsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	return s.queryDerivedPools(ctx, req)
}

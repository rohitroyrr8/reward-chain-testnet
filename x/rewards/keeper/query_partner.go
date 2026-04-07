package keeper

import (
	"context"

	"cosmossdk.io/collections"
	"github.com/cosmos/cosmos-sdk/types/query"
	"rewardchain/x/rewards/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (q queryServer) ListPartner(ctx context.Context, req *types.QueryAllPartnerRequest) (*types.QueryAllPartnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	partners, pageRes, err := query.CollectionPaginate(
		ctx,
		q.k.Partner,
		req.Pagination,
		func(_ string, value types.Partner) (types.Partner, error){
			return value, nil
		},
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPartnerResponse{Partner: partners, Pagination: pageRes}, nil
}

func (q queryServer) GetPartner(ctx context.Context, req *types.QueryGetPartnerRequest) (*types.QueryGetPartnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	val, err := q.k.Partner.Get(ctx, req.Index)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}

		return nil, status.Error(codes.Internal, "internal error")
	}

	return &types.QueryGetPartnerResponse{Partner: val}, nil
}
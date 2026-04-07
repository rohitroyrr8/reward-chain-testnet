package keeper_test

import (
	"context"
	"testing"

    
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"rewardchain/x/rewards/keeper"
	"rewardchain/x/rewards/types"
)

func createNPartner(keeper keeper.Keeper, ctx context.Context, n int) []types.Partner {
	items := make([]types.Partner, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)
		items[i].Name = strconv.Itoa(i)
		items[i].Website = strconv.Itoa(i)
		items[i].Description = strconv.Itoa(i)
		items[i].Wallet = strconv.Itoa(i)
		items[i].Status = strconv.Itoa(i)
		_ = keeper.Partner.Set(ctx, items[i].Index, items[i])
	}
	return items
}

func TestPartnerQuerySingle(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNPartner(f.keeper, f.ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetPartnerRequest
		response *types.QueryGetPartnerResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetPartnerRequest{
			    Index: msgs[0].Index,
			},
			response: &types.QueryGetPartnerResponse{Partner: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetPartnerRequest{
			    Index: msgs[1].Index,
			},
			response: &types.QueryGetPartnerResponse{Partner: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetPartnerRequest{
				Index: strconv.Itoa(100000),
			},
			err:     status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := qs.GetPartner(f.ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.EqualExportedValues(t, tc.response, response)
			}
		})
	}
}

func TestPartnerQueryPaginated(t *testing.T) {
	f := initFixture(t)
	qs := keeper.NewQueryServerImpl(f.keeper)
	msgs := createNPartner(f.keeper, f.ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPartnerRequest {
		return &types.QueryAllPartnerRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListPartner(f.ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Partner), step)
			require.Subset(t, msgs, resp.Partner)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := qs.ListPartner(f.ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Partner), step)
			require.Subset(t, msgs, resp.Partner)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := qs.ListPartner(f.ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.EqualExportedValues(t, msgs, resp.Partner)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := qs.ListPartner(f.ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}

package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/greenbahar/vot/testutil/keeper"
	"github.com/greenbahar/vot/testutil/nullify"
	"github.com/greenbahar/vot/x/vot/types"
)

func TestNextVoteQuery(t *testing.T) {
	keeper, ctx := keepertest.VotKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestNextVote(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetNextVoteRequest
		response *types.QueryGetNextVoteResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetNextVoteRequest{},
			response: &types.QueryGetNextVoteResponse{NextVote: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.NextVote(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

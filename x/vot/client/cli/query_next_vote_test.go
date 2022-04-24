package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"github.com/greenbahar/vot/testutil/network"
	"github.com/greenbahar/vot/testutil/nullify"
	"github.com/greenbahar/vot/x/vot/client/cli"
	"github.com/greenbahar/vot/x/vot/types"
)

func networkWithNextVoteObjects(t *testing.T) (*network.Network, types.NextVote) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	nextVote := &types.NextVote{}
	nullify.Fill(&nextVote)
	state.NextVote = nextVote
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.NextVote
}

func TestShowNextVote(t *testing.T) {
	net, obj := networkWithNextVoteObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.NextVote
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowNextVote(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetNextVoteResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.NextVote)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.NextVote),
				)
			}
		})
	}
}

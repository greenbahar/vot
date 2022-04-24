package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/greenbahar/vot/x/vot/types"
	"github.com/spf13/cobra"
)

func CmdShowNextVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-next-vote",
		Short: "shows nextVote",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetNextVoteRequest{}

			res, err := queryClient.NextVote(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

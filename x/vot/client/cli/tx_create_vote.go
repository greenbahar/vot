package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/greenbahar/vot/x/vot/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateVote() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-vote [question] [options]",
		Short: "Broadcast message createVote",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argQuestion := args[0]
			argOptions := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateVote(
				clientCtx.GetFromAddress().String(),
				argQuestion,
				argOptions,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

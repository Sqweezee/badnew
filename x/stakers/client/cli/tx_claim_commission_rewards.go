package cli

import (
	"github.com/KYVENetwork/chain/x/stakers/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

func CmdClaimCommissionRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-commission-rewards [amount]",
		Short: "Broadcast message claim-commission-rewards",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argAmount, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			msg := types.MsgClaimCommissionRewards{
				Creator: clientCtx.GetFromAddress().String(),
				Amount:  argAmount,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

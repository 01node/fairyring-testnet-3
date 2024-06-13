package cli

import (
	"strconv"

	"github.com/Fairblock/fairyring/x/pep/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdSubmitEncryptedTx() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-encrypted-tx [data] [target-block-height]",
		Short: "Submit an encrypted transaction along with its execution height (execution height refers to the height in the FairyRing chain)",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argData := args[0]
			argTargetBlockHeight, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitEncryptedTx(
				clientCtx.GetFromAddress().String(),
				argData,
				argTargetBlockHeight,
			)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

package cmd

import (
	"fmt"
	"log/slog"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/mavryk-network/mavpay/common"
	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mvgo/codec"
	"github.com/mavryk-network/mvgo/mavryk"
	"github.com/mavryk-network/mvgo/rpc"
	"github.com/spf13/cobra"
)

var transferCmd = &cobra.Command{
	Use:   "transfer <destination> <amount mav>",
	Short: "transfers mav to specified address",
	Long:  "transfers mav to specified address from payout wallet",
	Run: func(cmd *cobra.Command, args []string) {
		_, _, signer, transactor := assertRunWithResult(loadConfigurationEnginesExtensions, EXIT_CONFIGURATION_LOAD_FAILURE).Unwrap()
		mumav, _ := cmd.Flags().GetBool(MUMAV_FLAG)

		if len(args)%2 != 0 {
			slog.Error("invalid number of arguments (expects pairs of destination and amount)")
			os.Exit(EXIT_IVNALID_ARGS)
		}
		total := int64(0)

		destinations := make([]string, 0)

		op := codec.NewOp().WithSource(signer.GetPKH())
		op.WithTTL(constants.MAX_OPERATION_TTL)
		for i := 0; i < len(args); i += 2 {
			destination, err := mavryk.ParseAddress(args[i])
			if err != nil {
				slog.Error("invalid destination address", "address", args[i], "error", err.Error())
				os.Exit(EXIT_IVNALID_ARGS)
			}

			amount, err := strconv.ParseFloat(args[i+1], 64)
			if err != nil {
				slog.Error("invalid amount", "amount", args[i+1], "error", err.Error())
				os.Exit(EXIT_IVNALID_ARGS)
			}
			if !mumav {
				amount *= constants.MUMAV_FACTOR
			}

			mumav := int64(math.Floor(amount))
			total += mumav
			destinations = append(destinations, destination.String())
			op.WithTransfer(destination, mumav)
		}

		if err := requireConfirmation(fmt.Sprintf("do you really want to transfer %s to %s", common.MumavToMavS(total), strings.Join(destinations, ", "))); err != nil {
			os.Exit(EXIT_OPERTION_CANCELED)
		}
		slog.Info("transferring mav", "total", common.MumavToMavS(total), "destinations", strings.Join(destinations, ", "), "confirmations_required", constants.DEFAULT_REQUIRED_CONFIRMATIONS)
		opts := rpc.DefaultOptions
		opts.Confirmations = constants.DEFAULT_REQUIRED_CONFIRMATIONS
		opts.Signer = signer.GetSigner()

		rcpt, err := transactor.Send(op, &opts)
		if err != nil {
			slog.Error("failed to confirm tx", "error", err.Error())
			os.Exit(EXIT_OPERTION_FAILED)
		}
		if !rcpt.IsSuccess() {
			slog.Error("tx failed", "error", rcpt.Error().Error())
			os.Exit(EXIT_OPERTION_FAILED)
		}
		slog.Info("transfer successful")
	},
}

func init() {
	transferCmd.Flags().Bool(MUMAV_FLAG, false, "amount in mumav")
	transferCmd.Flags().Bool(CONFIRM_FLAG, false, "automatically confirms transfer")
	RootCmd.AddCommand(transferCmd)
}

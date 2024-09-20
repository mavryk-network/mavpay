package cmd

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/mavryk-network/mavpay/configuration/seed"
	"github.com/mavryk-network/mavpay/constants"
	"github.com/mavryk-network/mavpay/constants/enums"
	"github.com/mavryk-network/mavpay/state"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slices"
)

var generateConfigurationCmd = &cobra.Command{
	Use:     "import-configuration <kind> <source-file>",
	Short:   "seed configuration from",
	Aliases: []string{"import-config"},
	Long: `Generates configuration based on configuration from others payout distribution tools.

	Currently supported sources are: ` + strings.Join(lo.Map(enums.SUPPORTED_CONFIGURATION_SEED_KINDS, func(item enums.EConfigurationSeedKind, _ int) string {
		return string(item)
	}), ", ") + `

	To import configuration from supported sources copy configuration file to directory where you plan to store mavpay configuration and run command with source file path as argument.

	Example:
		mavpay import-configuration bc ./bc.json
		mavpay import-configuration trd ./trd.yaml
`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(2)(cmd, args); err != nil {
			return err
		}

		seedKind := enums.EConfigurationSeedKind(args[0])
		if !slices.Contains(enums.SUPPORTED_CONFIGURATION_SEED_KINDS, seedKind) {
			return errors.Join(constants.ErrInvalidConfigurationImportSource, fmt.Errorf("invalid seed: %s", seedKind))
		}
		if _, err := os.Stat(args[1]); err != nil {
			return errors.Join(constants.ErrInvalidConfigurationImportSource, fmt.Errorf("invalid source: %s", args[1]))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		sourceFile := args[1]
		destiantionFile := state.Global.GetConfigurationFilePath()

		if _, err := os.Stat(destiantionFile); err == nil {
			assertRequireConfirmation("configuration file already exists, overwrite?")
		}

		// load source bytes
		sourceBytes := assertRunWithResultAndErrorMessage(func() ([]byte, error) {
			return os.ReadFile(sourceFile)
		}, EXIT_CONFIGURATION_LOAD_FAILURE, "failed to read source file - %s")

		seededBytes, err := seed.Generate(sourceBytes, enums.EConfigurationSeedKind(args[0]))
		if err != nil {
			slog.Error("failed to generate configuration", "error", err.Error())
			os.Exit(EXIT_CONFIGURATION_GENERATE_FAILURE)
		}
		assertRunWithErrorMessage(func() error {
			if target, err := os.Stat(destiantionFile); err == nil {
				if source, err := os.Stat(sourceFile); err == nil {
					if os.SameFile(target, source) {
						// backup old configuration file
						return os.Rename(destiantionFile, destiantionFile+constants.CONFIG_FILE_BACKUP_SUFFIX)
					}
				}
			}
			return os.WriteFile(destiantionFile, seededBytes, 0644)
		}, EXIT_CONFIGURATION_SAVE_FAILURE, "failed to save configuration file - %s")
		slog.Info("configuration imported successfully")
	},
}

func init() {
	RootCmd.AddCommand(generateConfigurationCmd)
}

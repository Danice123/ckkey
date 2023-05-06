package cmd

import (
	"errors"
	"strconv"

	"github.com/Danice123/ckkey/internal/roll"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate [trainer id] [dex number] [level]",
	Short: "Generate a deterministic encounter",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		trainerId, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if trainerId < 0 || trainerId > 65535 {
			return errors.New("trainer ID invalid")
		}
		dexNumber, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		if dexNumber < 1 || dexNumber > 251 {
			return errors.New("dex ID invalid")
		}
		level, err := strconv.Atoi(args[2])
		if err != nil {
			return err
		}
		if level < 1 {
			return errors.New("level invalid")
		}

		e := roll.CalcDVs(trainerId, dexNumber, level)
		e.Print()
		e.PrintHex()
		return nil
	},
}

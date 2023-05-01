package cmd

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/Danice123/ckkey/internal"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate [trainer id] [dex number] [level]",
	Short: "Generate a deterministic encounter",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		trainerId, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		dexNumber, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}
		level, err := strconv.Atoi(args[2])
		if err != nil {
			panic(err)
		}

		e, err := internal.CalcEncounter(trainerId, dexNumber, level)
		if err != nil {
			panic(err)
		}
		e.Print()
	},
}

package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Danice123/ckkey/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(findCmd)
}

var findCmd = &cobra.Command{
	Use:   "find [trainer id] [area name] [type]",
	Short: "Generate a deterministic encounter",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		trainerId, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		area := strings.ReplaceAll(strings.ToLower(args[1]), " ", "")
		eType := strings.ToLower(args[2])

		orderMorning, err := internal.GenerateEncounterOrder(trainerId, area, eType, "morning")
		if err != nil {
			panic(err)
		}
		orderDay, err := internal.GenerateEncounterOrder(trainerId, area, eType, "day")
		if err != nil {
			panic(err)
		}
		orderNight, err := internal.GenerateEncounterOrder(trainerId, area, eType, "night")
		if err != nil {
			panic(err)
		}

		fmt.Printf("%s: %s Day\n%v\n", area, eType, orderDay)
		fmt.Printf("%s: %s Night\n%v\n", area, eType, orderNight)
		fmt.Printf("%s: %s Morning\n%v\n", area, eType, orderMorning)
	},
}

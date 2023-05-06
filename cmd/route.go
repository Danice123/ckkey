package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/Danice123/ckkey/internal/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(findCmd)
}

var findCmd = &cobra.Command{
	Use:   "find [trainer id] [area name]",
	Short: "Generate a deterministic encounter",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		trainerId, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if trainerId < 0 || trainerId > 65535 {
			return errors.New("trainer id invalid")
		}
		areaName := strings.ReplaceAll(strings.ToLower(args[1]), " ", "")

		resp := api.BuildEncounterResponse(trainerId, areaName)
		if resp == nil {
			return errors.New("area invalid")
		}

		fmt.Printf("%s\n", resp.Area)
		if resp.Walking != nil {
			fmt.Println("Day:")
			ppEncounters(resp.Walking.Day)
			if resp.Walking.Night != nil {
				fmt.Println("Night:")
				ppEncounters(resp.Walking.Night)
			}
			if resp.Walking.Morning != nil {
				fmt.Println("Morning:")
				ppEncounters(resp.Walking.Morning)
			}
		}

		if resp.Surfing != nil {
			fmt.Println("Surfing:")
			ppEncounters(resp.Surfing)
		}

		if resp.Fishing != nil {
			fmt.Println("Old Rod:")
			ppEncounters(resp.Fishing["Old"].Day)

			fmt.Println("Good Rod Day:")
			ppEncounters(resp.Fishing["Good"].Day)

			fmt.Println("Good Rod Night:")
			ppEncounters(resp.Fishing["Good"].Night)

			fmt.Println("Super Rod Day:")
			ppEncounters(resp.Fishing["Super"].Day)

			fmt.Println("Super Rod Night:")
			ppEncounters(resp.Fishing["Super"].Night)
		}

		if resp.Headbutt != nil {
			fmt.Println("Headbutt:")
			ppEncounters(resp.Headbutt)
		}

		if resp.RockSmash != nil {
			fmt.Println("Rock Smash:")
			ppEncounters(resp.RockSmash)
		}

		if resp.Special != nil {
			for _, s := range resp.Special {
				if s.Pool != nil {
					fmt.Printf("%s:\n", s.Type)
					ppEncounters(s.Pool)
				}
			}
		}

		return nil
	},
}

func ppEncounters(encounters []api.Encounter) {
	for i, e := range encounters {
		name := e.Pokemon
		if len(name) < 6 {
			name += "\t"
		}
		fmt.Printf("\t%d-%s\t(H %d,\tA %d,\tD %d,\tSP %d,\tS %d)\n",
			i+1,
			name,
			e.HealthDV,
			e.DVs.Attack,
			e.DVs.Defense,
			e.DVs.Special,
			e.DVs.Speed,
		)
	}
}

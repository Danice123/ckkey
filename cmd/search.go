package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/Danice123/ckkey/package/roll"
	"github.com/spf13/cobra"
)

func init() {
	searchCmd.Flags().IntVarP(&searchAttack, "attack", "", -1, "Attack value to search for")
	searchCmd.Flags().IntVarP(&searchDefense, "defense", "", -1, "Defense value to search for")
	searchCmd.Flags().IntVarP(&searchSpeed, "speed", "", -1, "Speed value to search for")
	searchCmd.Flags().IntVarP(&searchSpecial, "special", "", -1, "Special value to search for")
	searchCmd.Flags().IntVarP(&searchHp, "hp", "", -1, "HP value to search for")
	rootCmd.AddCommand(searchCmd)
}

var searchAttack int
var searchDefense int
var searchSpeed int
var searchSpecial int
var searchHp int

var searchCmd = &cobra.Command{
	Use:   "search [dex number] [level]",
	Short: "Search for perfect encounters",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		dexNumber, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if dexNumber < 1 || dexNumber > 251 {
			return errors.New("dex ID invalid")
		}
		level, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		if level < 1 {
			return errors.New("level invalid")
		}

		for i := 0; i < 65535; i++ {
			e := roll.CalcDVs(i, dexNumber, level)
			if searchAttack != -1 && e.Attack != searchAttack {
				continue
			}
			if searchDefense != -1 && e.Defense != searchDefense {
				continue
			}
			if searchSpeed != -1 && e.Speed != searchSpeed {
				continue
			}
			if searchSpecial != -1 && e.Special != searchSpecial {
				continue
			}
			if searchHp != -1 && e.CalcHealth() != searchHp {
				continue
			}
			fmt.Printf("Trainer Id: %d\n", i)
			e.Print()
		}
		return nil
	},
}

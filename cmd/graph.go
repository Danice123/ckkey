package cmd

import (
	"os"
	"strconv"

	"github.com/Danice123/ckkey/internal"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(graphCmd)
}

var graphCmd = &cobra.Command{
	Use:   "graph [dex number] [level]",
	Short: "Graph encounters",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		dexNumber, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		level, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}

		encounters := make([]internal.Encounter, 65535)
		for i := 0; i < 65535; i++ {
			e, err := internal.CalcEncounter(i, dexNumber, level)
			if err != nil {
				panic(err)
			}
			encounters[i] = e
		}

		graph := charts.NewScatter()

		start := 23456
		length := 500
		axis := make([]int, length)
		for i := 0; i < length; i++ {
			axis[i] = start + i
		}

		graph.SetXAxis(axis)
		graph.AddSeries("Attack", makeSeries(func(i int) int { return encounters[i].Attack }, start, length))
		graph.AddSeries("Defense", makeSeries(func(i int) int { return encounters[i].Defense }, start, length))
		graph.AddSeries("Speed", makeSeries(func(i int) int { return encounters[i].Speed }, start, length))
		graph.AddSeries("Special", makeSeries(func(i int) int { return encounters[i].Special }, start, length))
		graph.AddSeries("Health", makeSeries(func(i int) int { return encounters[i].CalcHealth() }, start, length))

		f, _ := os.Create("bar.html")
		graph.Render(f)
	},
}

func makeSeries(d func(int) int, start int, length int) []opts.ScatterData {
	items := make([]opts.ScatterData, length)
	for i := 0; i < length; i++ {
		items[i] = opts.ScatterData{
			Value:      d(start + i),
			SymbolSize: 3,
		}
	}
	return items
}

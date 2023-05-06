package cmd

import (
	"net/http"

	"github.com/Danice123/ckkey/internal/api"
	"github.com/julienschmidt/httprouter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve [address]",
	Short: "Start api webserver",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		router := httprouter.New()

		router.GET("/api/:trainerId/rolldv/:dexId", api.RollDVs)
		router.GET("/api/:trainerId/rollencounter/:area", api.RollEncounter)

		return http.ListenAndServe(args[0], router)
	},
}

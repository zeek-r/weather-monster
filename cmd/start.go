package cmd

import (
	"os"

	"github.com/spf13/cobra"
	server "github.com/zeek-r/weather-monster/api/http"
	"github.com/zeek-r/weather-monster/app"
	logger "github.com/zeek-r/weather-monster/pkg/logger"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts the api server",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.InitLogger(os.Stdout)

		application := new(app.App)

		startedApp, err := application.Start()

		if err != nil {
			logger.Error(err, "Error starting app")
			logger.Info("env", map[string]interface{}{"environments": os.Environ()})
			return err
		}

		api := server.NewAPI(startedApp)

		if err := api.Serve(); err != nil {
			startedApp.Stop()
			logger.Error(err, "Error starting API")
			return err
		}
		return nil
	},
}

func init() {
	rootCommand.AddCommand(startCmd)
}

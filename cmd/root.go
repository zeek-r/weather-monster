package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCommand = &cobra.Command{
	Use:   "weather-monster",
	Short: "Weather Monster API",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCommand.PersistentFlags().StringVar(&configFile, "config", "", "config fie(default is run.env)")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		// Load default config.json
		viper.AddConfigPath("./config")
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		viper.ReadInConfig()

		// Override config.json with run.env, if available
		viper.AddConfigPath(".")
		viper.SetConfigName("run")
		viper.MergeInConfig()
		viper.SetEnvPrefix("MonsterAPI")
		viper.AutomaticEnv()
		fmt.Println(os.Environ())
	}
}

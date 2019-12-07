// Copyright © 2017 Pivotal Software

package cli

import (
	"github.com/pivotal-cf/on-demand-service-broker/brokerinitiator"
	"github.com/pivotal-cf/on-demand-service-broker/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var Version = "dev"

func CreateRootCommand(conf *config.Config, odbToolkit *brokerinitiator.ODBToolkit) *cobra.Command {

	// root represents the base command when called without any subcommands
	root := &cobra.Command{
		Use:   "on-demand-service-broker",
		Short: "The service broker API server modified and turned into a CLI",
		Long: "The service broker API server modified and turned into a CLI",
		Version: Version,
	}
	root.Flags().String("configFilePath", "configFilePath", "The config file path")
	root.AddCommand(createGetManifestsCommand(conf, odbToolkit))
	return root
}


func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
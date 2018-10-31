package cmd

import (
	"github.com/spf13/cobra"
	"github.com/openlab-red/vault-secret-fetcher/pkg/fetcher"
	"github.com/spf13/viper"
)

var handlerCmd = &cobra.Command{
	Use:   "start",
	Short: "Retrieves credentials managed by the vault agent and fetch the secret",
	Long:  `Retrieves credentials managed by the vault agent and fetch the secret`,
	Run: func(cmd *cobra.Command, args []string) {
		fetcher.Start()
	},
}

func init() {
	RootCmd.AddCommand(handlerCmd)
	viper.SetDefault("log-level", "INFO")
	viper.SetDefault("vault-insecure", false)
	viper.SetDefault("vault-token-fetcher-cron", "1 * * * * *")
	viper.SetDefault("properties-type", "yaml")
}

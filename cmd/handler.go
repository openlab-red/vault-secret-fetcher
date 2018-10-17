package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vault-agent-token-handler/handler"
	"github.com/spf13/viper"
)

var handlerCmd = &cobra.Command{
	Use:   "handler",
	Short: "handler retrieves credentials managed by the vault agent",
	Long:  `handler retrieves credentials managed by the vault agent`,
	Run: func(cmd *cobra.Command, args []string) {
		handler.Start()
	},
}

func init() {
	RootCmd.AddCommand(handlerCmd)
	viper.SetDefault("log-level", "INFO")
	viper.SetDefault("vault-insecure", false)
	viper.SetDefault("vault-token-handler-cron", "1 * * * * *")
}

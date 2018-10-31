package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "vault-token-fetcher",
	Short: "Vault Token Handler sidecar container for kubernetes applications",
	Long:  "Vault Token Handler sidecar container for kubernetes applications",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv() // read in environment variables that match
}
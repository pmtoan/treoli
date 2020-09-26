package cmd

import (
	"github.com/pmtoan/treoli/config"
	"github.com/pmtoan/treoli/trello"
	"github.com/spf13/cobra"
)

var key string
var token string

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(setConfigCmd)
	setConfigCmd.Flags().StringVarP(&key, "key", "k", "", "Application key")
	setConfigCmd.Flags().StringVarP(&token, "token", "t", "", "Account token")
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set or get config for treoli",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var setConfigCmd = &cobra.Command{
	Use:   "set",
	Short: "Set config for treoli",
	Run: func(cmd *cobra.Command, args []string) {
		client = trello.InitClient()
		client.SetConfig(config.TrelloAppKeyConfigName, key)
		client.SetConfig(config.TrelloTokenConfigName, token)
	},
}

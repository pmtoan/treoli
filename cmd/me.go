package cmd

import (
	"fmt"

	"github.com/pmtoan/treoli/config"
	"github.com/pmtoan/treoli/lib"
	"github.com/pmtoan/treoli/trello"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(meCmd)
}

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Get Trello account information",
	Run: func(cmd *cobra.Command, args []string) {
		client = trello.InitClient()
		myInfo, err := client.GetMe()
		if err != nil {
			fmt.Print(config.MessageTrelloAccountFound)
		} else {
			fmt.Println("Your Trello account information:")
			lib.DefaultPrint(myInfo)
		}
	},
}

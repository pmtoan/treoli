package cmd

import (
	"fmt"
	"os"

	"github.com/pmtoan/treoli/config"
	"github.com/pmtoan/treoli/lib"
	"github.com/pmtoan/treoli/trello"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(boardCmd)
	boardCmd.AddCommand(listBoardCmd)
}

var boardCmd = &cobra.Command{
	Use:   "board",
	Short: "Boards management",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var listBoardCmd = &cobra.Command{
	Use:   "list",
	Short: "List your boards",
	Run: func(cmd *cobra.Command, args []string) {
		command := trello.InitClient()
		if command == nil {
			fmt.Print(config.MessageTrelloAccountFound)
		} else {
			client = command
			boards, err := client.GetMyBoards()
			if err != nil {
				fmt.Print(config.MessageTrelloAccountFound)
				os.Exit(1)
			} else {
				lib.DefaultPrint(boards)
			}
		}
	},
}

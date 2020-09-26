package cmd

import (
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
		client = trello.InitClient()
		boards, err := client.GetMyBoards()
		if err != nil {
			panic(err)
		} else {
			lib.DefaultPrint(boards)
		}
	},
}

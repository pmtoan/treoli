package cmd

import (
	"fmt"
	"os"

	"github.com/pmtoan/treoli/config"
	"github.com/pmtoan/treoli/trello"
	"github.com/spf13/cobra"
)

var client trello.Client

var rootCmd = &cobra.Command{
	Use:   "treoli",
	Short: "Treoli - Trello CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use \"treoli --help\" for more information about a command.")
	},
}

func Execute() {
	config.LoadConfig()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List items",
	Long:  "Lists items related to fpm.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Listing items...")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	

	"github.com/spf13/cobra"
	"mycli/data"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "All of our notes can be seen in a list",
	Long: `All of our notes can be seen in a list`,
	Run: func(cmd *cobra.Command, args []string) {
		data.DisplayAllNotes()
	},
}

func init() {
	noteCmd.AddCommand(listCmd)

	
}

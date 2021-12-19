/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	

	"github.com/spf13/cobra"
)

// noteCmd represents the note command
var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "A brief description of your note you want to store",
	Long: `A  brief description of your note you want to store`,
	
}

func init() {
	rootCmd.AddCommand(noteCmd)

}

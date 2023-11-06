/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package info

import (
	"fmt"

	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var InfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Command is related to info",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println("Info called")
	},
}

func init() {
	

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

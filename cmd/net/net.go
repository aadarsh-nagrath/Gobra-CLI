/*
Copyright © 2023 NAME HERE anagrath1@gmail.com

*/
package net

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NetCmd represents the net command
var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "net is a pkg to hold network related commands",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println("so that if somebody called net command directly it could display the help instructions")
		// so that if somebody called net command directly it could display the help instructions
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// netCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// netCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

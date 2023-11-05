/*
Copyright © 2023 NAME HERE anagrath1@gmail.com

*/
package net

import (
	"fmt"

	"github.com/spf13/cobra"
)
var (
	urlPath string
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This command pings a remote url and return a response",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ping called")
		//Here we will define our ping logic
	},
}

func init() {
	pingCmd.Flags().StringVarP(&urlPath, "url" , "u", "", "THis is url to ping")

	//persistent url flag
	if err := pingCmd.MarkFlagRequired("url"); err != nil {
		fmt.Println(err)
	}
	NetCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2023 NAME HERE anagrath1@gmail.com

*/
package net

import (
	"fmt"
	"net/http"
	"time"
	"github.com/spf13/cobra"
)
var (
	urlPath string
	//logic
	client = http.Client{
		Timeout: 2* time.Second,
	}
)

func ping(domain string) (int,error) {
	url:= "http://" + domain
	req, err := http.NewRequest("HEAD", url, nil)
	if err!= nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "This command pings a remote url and return a response",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ping called")
		
		if resp, err := ping(urlPath); err!= nil {
			fmt.Println(err)
		}else{
			fmt.Println(resp)
		}
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

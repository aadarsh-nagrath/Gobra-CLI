package info

import (
	"fmt"
	"github.com/ricochet2200/go-disk-usage/du"
	"github.com/spf13/cobra"
)

// diskUsageCmd represents the discUsage command
var diskUsageCmd = &cobra.Command{
	Use:   "diskUsage",
	Short: "THe command briefs about the disk usage",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		usage := du.NewDiskUsage(".")

		fmt.Println(usage)
	},
}

func init() {
	InfoCmd.AddCommand(diskUsageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// discUsageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// discUsageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"reminders/shared"

	"github.com/spf13/cobra"
)

// startShiftCmd represents the startShift command
var startShiftCmd = &cobra.Command{
	Use:   "startShift",
	Short: "startShift will start your shift",
	Long:  "startShift will start your shift",

	Run: func(cmd *cobra.Command, args []string) {
		shared.Scheduler()
	},
}

func init() {
	rootCmd.AddCommand(startShiftCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startShiftCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startShiftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

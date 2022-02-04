/*
Package cmd
Copyright © 2022 snapsh.io tomer@snapsh.io

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Creating basic http web server",
		Long:  fmt.Sprintf("Creating basic http web server without API and docs"),
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func addInitCommand() {
	rootCmd.AddCommand(initCmd)
}

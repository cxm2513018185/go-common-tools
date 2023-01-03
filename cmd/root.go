package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{}

func init() {
	// 注册子命令
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

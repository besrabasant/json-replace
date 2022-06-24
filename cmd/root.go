package cmd

import (
	"github.com/spf13/cobra"
)

func RootCmd() {
	cmdReplace := ReplaceCmd()
	cmdGet := GetCmd()

	var rootCmd = &cobra.Command{Use: "json-replace"}
	rootCmd.AddCommand(cmdReplace, cmdGet)
	rootCmd.Execute()
}
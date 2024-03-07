/*
Copyright © 2024 MATHIAS MARCHETTI aquemaati@gmail.com
*/
package cmd

import (
	"fmt"
	cmd "gogol/cmd/new/basic"

	"github.com/spf13/cobra"
)

// NewCmd represents the new command
var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "create a new application from scratch",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called, create default user project")
	},
}
var foo string

func init() {
	NewCmd.AddCommand(cmd.BasicCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	NewCmd.PersistentFlags().StringVarP(&foo, "foo", "f", "default", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:

}

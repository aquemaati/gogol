/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gogol/internal/tools"
	"log"

	"github.com/spf13/cobra"
)

// basicCmd represents the basic command
var BasicCmd = &cobra.Command{
	Use:   "basic",
	Short: "Create a simple project to learn or test ideas",
	Long:  `Language avaiable : go, python, julia, ...`,
	Run: func(cmd *cobra.Command, args []string) {	
		fmt.Println("basic called")

		data, err := tools.GetDatas("basic", lang)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(data)
		f, _ := tools.LangIsInstalled(lang)
		if !f {
			log.Fatalln(lang, "programming language not installed")
		}
	},
}
var lang string

func init() {
	//rootCmd.AddCommand(basicCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// basicCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// basicCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	BasicCmd.Flags().StringVarP(&lang, "lang", "l", "none", "Specify the programming language")
}

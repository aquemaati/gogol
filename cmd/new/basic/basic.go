/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"gogol/internal/tools"
	"log"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

// basicCmd represents the basic command
var BasicCmd = &cobra.Command{
	Use:   "basic",
	Short: "Create a simple project to learn or test ideas",
	Long:  `Language avaiable : go, python, julia, ...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("basic called")

		// STEP 1
		// get generald datas for basics projects with the
		// programming language selected
		data, err := tools.GetDatas("basic", lang)
		if err != nil {
			log.Fatalln(err)
		}

		// STEP 2
		// get command to check if the programming language
		// is properly installed
		setUp, _ := tools.GetCmdCheckInstall(data.LinkSetup)

		// STEP 3
		// check if the language is wel installed, if not
		// Show the what the user have to do
		f, _ := tools.LangIsInstalled(lang, setUp.Cmd)
		if !f {
			arch := runtime.GOARCH
			// TODO check as well with runtime.goose ans update json
			log.Fatalln(lang, "programming language not installed", "go to ", setUp.Install[arch], "to download it for", runtime.GOARCH)
		}

		// STEP 4
		// create de root directory for the project and then go inside
		if err := os.Mkdir(name, 0777); err != nil {
			log.Fatalln(err)
		}
		if err := os.Chdir(name); err != nil {
			log.Fatalln(err)
		}

		// Step 5
		// Executing get json instructions
		basic, err := GetBasicJson(data.Link)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(basic)

	},
}
var lang, name string

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
	BasicCmd.Flags().StringVarP(&name, "name", "n", "project", "Specify the name")
}

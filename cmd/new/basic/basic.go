/*
Copyright © 2024 MATHIAS MARCHETTI aquemaati@gmail.com
*/
package cmd

import (
	"fmt"
	"gogol/internal/tools"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

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
		// Going inside this directory to execute new cmd
		if err := os.Chdir(name); err != nil {
			log.Fatalln(err)
		}

		// STEP 5
		// get json instructions
		basic, err := GetBasicJson(data.Link)
		if err != nil {
			log.Fatalln(err)
		}

		//STEP 6
		// Execute precommand
		// TODO create condition for dependancies
		for imap, cm := range basic.PreCmd {
			for i, v := range cm {
				if v == "%s" {
					basic.PreCmd[imap][i] = name
				}
			}
			if err := exec.Command(cm[0], cm[1:]...).Run(); err != nil {
				log.Fatalln(err)
			}
		}

		// STEP 7
		// Create files and write inside
		for _, f := range basic.Files {
			file, err := os.Create(f.Name)
			if err != nil {
				log.Fatalln(err)
			}
			defer file.Close()
			fmt.Printf("%s has been created\n", f.Name)
			// Writing inside
			text := strings.ReplaceAll(f.Content, "%s", name)
			if _, err := file.WriteString(text); err != nil {
				log.Fatalln(err)
			}
		}
		// STEP 8
		// Create all subdirectories
		if err := CreateFilesAndPackages(basic.Subdirectories, name); err != nil {
			log.Fatalln(err)
		}

	},
}
var lang, name string

func init() {
	// Here you will define your flags and configuration settings.

	BasicCmd.Flags().StringVarP(&lang, "lang", "l", "none", "Specify the programming language")
	BasicCmd.Flags().StringVarP(&name, "name", "n", "project", "Specify the name")
}

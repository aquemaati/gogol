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
		fmt.Println("")
		fmt.Println("Creating a", cmd.Use, name, "application with", lang, "programming language...")
		fmt.Println("")

		// STEP 1
		// get general datas for basics projects with the
		// programming language selected
		fmt.Println("---> Fetching datas for", cmd.Use, "app with", lang, "Programming language")
		data, err := tools.GetDatas("basic", lang)
		if err != nil {
			log.Fatalln(err)
		}

		// STEP 2
		// get command to check if the programming language
		// is properly installed
		fmt.Println("---> Checking if your os has the requirements for", lang, "programming languages")
		setUp, _ := tools.GetCmdCheckInstall(data.LinkSetup)

		// STEP 3
		// check if the language is wel installed, if not
		// Show the what the user have to do
		f, _ := tools.LangIsInstalled(lang, setUp.Cmd)
		if !f {
			arch := runtime.GOARCH
			// TODO check as well with runtime.goose ans update json
			log.Fatalln("ERROR :", lang, "programming language not installed", "go to ", setUp.Install[arch], "to download it for", runtime.GOARCH)
		}

		// STEP 4
		// create de root directory for the project and then go inside
		fmt.Println("+++> Creating root directory", name)
		if err := os.Mkdir(name, 0777); err != nil {
			log.Fatalln(err)
		}
		// Going inside this directory to execute new cmd
		if err := os.Chdir(name); err != nil {
			log.Fatalln(err)
		}

		// STEP 5
		// get json instructions
		fmt.Println("---> Fetching", lang, cmd.Use, "instructions")
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
			fmt.Println("")
			fmt.Print("***> Executing command : ")
			out, err := exec.Command(cm[0], cm[1:]...).CombinedOutput()
			if err != nil {
				fmt.Println(string(out))
				fmt.Println("ERROR : while executing commands")
				log.Fatalln(err)
			}
			for _, s := range cm {
				fmt.Print(s, " ")
			}
			fmt.Println("")
			fmt.Println(string(out))
		}

		// STEP 7
		// Create files and write inside
		for _, f := range basic.Files {
			file, err := os.Create(f.Name)
			if err != nil {
				log.Fatalln(err)
			}
			defer file.Close()
			fmt.Printf("+++> Creating file : %s\n", f.Name)
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
		fmt.Println("")
		fmt.Println("CONGRATULATION", cmd.Use, name, "Has been created succesfully!!!")
		fmt.Println("end instuctions")
	},
}
var lang, name string

func init() {
	// Here you will define your flags and configuration settings.

	BasicCmd.Flags().StringVarP(&lang, "lang", "l", "none", "Specify the programming language")
	BasicCmd.Flags().StringVarP(&name, "name", "n", "project", "Specify the name")
}

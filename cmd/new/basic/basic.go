/*
Package cmd handles the basic command for creating simple projects to learn or test ideas.
*/
package cmd

import (
	"fmt"
	"gogol/internal/messages"
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
	Long:  `Languages available: Go, Python, Julia, ...`,
	Run: func(cmd *cobra.Command, args []string) {
		// Starting message
		fmt.Println(messages.StartingProject(cmd.Use, name, lang))

		// STEP 1: Get general data for basic projects with the selected programming language
		data, err := tools.GetDatas("basic", lang)
		if err != nil {
			fmt.Println(messages.ErrFetch(lang), err)
			return
		}

		// STEP 2: Get command to check if the programming language is properly installed
		setUp, err := tools.GetCmdCheckInstall(data.LinkSetup)
		if err != nil {
			fmt.Println(messages.ErrFetch(data.LinkSetup), err)
			return
		}

		// STEP 3: Check if the language is installed properly; show instructions if not
		arch := runtime.GOARCH // Architecture of the user computer
		ops := runtime.GOOS    // OS of the user
		outs, f, err := tools.LangIsInstalled(setUp.CheckCommand[ops], lang)
		if err != nil {
			fmt.Println(messages.ERR, err)
		}
		fmt.Println(outs)
		if !f {
			fmt.Println(messages.ErrLangInstall(lang, ops, arch))
			err := tools.HandleSetUp(ops, arch, setUp)
			if err != nil {
				fmt.Println(messages.ERR, err, "could not install the programming language properly")
				return
			}
			return
		}

		// STEP 4: Create the root directory for the project and navigate inside
		fmt.Println(messages.DirBuilding(name))
		if err := os.Mkdir(name, 0777); err != nil {
			fmt.Println(messages.ErrDir(name), err)
			return
		}
		if err := os.Chdir(name); err != nil {
			log.Fatalln(messages.ERR, err)
		}

		// STEP 5: Get JSON instructions
		basic, err := GetBasicJson(data.Link)
		if err != nil {
			log.Fatalln(messages.ErrFetch(data.Link), err)
		}

		// STEP 6: Execute pre-commands
		if err := ExecutePreCommands(basic.PreCmd, name); err != nil {
			log.Fatalln(err)
		}

		// STEP 7: Create files and write inside
		if err := CreateAndWriteFiles(basic.Files, name); err != nil {
			log.Fatalln(err)
		}

		// STEP 8: Create all subdirectories
		if err := CreateFilesAndPackages(basic.Subdirectories, name); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("")

		// STEP 9: Add more files as requested by the user with parent directories
		if err := ChangeToParentDirectoryWithSpecificName(name); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Files to add:", addf)
		if err := CreateAdditionalFiles(addf); err != nil {
			log.Fatalln(err)
		}

		// TODO NEW STEP : ask if the user wants to create a repo on github, and init a git
		// Then create steps, ands ask user enter information or dowllad github cli, then config users settinfs
		// STEP 10: Print final instructions
		fmt.Println(messages.Congrat(cmd.Use, name, lang))
		fmt.Print("\t\t - Access your new directory: cd ", name, "\n")
		for _, v := range basic.EndInstruction {
			fmt.Println("\t\t", "-", v)
		}
		fmt.Println("")
	},
}

var lang, name string
var dep, addf []string

func init() {
	// Define flags
	BasicCmd.Flags().StringVarP(&lang, "lang", "l", "none", "Specify the programming language")
	BasicCmd.Flags().StringVarP(&name, "name", "n", "project", "Specify the name")
	BasicCmd.Flags().StringArrayVarP(&dep, "dep", "d", []string{}, "Specify dependencies")
	BasicCmd.Flags().StringArrayVarP(&addf, "addf", "", []string{}, "Specify files to add")
}

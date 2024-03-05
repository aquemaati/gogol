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
		// FLag test

		//		log.Fatalln(dep, len(dep))
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
		// TODO create conditions for type of commands => create a switch ?

		for key, cm := range basic.PreCmd {
			if key == "init" {
				for i, v := range cm {
					if v == "%s" {
						basic.PreCmd[key][i] = name
					}
				}
				executeCommand(cm)
			}
			if key == "dep" {
				for _, d := range dep {
					cm = append(cm, d)
					executeCommand(cm)
				}
			}
		}
		// for id := range dep {
		// 	if key == "dep" && len(dep) > 0 {
		// 		for i, v := range cm {
		// 			if v == "%s" {
		// 				if len(dep) > 0 {
		// 					basic.PreCmd[key][i] = dep[id]
		// 					continue
		// 				}
		// 			}
		// 		}
		// 	} else if key == "dep" && len(dep) == 0 {
		// 		break
		// 	}
		// 	executeCommand(cm)
		// }

		// fmt.Println("")
		// fmt.Print("***> Executing command : ")
		// out, err := exec.Command(cm[0], cm[1:]...).CombinedOutput()
		// if err != nil {
		// 	fmt.Println("\n", string(out))
		// 	fmt.Println("ERROR : while executing commands")
		// 	log.Fatalln(err)
		// }
		// for _, s := range cm {
		// 	fmt.Print(s, " ")
		// }
		// fmt.Println("")
		// fmt.Println(string(out))

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
var dep []string

func init() {
	// Here you will define your flags and configuration settings.

	BasicCmd.Flags().StringVarP(&lang, "lang", "l", "none", "Specify the programming language")
	BasicCmd.Flags().StringVarP(&name, "name", "n", "project", "Specify the name")
	BasicCmd.Flags().StringArrayVarP(&dep, "dep", "d", []string{}, "specify dependacies")
}

func executeCommand(cm []string) {
	fmt.Println("")
	fmt.Print("***> Executing command : ")

	// Exécute la commande et récupère la sortie combinée
	out, err := exec.Command(cm[0], cm[1:]...).CombinedOutput()

	// Vérifie s'il y a une erreur lors de l'exécution de la commande
	if err != nil {
		fmt.Println("\n", string(out))
		fmt.Println("ERROR : while executing commands")
		log.Fatalln(err)
	}

	// Affiche la commande exécutée
	for _, s := range cm {
		fmt.Print(s, " ")
	}
	fmt.Println("")

	// Affiche la sortie de la commande
	fmt.Println("\n", string(out))
}

/*
Copyright © 2024 MATHIAS MARCHETTI aquemaati@gmail.com
*/
package cmd

import (
	"fmt"
	"gogol/internal/messages"
	"gogol/internal/tools"
	"log"
	"os"
	"os/exec"
	"path/filepath"
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

		// Sarter message
		fmt.Println(messages.StartingProject(cmd.Use, name, lang))

		// STEP 1
		// get general datas for basics projects with the
		// programming language selected
		data, err := tools.GetDatas("basic", lang)
		if err != nil {
			log.Fatalln(err)
		}

		// STEP 2
		// get command to check if the programming language
		// is properly installed
		//TODO include in LangIsInstalled
		setUp, _ := tools.GetCmdCheckInstall(data.LinkSetup)

		// STEP 3
		// check if the language is wel installed, if not
		// Show what the user have to do
		arch := runtime.GOARCH
		ops := runtime.GOOS
		outs, f, _ := tools.LangIsInstalled(setUp.CheckCommand[ops], lang)
		fmt.Println(outs)
		if !f {
			fmt.Println(messages.ErrLangInstall(lang, ops, arch))
			// allow the user to download what he need
			// TODO check if brew or choco is installed
			err := tools.HandleSetUp(ops, arch, setUp)
			if err != nil {
				fmt.Println(err, "could not install properly the programming language")
				return
			}
		}

		// STEP 4
		// create de root directory for the project and then go inside
		fmt.Println(messages.DirBuilding(name))
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

		for key, cm := range basic.PreCmd {
			if key == "init" && len(cm) != 0 {
				for i, v := range cm {
					if v == "%s" {
						basic.PreCmd[key][i] = name
					}
				}
				fmt.Println("")
				executeCommand(cm)
			} else if key == "dep" && len(cm) != 0 {
				for _, d := range dep {
					cm = append(cm, d)
					fmt.Println(messages.Fetching(d))
					executeCommand(cm)
				}
			} else {
				continue
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
			// fmt.Printf("+++> Creating file : %s\n", f.Name)
			fmt.Println(messages.FileBuilding(f.Name))
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

		// STEP 9
		// Add more files as asked by the user with parent directories
		// Back to main repertory
		if err := ChangeToParentDirectoryWithSpecificName(name); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("files to add", addf)
		for _, filePath := range addf {

			err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
			if err != nil {
				fmt.Println("Erreur lors de la création des répertoires parents:", err)
				return
			}
			// Crée le fichier
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Println("Erreur lors de la création du fichier:", err)
				return
			}
			fmt.Println(messages.FileBuilding(filePath))
			defer file.Close()
		}

		// SAY everything went ok
		// fmt.Println("\nCONGRATULATION", cmd.Use, name, "Has been created successfully!!!")
		// // Print last instructions
		fmt.Println(messages.Congrat(cmd.Use, name, lang))
		fmt.Print("\t\t Access your new directory: cd ", name, "\n")
		for _, v := range basic.EndInstruction {
			fmt.Println("\t\t", v)
		}
	},
}

var lang, name string
var dep, addf []string

func init() {
	// Here you will define your flags and configuration settings.
	BasicCmd.Flags().StringVarP(&lang, "lang", "l", "none", "Specify the programming language")
	BasicCmd.Flags().StringVarP(&name, "name", "n", "project", "Specify the name")
	BasicCmd.Flags().StringArrayVarP(&dep, "dep", "d", []string{}, "specify dependacies")
	BasicCmd.Flags().StringArrayVarP(&addf, "addf", "", []string{}, "specify file to add")
}

func executeCommand(cm []string) {
	fmt.Println("")
	// fmt.Print("***> Executing command : ")

	// Exécute la commande et récupère la sortie combinée
	out, err := exec.Command(cm[0], cm[1:]...).CombinedOutput()
	fmt.Println(messages.ExecCmd(strings.Join(cm, " "), string(out)))
	// Vérifie s'il y a une erreur lors de l'exécution de la commande
	if err != nil {
		fmt.Println("\n", string(out))
		fmt.Println("ERROR : while executing commands")
		log.Fatalln(err)
	}
}

// allow to come back to the main directory
func ChangeToParentDirectoryWithSpecificName(targetDir string) error {
	// Get the absolute path of the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	// Iterate through parent directories until the target directory is found
	for {
		// Check if the current directory is the target directory
		if filepath.Base(currentDir) == targetDir {
			// If so, you are in the correct directory, no need to go further up
			break
		}

		// Move up one directory level
		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			// If the current directory is already the root, stop the search
			return fmt.Errorf("target directory not found: %s", targetDir)
		}

		// Update the current directory to the parent directory
		currentDir = parentDir
	}

	// Change the working directory to the target directory
	if err := os.Chdir(currentDir); err != nil {
		return err
	}

	fmt.Println("Working directory updated:", currentDir)
	return nil
}

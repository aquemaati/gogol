package cmd

import (
	"fmt"
	"gogol/internal/messages"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// This function use recursivity to create subfiles and subdirectories
// Launch aswell comandlines
func CreateFilesAndPackages(subdirs []Subdirectory, name string) error {
	for _, subdir := range subdirs {
		fmt.Println(messages.DirBuilding(subdir.Name))

		if err := os.Mkdir(subdir.Name, 0777); err != nil {
			return fmt.Errorf("ERROR: can't create the directory %s", subdir.Name)
		}
		// Changing directory
		if err := os.Chdir(subdir.Name); err != nil {
			return err
		}

		for _, file := range subdir.Files {
			fmt.Println(messages.FileBuilding(file.Name))
			// Create file
			fil, err := os.Create(file.Name)
			if err != nil {
				return fmt.Errorf("%s can't be created", file.Name)
			}
			//fmt.Printf("%s has been created\n", file.Name)
			defer fil.Close()

			text := strings.ReplaceAll(file.Content, "%s", name)

			if _, err := fil.Write([]byte(text)); err != nil {
				return err
			}
		}
		CreateFilesAndPackages(subdir.Subdirectories, name)
	}
	return nil
}

func ExecutePreCommands(preCmd map[string][]string, name string) error {
	// Execute pre-commands
	for key, cm := range preCmd {
		if key == "init" && len(cm) != 0 {
			for i, v := range cm {
				if v == "%s" {
					cm[i] = name
				}
			}
			if err := ExecuteCommand(cm); err != nil {
				return err
			}
		} else if key == "dep" && len(cm) != 0 {
			for _, d := range dep {
				cm = append(cm, d)
				fmt.Println(messages.Fetching(d))
				if err := ExecuteCommand(cm); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func CreateAndWriteFiles(files []File, name string) error {
	// Create files and write inside
	for _, f := range files {
		file, err := os.Create(f.Name)
		if err != nil {
			return err
		}
		defer file.Close()
		fmt.Println(messages.FileBuilding(f.Name))
		text := strings.ReplaceAll(f.Content, "%s", name)
		if _, err := file.WriteString(text); err != nil {
			return err
		}
	}
	return nil
}

func ExecuteCommand(cm []string) error {
	// Execute the command
	cmd := exec.Command(cm[0], cm[1:]...)
	out, err := cmd.CombinedOutput()
	fmt.Println(messages.ExecCmd(strings.Join(cm, " "), string(out)))
	if err != nil {
		fmt.Println("\n", string(out))
		fmt.Println("ERROR: while executing commands")
		return err
	}
	return nil
}

func CreateAdditionalFiles(addf []string) error {
	// Create additional files
	for _, filePath := range addf {
		err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		if err != nil {
			return err
		}
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		fmt.Println(messages.FileBuilding(filePath))
		defer file.Close()
	}
	return nil
}

// ChangeToParentDirectoryWithSpecificName moves to the parent directory with the specified name
func ChangeToParentDirectoryWithSpecificName(targetDir string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	for {
		if filepath.Base(currentDir) == targetDir {
			break
		}
		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return fmt.Errorf("target directory not found: %s", targetDir)
		}
		currentDir = parentDir
	}
	if err := os.Chdir(currentDir); err != nil {
		return err
	}
	fmt.Println("Working directory updated:", currentDir)
	return nil
}

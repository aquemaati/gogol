package cmd

import (
	"fmt"
	"os"
	"strings"
)

// This function use recursivity to create files and directories
// Launch aswell comandlines
func CreateFilesAndPackages(subdirs []Subdirectory, name string) error {
	for _, subdir := range subdirs {
		fmt.Printf("Package: %s\n", subdir.Name)
		if err := os.Mkdir(subdir.Name, 0777); err != nil {
			return fmt.Errorf("can't create the directory %s", subdir.Name)
		}
		// Changing directory
		if err := os.Chdir(subdir.Name); err != nil {
			return err
		}

		for _, file := range subdir.Files {
			fmt.Printf("File: %s\n", file.Name)
			// Create file
			fil, err := os.Create(file.Name)
			if err != nil {
				return fmt.Errorf("%s can't be created", file.Name)
			}
			fmt.Printf("%s has been created\n", file.Name)
			defer fil.Close()

			text := strings.ReplaceAll(file.Content, "%s", name)

			if _, err := fil.Write([]byte(text)); err != nil {
				return err
			}
		}
		CreateFilesAndPackages(subdirs, name)
	}
	return nil
}

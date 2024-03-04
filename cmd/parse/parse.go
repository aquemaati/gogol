/*
Copyright © 2024 MATHIAS MARCHETTI aquemaati@gmail.com
// allow the user to create a template and transform it in json format with a help of config file
must be a flag? complete ?
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var ParseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Transform file in strinf",
	Long: `for now, transform one file in string, but later, can transform repo in json
with a yaml file to help`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("parse called")
		file, err := os.ReadFile(args[0])
		if err != nil {
			log.Fatalln(err)
		}
		//text := string(file)
		text := formatString(string(file))
		fmt.Println(text)
	},
}

// Function to format a string into an escaped string
func formatString(input string) string {
	// Replacement map for special characters
	replacements := map[string]string{
		"\\": "\\\\", // Backslash
		"\n": "\\n",  // Newline
		"\r": "\\r",  // Carriage return
		"\t": "\\t",  // Horizontal tab
		"\"": "\\\"", // Double quote
		"'":  "\\'",  // Single quote
	}

	// Iterate over each character in the input string
	result := ""
	for _, char := range input {
		// Convert the character to a string
		charStr := string(char)
		// Check if there is a replacement for this character
		replacement, found := replacements[charStr]
		if found {
			// If there is a replacement, append it to the final result
			result += replacement
		} else {
			// Otherwise, simply append the character to the final result
			result += charStr
		}
	}

	return result
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

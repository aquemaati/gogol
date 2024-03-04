package tools

import (
	"fmt"
	"os/exec"
	"strings"
)

// We want to be sure that the programming language is well installed
func LangIsInstalled(s string) (bool, error) {
	cmd := exec.Command(s, "version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err, "1")
		return false, err
	}
	if strings.Contains(string(out), "not found") {
		fmt.Println(err)
		return false, nil
	}
	return true, nil
}

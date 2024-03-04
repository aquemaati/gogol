package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

type SetUp struct {
	Cmd     []string          `json:"cmd"`
	Install map[string]string `json:"install"`
}

// We want to be sure that the programming language is well installed
func LangIsInstalled(s string, cmdLang []string) (bool, error) {
	cmd := exec.Command(cmdLang[0], cmdLang[1:]...)
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

// function to get the command line to check if the programming language
// is properly installed
func GetCmdCheckInstall(url string) (SetUp, error) {
	setUp := new(SetUp)
	resp, err := http.Get(url)
	if err != nil {
		return *setUp, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(setUp); err != nil {
		return *setUp, err
	}
	return *setUp, nil
}

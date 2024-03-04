package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

type Installer struct {
	Os  string `json:"os"`
	Cmd string `json:"link"`
}

type SetUp struct {
	Cmd     []string    `json:"cmd"`
	Install []Installer `json:"install"`
}

// We want to be sure that the programming language is well installed
func LangIsInstalled(s string, cmdLang []string) (bool, error) {
	cmd := exec.Command(cmdLang[0], cmdLang...)
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
func GetCmdCheckInstall(url string) ([]string, error) {
	setUp := new(SetUp)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(setUp); err != nil {
		return nil, err
	}
	return setUp.Cmd, nil
}

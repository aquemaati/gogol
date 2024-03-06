package tools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type CheckConfig struct {
	CheckCommand map[string]string              `yaml:"check_command"`
	Instructions map[string]map[string][]string `yaml:"instructions"`
}

// We want to be sure that the programming language is well installed
func LangIsInstalled(s string) (string, bool, error) {
	cmdLang := strings.Split(s, " ")
	if len(cmdLang) > 1 {
		return "", false, errors.New("bad command")
	}
	cmd := exec.Command(cmdLang[0], cmdLang[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err, "1")
		return string(out), false, err
	}
	// TODO adapt to architecture
	if string(out) == "" {
		fmt.Println(err, "not installed")
		return string(out), false, nil
	}

	return string(out), true, nil
}

// function to get the command line to check if the programming language
// is properly installed on the user'os
func GetCmdCheckInstall(url string) (CheckConfig, error) {
	setUp := new(CheckConfig)
	resp, err := http.Get(url)
	if err != nil {
		return *setUp, err
	}
	defer resp.Body.Close()
	yamlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return *setUp, err
	}
	err = yaml.Unmarshal(yamlData, &setUp)
	if err != nil {
		log.Fatalf("Erreur de parsing du fichier YAML : %v", err)
	}

	return *setUp, nil
}

func OpenWebpage(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}
func IsURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}
	u, err := url.Parse(input)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}
func HandleSetUp(ops, arch string, config CheckConfig) {
	for i, v := range config.Instructions[ops][arch] {
		fmt.Println(i+1, v)
	}
	var num string
	fmt.Println("select")
	fmt.Scanf("%s", &num)

	selected, _ := strconv.Atoi(num)

	switch {
	case selected-1 >= 0 && selected-1 <= len(config.Instructions[ops][arch]) && IsURL(config.Instructions[ops][arch][selected-1]):

		err := OpenWebpage(config.Instructions[ops][arch][selected-1])
		if err != nil {
			log.Fatalln(err)
		}
	}
}

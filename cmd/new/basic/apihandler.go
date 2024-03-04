package cmd

import (
	"encoding/json"
	"net/http"
)

// handling basic api
type BASIC struct {
	PreCmd         map[string][]string
	Files          []File
	Subdirectories []Subdirectory
}

type File struct {
	Name        string
	Description string
	Content     string
	Cmd         [][]string
}
type Subdirectory struct {
	Name           string
	Description    string
	Content        string
	Cmd            [][]string
	Files          []File
	Subdirectories []Subdirectory
}

func GetBasicJson(url string) (BASIC, error) {
	basic := new(BASIC)

	resp, err := http.Get(url)
	if err != nil {
		return *basic, err
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(basic); err != nil {
		return *basic, err
	}
	return *basic, nil

}

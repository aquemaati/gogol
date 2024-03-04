package cmd

import (
	"encoding/json"
	"net/http"
)

// handling basic api
type BASIC struct {
	PreCmd         map[string][]string `json:"precmd"`
	Files          []File              `json:"files"`
	Subdirectories []Subdirectory      `json:"subdirectories"`
}

type File struct {
	Name        string     `json:"name"`
	Content     string     `json:"content"`
	Description string     `json:"description"`
	Cmd         [][]string `json:"cmd"`
}
type Subdirectory struct {
	Name           string         `json:"name"`
	Description    string         `json:"description"`
	Cmd            [][]string     `json:"cmd"`
	Files          []File         `json:"files"`
	Subdirectories []Subdirectory `json:"subdirectories"`
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

package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"gogol/internal/messages"
	"net/http"
)

// handling basic api
type BASIC struct {
	Description    string              `json:"description"`
	PreCmd         map[string][]string `json:"precmd"`
	Files          []File              `json:"files"`
	Subdirectories []Subdirectory      `json:"subdirectories"`
	PostCmd        map[string][]string `json:"postcmd"`
	EndInstruction []string            `json:"endinstruction"`
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
	fmt.Println(messages.Fetching(url))
	basic := new(BASIC)

	resp, err := http.Get(url)
	if err != nil {
		return *basic, errors.New("get")
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(basic); err != nil {
		return *basic, errors.New("decoder")
	}
	return *basic, nil

}

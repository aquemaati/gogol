package tools

import (
	"encoding/json"
	"fmt"
	"gogol/internal/messages"
	"net/http"
)

// Allow to acces the data.json api

type DATAS []Data

type Data struct {
	Name string `json:"name"`
	Lang []Lang `json:"languages"`
}
type Lang struct {
	Name      string `json:"name"`
	Link      string `json:"link"`
	LinkSetup string `json:"linksetup"`
}

// function to get the language endpoints for
// the kind of app asked
func GetDatas(kind, lang string) (Lang, error) {
	datas := new(DATAS)
	datal := new(Lang)
	fmt.Println(messages.Fetching("https://raw.githubusercontent.com/aquemaati/gogol-api/main/data.json"))
	resp, err := http.Get("https://raw.githubusercontent.com/aquemaati/gogol-api/main/data.json")
	if err != nil {
		return *datal, err
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(datas); err != nil {
		return *datal, err
	}
	// Find the requested kind of data
	for _, d := range *datas {
		if d.Name == kind {
			// Find the requested language
			for _, l := range d.Lang {
				if l.Name == lang {
					return l, nil
				}
			}
		}
	}

	return *datal, fmt.Errorf("language %s for app %s not found", lang, kind)
}

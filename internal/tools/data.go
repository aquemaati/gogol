package tools

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

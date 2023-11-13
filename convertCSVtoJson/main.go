package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/tealeg/xlsx"
)

type Game struct {
	SrNo     string `json:"srNo"`
	GameName string `json:"gameName"`
	CNName   string `json:"cnName"`
	CHSName  string `json:"chsName"`
	GameType string `json:"gameType"`
	GameID   string `json:"gameID"`
}
type Result struct {
	Type   string   `json:"type"`
	Name   string   `json:"name"`
	Tag    []string `json:"tag"`
	Code   string   `json:"code"`
	Show   int      `json:"show"`
	NameCN string   `json:"nameCN"`
}

func main() {
	excelFileName := "PGSoft Game List_20231030 (Asia)-PG5.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	var games []Result
	for _, sheet := range xlFile.Sheets {
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				continue // Skip header row
			}
			var game Result
			for cellIndex, cell := range row.Cells {
				text := cell.String()
				switch cellIndex {
				case 0:
					game.Type = text
				case 1:
					game.Name = text
				case 2:
					game.NameCN = text
				case 4:
					game.Tag = append(game.Tag, "all")
					game.Tag = append(game.Tag, text)
				case 5:
					game.Code = text
				}
				game.Show = 2
			}
			games = append(games, game)
		}
	}

	jsonData, err := json.MarshalIndent(games, "", "    ")
	if err != nil {
		log.Fatal("Error marshaling JSON: ", err)
	}

	jsonFile, err := os.Create("games.json")
	if err != nil {
		log.Fatal("Error creating JSON file: ", err)
	}
	defer jsonFile.Close()

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		log.Fatal("Error writing JSON data to file: ", err)
	}
}

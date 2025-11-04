package services

import (
	"encoding/json"
	"os"

	"CrosswordBackend/model"
)

func JsonGenerator(){
	file, err := os.OpenFile("crosswordJSON.json", os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rows := 7
	columns := 7
	acrossClues := 5
	downClues := 5


	defaultCrossword := model.Crossword{}
	defaultCrossword.Rows = rows;
	defaultCrossword.Columns = columns;
	defaultCrossword.Clues.Across = make([]model.UnitClue, acrossClues)
	defaultCrossword.Clues.Down = make([]model.UnitClue, downClues)
	defaultCrossword.Grid = make([][]model.Cell, rows) 
	for i := range acrossClues{
		defaultCrossword.Clues.Across[i].ClueID = i + 1
		defaultCrossword.Clues.Across[i].ClueText = ""
	}

	for i := range downClues{
		defaultCrossword.Clues.Down[i].ClueID = i + 1
		defaultCrossword.Clues.Down[i].ClueText = ""
	}

	for i := range rows{
		defaultCrossword.Grid[i] = make([]model.Cell, columns)
		for j := range columns{
			defaultCrossword.Grid[i][j].IsBlank = false
			defaultCrossword.Grid[i][j].NumberAssociated = -1
		}
	}

	data, _ := json.Marshal(defaultCrossword)
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}
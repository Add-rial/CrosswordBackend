package services

import (
	"encoding/json"
	"os"

	"CrosswordBackend/model"
)

func CrosswordGenerator(){
	file, err := os.OpenFile("crosswordJSON.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rows := 7
	columns := 7
	acrossClues := 5
	downClues := 5


	templateCrossword := model.Crossword{}
	templateCrossword.Rows = rows;
	templateCrossword.Columns = columns;
	templateCrossword.Clues.Across = make([]model.UnitClue, acrossClues)
	templateCrossword.Clues.Down = make([]model.UnitClue, downClues)
	templateCrossword.Grid = make([][]model.Cell, rows) 
	for i := 0; i < acrossClues; i++{
		templateCrossword.Clues.Across[i].ClueID = i + 1
		templateCrossword.Clues.Across[i].ClueText = ""
	}

	for i := 0; i < downClues; i++{
		templateCrossword.Clues.Down[i].ClueID = i + 1
		templateCrossword.Clues.Down[i].ClueText = ""
	}

	for i := 0; i < rows; i++{
		templateCrossword.Grid[i] = make([]model.Cell, columns)
		for j := 0; j < columns; j++{
			templateCrossword.Grid[i][j].IsBlank = false
			templateCrossword.Grid[i][j].NumberAssociated = -1
		}
	}

	data, _ := json.Marshal(templateCrossword)
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}

func SolutionGenerator(){
	file, err := os.OpenFile("solutionJSON.json", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	clues := 10
	templateSolution := []model.UnitClue{}

	for i := 0; i < clues; i++{
		templateSolution[i].ClueID = i + 1
		templateSolution[i].ClueText = ""
	}

	data, _ := json.Marshal(templateSolution)
	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}
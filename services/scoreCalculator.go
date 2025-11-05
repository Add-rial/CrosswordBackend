package services

import (
	"CrosswordBackend/model"
	"encoding/json"
	"log"
	"os"
)

func LoadOfficialSolution() ([]model.UnitClue, int, error){
	file, err := os.ReadFile("solutionJSON.json")
	if err != nil {
		log.Println("Solutions haven't been uploaded")
	}

	var jsonExtacted model.CrosswordSolution
	err = json.Unmarshal(file, &jsonExtacted)
	if err != nil {
		log.Println("Error extracting soln:")
	}
	return jsonExtacted.Sol, jsonExtacted.Id, err
}

func CompareAnswer(userAns []model.UnitClue, sol []model.UnitClue) int{
	score := 0
	for k, clue := range sol{
		if(userAns[k].ClueText == clue.ClueText){
			score += 1
		}
	}
	return score
}

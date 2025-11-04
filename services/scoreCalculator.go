package services

import (
	"CrosswordBackend/model"
	"encoding/json"
	"log"
	"os"
)

func LoadOfficialSolution() ([]model.UnitClue, error){
	file, err := os.ReadFile("solutionJSON.json")
	if err != nil {
		log.Println("Solutions haven't been uploaded")
	}
	var sol []model.UnitClue
	err = json.Unmarshal(file, &sol)
	if err != nil {
		log.Println("Error extracting soln:")
	}
	return sol, err
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

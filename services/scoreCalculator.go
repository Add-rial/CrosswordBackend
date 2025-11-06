package services

import (
	"CrosswordBackend/model"
	"encoding/json"
	"log"
	"os"
	"strings"
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

func CompareAnswer(userAns []model.UnitClue, solMap map[int]string) int {
	score := 0

	for _, userClue := range userAns {
		if strings.EqualFold(strings.TrimSpace(userClue.ClueText),strings.TrimSpace(solMap[userClue.ClueID]),) {
			score++
		}
	}
	return score
}

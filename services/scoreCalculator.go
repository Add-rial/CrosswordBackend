package services

import (
	"CrosswordBackend/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func LoadOfficialSolution(crossword_id uint) ([]model.UnitClue, uint, error){
	filePath := fmt.Sprintf("data/day%d/solutionJSON.json", crossword_id)
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("Solutions haven't been uploaded")
		return nil, 0, err
	}

	var jsonExtacted model.CrosswordSolution
	err = json.Unmarshal(file, &jsonExtacted)
	if err != nil {
		log.Println("Error extracting soln:")
		return nil, 0, err
	}
	return jsonExtacted.Sol, jsonExtacted.Id, err
}

func CompareAnswer(userAns []model.UnitClue, solMap map[int]string) (int, bool) {
	score := 0
	allCorrect := true

	for _, userClue := range userAns {
		if strings.EqualFold(strings.TrimSpace(userClue.ClueText),strings.TrimSpace(solMap[userClue.ClueID]),) {
			score += len(strings.TrimSpace(solMap[userClue.ClueID])) 
		}else{
			allCorrect = false
		}
	}
	if len(userAns) != len(solMap){
		allCorrect = false
	}
	return score, allCorrect
}

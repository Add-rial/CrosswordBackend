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

func CompareAnswer(userAns []model.UnitClue, solMap map[int][]string) (int, bool) {
	score := 0
	allCorrect := true

	used := make(map[int][]bool)
	for id, answers := range solMap {
		used[id] = make([]bool, len(answers))
	}

	for _, userClue := range userAns {
		userText := strings.TrimSpace(userClue.ClueText)
		correctAnswers, exists := solMap[userClue.ClueID]
		if !exists {
			allCorrect = false
			continue
		}

		matchFound := false
		for i, sol := range correctAnswers {
			if !used[userClue.ClueID][i] &&
				strings.EqualFold(userText, strings.TrimSpace(sol)) {
				score += len(sol)
				used[userClue.ClueID][i] = true
				matchFound = true
				break
			}
		}

		if !matchFound {
			allCorrect = false
		}
	}
	
	for id, answers := range solMap {
		for i := range answers {
			if !used[id][i] {
				allCorrect = false
			}
		}
	}

	return score, allCorrect
}
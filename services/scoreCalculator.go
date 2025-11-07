package services

import (
	"CrosswordBackend/model"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

/* func LoadOfficialSolution(crossword_id uint) ([]model.UnitClue, uint, error){
	wd, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/data/day%d/solutionJSON.json", wd, crossword_id)
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
} */

// Replace your existing LoadOfficialSolution with this debug-friendly version.
func LoadOfficialSolution(crosswordID uint) ([]model.UnitClue, uint, error) {
    // 1) Log current working directory so we know where relative paths are resolved from
    if wd, err := os.Getwd(); err == nil {
        log.Println("LoadOfficialSolution - working dir:", wd)
    } else {
        log.Println("LoadOfficialSolution - could not get working dir:", err)
    }

    // 2) Log the requested ID
    log.Printf("LoadOfficialSolution - requested crosswordID: %d\n", crosswordID)

    // 3) Build file path and log it
    filePath := fmt.Sprintf("data/day%d/solutionJSON.json", crosswordID)
    log.Println("LoadOfficialSolution - looking for file:", filePath)

    // 4) Check existence & stat
    fi, err := os.Stat(filePath)
    if err != nil {
        if os.IsNotExist(err) {
            log.Println("LoadOfficialSolution - file does not exist:", filePath)
            return nil, 0, err
        }
        log.Println("LoadOfficialSolution - stat error:", err)
        return nil, 0, err
    }
    log.Printf("LoadOfficialSolution - file exists, size=%d\n", fi.Size())

    // 5) Read file
    fileBytes, err := os.ReadFile(filePath)
    if err != nil {
        log.Println("LoadOfficialSolution - read error:", err)
        return nil, 0, err
    }
    log.Printf("LoadOfficialSolution - read %d bytes\n", len(fileBytes))

    // 6) Quick sanity-print of a small prefix of the file so we can see content without huge logs
    if len(fileBytes) > 0 {
        preview := string(fileBytes)
        if len(preview) > 300 {
            preview = preview[:300] + "..."
        }
        log.Println("LoadOfficialSolution - file preview:", preview)
    } else {
        log.Println("LoadOfficialSolution - file is empty")
    }

    // 7) Unmarshal with clear error handling
    var jsonExtracted model.CrosswordSolution
    if err := json.Unmarshal(fileBytes, &jsonExtracted); err != nil {
        log.Println("LoadOfficialSolution - json unmarshal error:", err)
        return nil, 0, err
    }

    // 8) Final sanity checks on unmarshalled contents
    log.Printf("LoadOfficialSolution - parsed ID: %d, clues: %d\n", jsonExtracted.Id, len(jsonExtracted.Sol))
    return jsonExtracted.Sol, jsonExtracted.Id, nil
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

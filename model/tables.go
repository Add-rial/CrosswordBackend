package model

import (
    "database/sql/driver"
    "encoding/json"
    "log"
)

// @Description User entity with unique email and optional crossword answer
type User struct{
    ID uint `gorm:"primaryKey" json:"id" example:"1"`

    Email      string `gorm:"uniqueIndex;not null" example:"user@example.com"`
    Username       string `example:"JohnDoe"`
    Score int `gorm:"default:0;not null" example:"0"`
    DailyCrosswordAnswer CrosswordAnswer `gorm:"foreignKey:UserID" swaggerignore:"true" json:"-"`
}

// @Description A user's crossword answer submission
type CrosswordAnswer struct{
    ID uint `gorm:"primaryKey" json:"id,omitempty" swaggerignore:"true" example:"1"`
    
    UserID uint `json:"user_id,omitempty" swaggerignore:"true" gorm:"uniqueIndex:idx_user_crossword"`
    CrosswordID uint `json:"crossword_id" gorm:"uniqueIndex:idx_user_crossword" example:"1"`
    Answers Answers `gorm:"type:jsonb" json:"answers"`
    Scored bool `gorm:"default:false" json:"-" swaggerignore:"true"`
    TimeLeft int `json:"time_left" example:"137"`
}

// @Description A singular clue with its index
type UnitClue struct{
    ClueID int `example:"1"`
    ClueText string `example:"APPLE"`
}

type Answers []UnitClue

func (a Answers) Value() (driver.Value, error) {
    return json.Marshal(a)
}
func (a *Answers) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        log.Println("failed to scan Answers: value is not []byte")
    }
    return json.Unmarshal(bytes, a)
}

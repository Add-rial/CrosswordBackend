package model

import (
	"gorm.io/gorm"
)

// @Description User entity with unique email and optional crossword answer
type User struct{
    gorm.Model `swaggerignore:"true"`

    Email      string `gorm:"uniqueIndex;not null" example:"user@example.com"`
    Username       string `example:"JohnDoe"`
    Score int `gorm:"default:0;not null" example:"0"`
    DailyCrosswordAnswer CrosswordAnswer `gorm:"foreignKey:UserID"`
}

// @Description A user's crossword answer submission
type CrosswordAnswer struct{
    gorm.Model `swaggerignore:"true"`
    
    UserID uint 
    Answers []UnitClue `gorm:"type:jsonb" json:"answers"`
}

// @Description A singular clue with its index
type UnitClue struct{
    ClueID int `example:"1"`
    ClueText string `example:"APPLE"`
}
package model

import (
	"gorm.io/gorm"
)

// User represents an application user
// @Description User entity with unique email and optional crossword answer
type User struct{
    gorm.Model

    Email      string `gorm:"uniqueIndex;not null"`
    Username       string
    Score int `gorm:"default:0;not null"`
    DailyCrosswordAnswer CrosswordAnswer `gorm:"foreignKey:UserID"`
}

// CrosswordAnswer stores a user's crossword answers for the day
// @Description A user's crossword answer submission
type CrosswordAnswer struct{
    gorm.Model
    
    UserID uint 
    Answers []UnitClue `gorm:"type:jsonb" json:"answers"`
}

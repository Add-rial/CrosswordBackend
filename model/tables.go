package model

import (
	"gorm.io/gorm"
)

type User struct{
    gorm.Model

    Email      string `gorm:"uniqueIndex;not null"`
    Username       string
    Score int `gorm:"default:0;not null"`
    DailyCrosswordAnswer CrosswordAnswer `gorm:"foreignKey:UserID"`
}

type CrosswordAnswer struct{
    gorm.Model
    
    UserID uint 
    Answers []UnitClue `gorm:"type:jsonb" json:"answers`
}

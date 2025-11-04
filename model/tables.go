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
    Rows int
    Columns int
    Grid [][]byte `gorm:"type:json"`
}

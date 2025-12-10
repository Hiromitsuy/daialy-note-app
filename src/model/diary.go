package model

import (
	"time"

	"gorm.io/gorm"
)

type Diary struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	Note string
	UserID uint
	QuestionID uint
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
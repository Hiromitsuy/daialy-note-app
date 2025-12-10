package model

import (
	"time"

	"gorm.io/gorm"
)

type Diary struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	Note string `json:"note"`
	UserID uint `json:"userId"`
	QuestionID uint `json:"questionId"`
	Question Question `json:"question" gorm:"foreignKey:QuestionID"`
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
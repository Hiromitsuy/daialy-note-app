package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey,unique;autoIncrement"`
	Name string `json:"name" gorm:"unique;notNull"`
	Password []byte `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Diaries []Diary
	Questions []Question
}
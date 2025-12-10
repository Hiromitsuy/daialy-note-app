package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey,unique;autoIncrement"`
	Name string `json:"name" gorm:"unique;notNull"`
	Password string `json:"Password"`
	SaltWord string
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Diaries []Diary
}
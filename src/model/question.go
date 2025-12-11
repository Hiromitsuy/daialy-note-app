package model

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey;unique;autoIncrement"`
	QText string `json:"qtext"`
	UserID uint `json:"userId" gorm:"notNull"`
	Enabled bool `json:"enabled" gorm:"default:true"`
	CreatedAt time.Time `json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Diaries []Diary
}
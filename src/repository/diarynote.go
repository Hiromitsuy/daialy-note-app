package repository

import (
	"context"

	"github.com/lunasky-hy/dialy-note-app/src/model"
	"gorm.io/gorm"
)

type DiaryRepository struct {
	db *gorm.DB
	ctx context.Context
}

func (d DiaryRepository) QuestionCreate(question model.Question) error {
	err := gorm.G[model.Question](d.db).Create(d.ctx, &question)
	return err
}

func (d DiaryRepository) QuestionFind() ([]model.Question, error) {
	// questions, err := gorm.G[model.Question](d.db).Find(d.ctx)
	var questions []model.Question
	error := d.db.Find(&questions).Error
	return questions, error
}

func (d DiaryRepository) DiaryCreate(diary model.Diary) error {
	err := gorm.G[model.Diary](d.db).Create(d.ctx, &diary)
	return err
}

// func (d DiaryRepository) ReadDiaries(userId uint) ([]model.Diary, error) {
// 	diaries, error := gorm.G[model.User](d.db).Joins("Diary").Find(d.ctx)
// 	return diaries, error
// }

func CreateRepository(db *gorm.DB) DiaryRepository {
	context := context.Background()
	repos := DiaryRepository{db, context}
	return repos
}
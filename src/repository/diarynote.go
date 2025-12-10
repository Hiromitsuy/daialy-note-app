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

func (d DiaryRepository) QuestionsFind() ([]model.Question, error) {
	// questions, err := gorm.G[model.Question](d.db).Find(d.ctx)
	var questions []model.Question
	error := d.db.Find(&questions).Error
	return questions, error
}

func (d DiaryRepository) DiaryCreate(diary model.Diary) error {
	err := gorm.G[model.Diary](d.db).Create(d.ctx, &diary)
	return err
}

func (d DiaryRepository) DiariesFind(userId uint) ([]model.Diary, error) {
	var diaries []model.Diary
	error := d.db.Model(&model.Diary{}).Preload("Question").Where(&model.Diary{UserID: userId}).Find(&diaries).Error
	return diaries, error
}

func CreateRepository(db *gorm.DB) DiaryRepository {
	context := context.Background()
	repos := DiaryRepository{db, context}
	return repos
}
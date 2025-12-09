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

func (d DiaryRepository) CreateQuestion(question model.Question) error {
	err := gorm.G[model.Question](d.db).Create(d.ctx, &question)
	return err
}

func (d DiaryRepository) ReadQuestion() ([]model.Question, error) {
	questions, err := gorm.G[model.Question](d.db).Find(d.ctx)
	return questions, err
}

func CreateRepository(db *gorm.DB) DiaryRepository {
	context := context.Background()
	repos := DiaryRepository{db, context}
	return repos
}
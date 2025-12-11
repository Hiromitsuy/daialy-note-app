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
	var questions []model.Question
	error := d.db.Find(&questions).Error
	return questions, error
}

func (d DiaryRepository) QuestionsFindRand(lim int) ([]model.Question, error) {
	var questions []model.Question
	error := d.db.Order("RANDOM()").Limit(lim).Find(&questions).Error
	return questions, error
}

func (d DiaryRepository) QuestionsFindBy(q model.Question) ([]model.Question, error) {
	var questions []model.Question
	error := d.db.Where(&q).Find(&questions).Error
	return questions, error
}

func (d DiaryRepository) DiaryCreate(diary model.Diary) error {
	err := gorm.G[model.Diary](d.db).Create(d.ctx, &diary)
	return err
}

func (d DiaryRepository) DiaryFindById(id uint) (model.Diary, error) {
	var diaries model.Diary
	error := d.db.Model(&model.Diary{}).Where(&model.Diary{ID: id}).First(&diaries).Error
	return diaries, error
}

func (d DiaryRepository) DiariesFind(userId uint) ([]model.Diary, error) {
	var diaries []model.Diary
	error := d.db.Model(&model.Diary{}).Preload("Question").Where(&model.Diary{UserID: userId}).Find(&diaries).Error
	return diaries, error
}

func (d DiaryRepository) DiaryDelete(id uint) error {
	err := d.db.Model(&model.Diary{}).Delete(&model.Diary{ID: id}).Error
	return err
}

func (d DiaryRepository) UserCreate(user model.User) error {
	err := gorm.G[model.User](d.db).Create(d.ctx, &user)
	return err
}

func (d DiaryRepository) UserGet(username string) (model.User, error) {
	var user model.User
	err := d.db.Model(&model.User{}).Where("name = ?", username).First(&user).Error
	return user, err
}

func CreateRepository(db *gorm.DB) DiaryRepository {
	context := context.Background()
	repos := DiaryRepository{db, context}
	return repos
}
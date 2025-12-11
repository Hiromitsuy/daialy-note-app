package service

import (
	"errors"

	"github.com/lunasky-hy/dialy-note-app/src/model"
	"github.com/lunasky-hy/dialy-note-app/src/repository"
)

type DiaryService struct {
	repos repository.DiaryRepository
}

func (s DiaryService) Find(userId uint) ([]model.Diary, error) {
	diaries, error := s.repos.DiariesFind(userId)
	return diaries, error
}

func (s DiaryService) Create(d model.Diary, u model.User) (error) {
	newData := model.Diary{UserID: u.ID, Note: d.Note, QuestionID: d.QuestionID}
	error := s.repos.DiaryCreate(newData)
	return error
}

func (s DiaryService) Delete(diaryId uint, u model.User) (error) {
	targetItem, db_err := s.repos.DiaryFindById(diaryId)
	if db_err != nil {
		return db_err
	}
	if targetItem.UserID != u.ID {
		return errors.New("cannot access item")
	}

	del_err := s.repos.DiaryDelete(targetItem.ID)
	return del_err
}

func CreateDiaryService(repos repository.DiaryRepository) DiaryService {
	s := DiaryService{repos: repos}
	return s
}
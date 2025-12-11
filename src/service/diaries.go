package service

import (
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

func CreateDiaryService(repos repository.DiaryRepository) DiaryService {
	s := DiaryService{repos: repos}
	return s
}
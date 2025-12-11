package service

import (
	"github.com/lunasky-hy/dialy-note-app/src/model"
	"github.com/lunasky-hy/dialy-note-app/src/repository"
)

type QuestionService struct {
	repos repository.DiaryRepository
}

func (s QuestionService) Find() ([]model.Question, error) {
	questions, error := s.repos.QuestionsFindRand(3)
	return questions, error
}

func (s QuestionService) FindByUser(uid uint) ([]model.Question, error) {
	questions, error := s.repos.QuestionsFindBy(model.Question{UserID: uid})
	return questions, error
}

func (s QuestionService) Create(q model.Question, u model.User) (error) {
	q.UserID = u.ID;
	error := s.repos.QuestionCreate(q)
	return error
}

func CreateQuestonService(repos repository.DiaryRepository) QuestionService {
	s := QuestionService{repos: repos}
	return s
}
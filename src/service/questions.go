package service

import (
	"github.com/lunasky-hy/dialy-note-app/src/model"
	"github.com/lunasky-hy/dialy-note-app/src/repository"
)

type QuestionService struct {
	repos repository.DiaryRepository
}

func (s QuestionService) Find() ([]model.Question, error) {
	questions, error := s.repos.QuestionFind()
	return questions, error
}

func (s QuestionService) Create(q model.Question) (error) {
	error := s.repos.QuestionCreate(q)
	return error
}

func CreateQuestonService(repos repository.DiaryRepository) QuestionService {
	s := QuestionService{repos: repos}
	return s
}
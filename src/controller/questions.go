package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunasky-hy/dialy-note-app/src/model"
	"github.com/lunasky-hy/dialy-note-app/src/repository"
)

type QuestionController struct {
	repos repository.DiaryRepository
}

func (qc QuestionController) Get(c *gin.Context) {
	ques, _ := qc.repos.ReadQuestion()
	c.JSON(http.StatusOK, ques)
}

func (qc QuestionController) Post(c *gin.Context) {
	var json model.Question
	if err := c.ShouldBindJSON(&json); err != nil {
		return
	}
	fmt.Println(json.QText);
	qc.repos.CreateQuestion(json);
	c.String(http.StatusAccepted, `sended`);
}

func CreateQuestionController(repos repository.DiaryRepository) QuestionController {
	controller := QuestionController{repos: repos}

	return controller
}
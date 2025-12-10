package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunasky-hy/dialy-note-app/src/model"
	"github.com/lunasky-hy/dialy-note-app/src/service"
)

type QuestionController struct {
	service service.QuestionService
}

func (qc QuestionController) Get(c *gin.Context) {
	ques, _ := qc.service.Find()
	c.JSON(http.StatusOK, ques)
}

func (qc QuestionController) Post(c *gin.Context) {
	var json model.Question
	if err := c.ShouldBindJSON(&json); err != nil {
		return
	}
	fmt.Println(json.QText);
	qc.service.Create(json);
	c.String(http.StatusAccepted, `sended`);
}

func CreateQuestionController(service service.QuestionService) QuestionController {
	controller := QuestionController{service: service}
	return controller
}
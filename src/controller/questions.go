package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunasky-hy/dialy-note-app/src/authorization"
	"github.com/lunasky-hy/dialy-note-app/src/model"
	"github.com/lunasky-hy/dialy-note-app/src/service"
)

type QuestionController struct {
	service service.QuestionService
	authHandler authorization.AuthHandler
}

func (qc QuestionController) Get(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization");
	user := model.User{}
	if auth_err := qc.authHandler.VerifyJwt(authHeader, &user); auth_err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": auth_err})
		return
	}

	isMine := c.Query("mine")
	if isMine == "true" {
		ques, _ := qc.service.FindByUser(user.ID)
		c.JSON(http.StatusOK, ques)
	} else {
		ques, _ := qc.service.Find()
		c.JSON(http.StatusOK, ques)
	}
}

func (qc QuestionController) Post(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization");
	user := model.User{}
	if auth_err := qc.authHandler.VerifyJwt(authHeader, &user); auth_err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": auth_err})
		return
	}

	var json model.Question
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	qc.service.Create(json, user);
	c.String(http.StatusAccepted, `sended`);
}

func CreateQuestionController(service service.QuestionService, authHandler authorization.AuthHandler) QuestionController {
	controller := QuestionController{service: service, authHandler: authHandler}
	return controller
}
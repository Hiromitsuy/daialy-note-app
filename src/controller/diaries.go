package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunasky-hy/dialy-note-app/src/authorization"
	"github.com/lunasky-hy/dialy-note-app/src/model"
	"github.com/lunasky-hy/dialy-note-app/src/service"
)

type DiariesController struct {
	service service.DiaryService
	authHandler authorization.AuthHandler
}

func (qc DiariesController) Get(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization");
	user := model.User{}
	if auth_err := qc.authHandler.VerifyJwt(authHeader, &user); auth_err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": auth_err})
		return
	}

	ques, _ := qc.service.Find(user.ID)
	c.JSON(http.StatusOK, ques)
}

func (qc DiariesController) Post(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization");
	user := model.User{}
	if auth_err := qc.authHandler.VerifyJwt(authHeader, &user); auth_err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": auth_err})
		return
	}

	var json model.Diary
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	qc.service.Create(json, user);
	c.String(http.StatusAccepted, `sended`);
}

func CreateDiaryController(service service.DiaryService, authHandler authorization.AuthHandler) DiariesController {
	controller := DiariesController{service: service, authHandler: authHandler}
	return controller
}
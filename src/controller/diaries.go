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
	auth_err := qc.authHandler.VerifyJwt(authHeader, &user);
	if auth_err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": auth_err})
	}

	ques, _ := qc.service.Find(user.ID)
	c.JSON(http.StatusOK, ques)
}

func (qc DiariesController) Post(c *gin.Context) {
	var json model.Diary
	if err := c.ShouldBindJSON(&json); err != nil {
		return
	}
	qc.service.Create(json);
	c.String(http.StatusAccepted, `sended`);
}

func CreateDiaryController(service service.DiaryService, authHandler authorization.AuthHandler) DiariesController {
	controller := DiariesController{service: service, authHandler: authHandler}
	return controller
}
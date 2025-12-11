package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunasky-hy/dialy-note-app/src/model"
	"github.com/lunasky-hy/dialy-note-app/src/service"
)

type AuthController struct {
	service service.AuthService
}

type RequestSignup struct {
	Name string `json:"name"`
	Password string `json:"password"`
}

func (ac AuthController) Signup(c *gin.Context) {
	var json RequestSignup
	if parse_err := c.ShouldBindJSON(&json); parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": parse_err.Error()})
		return
	}

	user := model.User{Name: json.Name, Password: []byte(json.Password)}
	token, register_err := ac.service.Register(user)
	
	if register_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": register_err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}

func (ac AuthController) Signin(c *gin.Context) {
	var json RequestSignup
	if parse_err := c.ShouldBindJSON(&json); parse_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": parse_err.Error()})
		return
	}

	user := model.User{Name: json.Name, Password: []byte(json.Password)}
	token, auth_err := ac.service.AuthorizeUser(user)
	
	if auth_err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": auth_err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"token": token})
}


func CreateAuthController(service service.AuthService) AuthController {
	controller := AuthController{service: service}
	return controller
}
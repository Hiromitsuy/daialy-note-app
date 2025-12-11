package main

import (
	"log"

	"github.com/lunasky-hy/dialy-note-app/src/authorization"
	"github.com/lunasky-hy/dialy-note-app/src/controller"
	"github.com/lunasky-hy/dialy-note-app/src/database"
	"github.com/lunasky-hy/dialy-note-app/src/repository"
	"github.com/lunasky-hy/dialy-note-app/src/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	enverr := godotenv.Load()
	if enverr != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.ConnectPostgres()
	repos := repository.CreateRepository(db)

	authHandler := authorization.CreateAuthHandler(repos)

	authService := service.CreateAuthService(repos)
	authController := controller.CreateAuthController(authService)

	questionService := service.CreateQuestonService(repos)
	questionController := controller.CreateQuestionController(questionService, authHandler)

	diaryService := service.CreateDiaryService(repos)
	diaryController := controller.CreateDiaryController(diaryService, authHandler)



	// loggerとrecoveryミドルウェア付きGinルーター作成
	r := gin.Default()

	{
		v1 := r.Group("/v1")

		v1.GET("/api/questions", questionController.Get)
		v1.POST("/api/questions", questionController.Post)

		v1.GET("/api/diaries", diaryController.Get)
		v1.POST("/api/diaries", diaryController.Post)

		v1.POST("/api/auth/signup", authController.Signup)
		v1.POST("/api/auth/signin", authController.Signin)
	}

	// ポート8080でサーバー起動（デフォルト）
	// 0.0.0.0:8080（Windowsではlocalhost:8080）で待機
	sverr := r.Run()
	if sverr != nil {
		return
	}
}

package main

import (
	"log"
	"os"
	"path"
	"path/filepath"

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
		v1.DELETE("/api/diaries/:postId", diaryController.Delete)

		v1.POST("/api/auth/signup", authController.Signup)
		v1.POST("/api/auth/signin", authController.Signin)
	}

	r.NoRoute(func(c *gin.Context) {
		_, file := path.Split(c.Request.RequestURI)
		ext := filepath.Ext(file)
		if file == "" || ext == "" {
			c.File("./frontend/dist" + "/index.html")
		} else {
			c.File("./frontend/dist" + c.Request.RequestURI)
		}
	})

	// ポート8080でサーバー起動（デフォルト）
	// 0.0.0.0:8080（Windowsではlocalhost:8080）で待機
	port := os.Getenv("PORT")
	sverr := r.Run(`localhost:` + port)
	if sverr != nil {
		return
	}
}

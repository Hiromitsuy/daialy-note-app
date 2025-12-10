package main

import (
	"log"
	"net/http"
	"time"

	"github.com/lunasky-hy/dialy-note-app/src/controller"
	"github.com/lunasky-hy/dialy-note-app/src/database"
	"github.com/lunasky-hy/dialy-note-app/src/repository"
	"github.com/lunasky-hy/dialy-note-app/src/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type question struct {
	Id int `json:"id"`
	QText string `json:"qtext"`
}

type daiary struct {
	Id int `json:"id"`
	Note string `json:"note"`
	UserId int `json:"userId"`
	Question question `json:"question"`
	CreatedAt time.Time `json:"createdAt"`
}

var mockQuestions = []question{
	{Id: 1, QText: "今の気分は？"},
	{Id: 2, QText: "今日の夕食は？"},
	{Id: 3, QText: "〇〇でいい感じ！"},
}

var mockDiaries = []daiary{
	{Id: 1, Note: "最高!", UserId: 1, Question: mockQuestions[0], CreatedAt: time.Now()},
	{Id: 2, Note: "ぼちぼち", UserId: 1, Question: mockQuestions[0], CreatedAt: time.Now()},
}

func main() {
	enverr := godotenv.Load()
	if enverr != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.ConnectPostgres()
	repos := repository.CreateRepository(db)
	
	questionService := service.CreateQuestonService(repos)
	questionController := controller.CreateQuestionController(questionService)

	// loggerとrecoveryミドルウェア付きGinルーター作成
	r := gin.Default()

	{
		v1 := r.Group("/v1")

		v1.GET("/api/questions", questionController.Get)
		v1.POST("/api/questions", questionController.Post)

		v1.GET("/api/diaries", func(c *gin.Context) {
			c.JSON(http.StatusOK, mockDiaries)
		})
		v1.POST("/api/diaries", func(c *gin.Context) {
			c.String(http.StatusAccepted, `sended`);
		})
	}

	// ポート8080でサーバー起動（デフォルト）
	// 0.0.0.0:8080（Windowsではlocalhost:8080）で待機
	sverr := r.Run()
	if sverr != nil {
		return
	}
}

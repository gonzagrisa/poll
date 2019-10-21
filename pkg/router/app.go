package router

import (
	"github.com/gin-gonic/gin"
	"poll/pkg/controllers"
)

func Start() {
	router := gin.Default()

	// Polls
	router.POST("poll", controllers.PostPoll)
	router.GET("poll", controllers.GetPoll)
	router.DELETE("poll", controllers.DeletePoll)

	// Answers
	router.PUT("poll/:answer", controllers.PutAnswer)
	router.GET("poll/results", controllers.GetResults)

	err := router.Run("localhost:8080")
	if err != nil {
		panic(err)
	}
}

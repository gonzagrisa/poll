package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"poll/pkg/domain"
	"poll/pkg/services"
)

func PostPoll(c *gin.Context) {
	var poll domain.Poll

	err := c.BindJSON(&poll)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	poll, err = services.CreatePoll(poll)
	if err != nil {
		apiErr := ParseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, poll)
}

func GetPoll(c *gin.Context) {
	poll, err := services.GetPoll()

	if err != nil {
		apiErr := ParseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, poll)
}

func PutAnswer(c *gin.Context) {
	answer := c.Param("answer")

	err := services.AnswerPoll(answer)
	if err != nil {
		apiErr := ParseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	results, err := services.GetResults()
	if err != nil {
		apiErr := ParseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, results)
}

func GetResults(c *gin.Context) {
	results, err := services.GetResults()
	if err != nil {
		apiErr := ParseError(err)
		c.String(apiErr.Status, apiErr.Error())
		return
	}

	c.JSON(http.StatusOK, results)
}

func DeletePoll(c *gin.Context) {
	services.DeletePoll()

	c.Status(http.StatusOK)
}

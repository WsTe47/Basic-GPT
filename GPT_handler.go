package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"net/http"
)

/*
r := gin.Default()

	r.GET("/GPTChat", func(c *gin.Context) {
		api.GPTChat(c, gptClient)
	})
*/
func GPTChat(c *gin.Context, gptClient *Client) {
	question := c.Query("question")
	log.Info(question)
	answer, err := gptClient.SendMsg(question)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	log.Info(answer)
	c.JSON(http.StatusOK, answer)

}

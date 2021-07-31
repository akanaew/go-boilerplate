package handlers

import "github.com/gin-gonic/gin"

type WelcomeResponse struct {
	Message string `json:"message"`
}

// @Summary Welcome
// @Description Default welcome endpoint
// @Tag Default
// @ID welcome-get
// @Success 200 {object} WelcomeResponse
// @Router / [get]
func (h Handler) welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}

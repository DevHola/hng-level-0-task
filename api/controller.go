package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Email string `json:"email"`
	Current_datetime string `json:"current_datetime"`
	Github_url string `json:"github_url"`
}

func Index(c *gin.Context) {
	response := Response{
		Email: "connectola@yahoo.com",
		Current_datetime: time.Now().UTC().Format(time.RFC3339),
		Github_url: "https://github.com/DevHola/hng-level-0-task.git",
	}

	c.JSON(200, response)
}
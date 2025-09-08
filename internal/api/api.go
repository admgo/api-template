package api

import (
	"api-template/internal/handler"

	"github.com/gin-gonic/gin"
)

type Api struct {
}

func (h Api) CreateUser(c *gin.Context, userID int64, params handler.CreateUserParams) {
	c.JSON(200, gin.H{
		"user_id": userID,
	})
}

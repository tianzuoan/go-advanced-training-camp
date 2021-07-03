package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"week02/global"
	"week02/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	return Service{
		ctx: ctx,
		dao: dao.New(global.DBEngine),
	}
}

func PingService(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetUsersService(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

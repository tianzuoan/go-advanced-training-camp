package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	geekerror "week02/error"
	"week02/error/errorcode"
	"week02/internal/service"
)

func GetUserList(c *gin.Context) {
	srv := service.New(c)
	name := c.Query("name")
	users, err := srv.GetUserList(name, 0, 10)
	if err != nil {
		c.JSON(200, gin.H{
			"data": users,
			"err":  geekerror.NewGeekError(errorcode.FAIL).SetUserMsg("%s", err.Error()),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": users,
		"err":  geekerror.NewGeekError(errorcode.SUCCESS),
	})
	return
}

func GetUserDetail(c *gin.Context) {
	srv := service.New(c)
	name := c.Query("name")
	user, err := srv.GetUserDetail(name)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(200, gin.H{
			"data": user,
			"err":  geekerror.CodeNotFound,
		})
		return
	}
	if err != nil {
		c.JSON(200, gin.H{
			"data": user,
			"err":  geekerror.NewGeekError(errorcode.FAIL).SetUserMsg("%s", err.Error()),
		})
		return
	}
	c.JSON(200, gin.H{
		"data": user,
		"err":  geekerror.NewGeekError(errorcode.SUCCESS),
	})
	return
}

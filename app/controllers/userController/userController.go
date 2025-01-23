package usercontroller

import (
	"EBUSU/app/apiException"
	"EBUSU/app/service/userService"
	"EBUSU/app/utils"

	"github.com/gin-gonic/gin"
)

func GetQrcode(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	qrcode, err := userService.GetQrcode(token)
	if err != nil {
		c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{"qrcode": qrcode})
}

func GetNotice(c *gin.Context) {
	token := c.GetHeader("Authorization")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	numPages := c.DefaultQuery("numPages", "")

	if token == "" {
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	notice, err := userService.GetNotice(token, page, pageSize, numPages)
	if err != nil {
		c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{"notice": notice})
}

func MarkReadedNotice(c *gin.Context) {
	token := c.GetHeader("Authorization")
	noticeID := c.Param("noticeID")

	if token == "" {
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	err := userService.MarkReaded(token, noticeID)
	if err != nil {
		c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func GetUnreadCount(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	count, err := userService.GetUnreadCount(token)
	if err != nil {
		c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{"count": count})
}

func CheckTokenAlive(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token == "" {
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	err := userService.CheckTokenAlive(token)
	if err != nil {
		c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{"alive": true})
}

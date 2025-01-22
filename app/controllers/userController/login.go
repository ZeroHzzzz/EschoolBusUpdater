package usercontroller

import (
	"EBUSU/app/apiException"
	"EBUSU/app/service/userService"
	"EBUSU/app/utils"

	"github.com/gin-gonic/gin"
)

// LoginRequest 手机登录请求结构体
type LoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// LoginByYxyRequest 易校园登录请求结构体
type LoginByYxyRequest struct {
	UnionID string `json:"unionID"`
}

// LoginByPhone 手机号登录处理函数
func LoginByPhone(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	token, err := userService.LoginByPhone(req.Phone, req.Password)
	if err != nil {
		c.AbortWithError(200, apiException.NoThatPasswordOrWrong)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"token": token,
	})
}

// LoginByYxy 易校园登录处理函数
func LoginByYxy(c *gin.Context) {
	var req LoginByYxyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	token, err := userService.LoginByYxy(req.UnionID)
	if err != nil {
		c.AbortWithError(200, apiException.NoThatPasswordOrWrong)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"token": token,
	})
}

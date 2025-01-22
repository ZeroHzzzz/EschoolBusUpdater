package updatercontroller

import (
	"EBUSU/app/service/updateService"
	"EBUSU/app/utils"
	"EBUSU/config/config"

	"github.com/gin-gonic/gin"
)

func GetUpdateStatus(c *gin.Context) {
	status, err := updateService.GetBusUpdateStatus()
	if err != nil {
		c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, status)
}

func UpdateBusInfo(c *gin.Context) {
	authToken := config.Config.GetString("eBus.token")
	// fmt.Println("authToken", authToken)
	err := updateService.BusInfoUpdater(authToken)
	if err != nil {
		c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

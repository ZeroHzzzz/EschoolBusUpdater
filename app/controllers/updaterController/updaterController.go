package updatercontroller

import (
	"EBUSU/app/service/updateService"
	"EBUSU/app/utils"
	constants "EBUSU/app/utils/const"

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
	// fmt.Println("authToken", authToken)
	err := updateService.BusInfoUpdater(constants.SystemUid)
	if err != nil {
		c.AbortWithError(200, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

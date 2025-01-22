package buscontroller

import (
	"EBUSU/app/apiException"
	"EBUSU/app/service/busService"
	"EBUSU/app/utils"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBusInfo(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	search := c.DefaultQuery("search", "")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		log.Printf("Invalid page parameter: %v", err)
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt <= 0 {
		log.Printf("Invalid pageSize parameter: %v", err)
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	busInfoList, err := busService.GetBusInfo(search)
	if err != nil {
		log.Printf("Error: failed to get bus info: %v", err)
		c.AbortWithError(200, apiException.NotFound)
		return
	}

	start := (pageInt - 1) * pageSizeInt
	end := pageInt * pageSizeInt

	utils.JsonSuccessResponse(c, busInfoList[max(start, 0):min(end, len(busInfoList))])
	// fmt.Println(busInfoList[start:end])
}

func GetBusRecords(c *gin.Context) {
	token := c.GetHeader("Authorization")
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	status := c.DefaultQuery("status", "")

	if token == "" {
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		log.Printf("Invalid page parameter: %v", err)
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt <= 0 {
		log.Printf("Invalid pageSize parameter: %v", err)
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	busRecords, err := busService.FetchBusRecords(token, page, pageSize, status)
	if err != nil {
		c.AbortWithError(200, err)
		return
	}

	start := (pageInt - 1) * pageSizeInt
	end := pageInt * pageSizeInt

	utils.JsonSuccessResponse(c, busRecords[max(start, 0):min(end, len(busRecords))])
}

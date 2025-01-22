package router

import (
	buscontroller "EBUSU/app/controllers/busController"
	updatercontroller "EBUSU/app/controllers/updaterController"
	usercontroller "EBUSU/app/controllers/userController"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	const pre = "/api"
	api := r.Group(pre)

	initBusRoutes(api)
	initUserRoutes(api)
	initUpdaterRoutes(api)

	for _, route := range r.Routes() {
		fmt.Println(route.Method, route.Path)
	}
}

func initBusRoutes(api *gin.RouterGroup) {
	bus := api.Group("/bus")
	{
		bus.GET("/info", buscontroller.GetBusInfo)
		bus.GET("/records", buscontroller.GetBusRecords)
	}
}

func initUserRoutes(api *gin.RouterGroup) {
	user := api.Group("/user")
	{
		// Auth routes
		user.POST("/login/yxy", usercontroller.LoginByYxy)
		user.POST("/login/phone", usercontroller.LoginByPhone)
		user.GET("/checkalive", usercontroller.CheckTokenAlive)

		// Notice related routes
		user.GET("/notice", usercontroller.GetNotice)
		user.PATCH("/notice/:noticeID", usercontroller.MarkReadedNotice)
		user.GET("/notice/unread", usercontroller.GetUnreadCount)

		// Other routes
		user.GET("/qrcode", usercontroller.GetQrcode)
	}
}

func initUpdaterRoutes(api *gin.RouterGroup) {
	updater := api.Group("/updater")
	{
		updater.GET("/status", updatercontroller.GetUpdateStatus)
		updater.POST("/update", updatercontroller.UpdateBusInfo)
	}
}

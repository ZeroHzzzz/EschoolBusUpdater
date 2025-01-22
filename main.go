package main

import (
	"EBUSU/app/midwares"
	"EBUSU/app/service/updateService"
	"EBUSU/config/config"
	"EBUSU/config/router"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	authToken := config.Config.GetString("eBus.token")
	interval := config.Config.GetDuration("timer.interval") * time.Hour
	port := config.Config.GetInt("server.port")

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	router.Init(r)

	go updateService.Run(authToken, interval)
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("ServerStartFailed", err)
	}
	// busService.GetCallback()
	// fmt.Println(userService.LoginByYxy("2408157831570432101"))
	// userService.LoginByYxy("2408157831570432101")
	// busService.FetchBusRecords("e0ae2516deebfdd720db927e9a268a6665328443", "1", "999", "30")
	// fmt.Println(busRecords)
}

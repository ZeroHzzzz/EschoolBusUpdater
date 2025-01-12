package main

import (
	"EBUSU/app/service/userService"
	"EBUSU/config/api"
	"fmt"
	"log"
)


func main() {
	data, err := userService.GetUnreadCount(api.SystemToken)
	if err != nil {
		log.Fatalf("Failed to get bus info: %v\n", err)
	}
	// fmt.Println(len(data))
	// for i := 0; i < len(data); i++ {
	// 	fmt.Println(data[i].ID)
	// }
	// fmt.Println(data[0].ID)
	fmt.Println(data)
}
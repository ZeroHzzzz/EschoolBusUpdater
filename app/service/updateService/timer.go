package updateService

import (
	"log"
	"time"
)

func Run(unionID string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				err := BusInfoUpdater(unionID)
				if err != nil {
					log.Printf("Error updating bus info: %v", err)
				}
			case <-quit:
				return
			}
		}
	}()

	// 防止主程序退出
	select {}
}

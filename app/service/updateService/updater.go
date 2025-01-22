package updateService

import (
	"EBUSU/app/models"
	"EBUSU/app/service/busService"
	"EBUSU/config/redis"
	"encoding/json"
	"log"
	"time"
)

func GetBusUpdateStatus() (models.BusStatus, error) {
	status, err := redis.RedisClient.Get("LastUpdateStatus").Result()
	if err != nil {
		return models.BusStatus{}, err
	}
	time, err := redis.RedisClient.Get("LastUpdateTime").Result()
	if err != nil {
		return models.BusStatus{}, err
	}

	return models.BusStatus{
		Status:       status,
		LastUpdateAt: time,
	}, nil
}

func BusInfoUpdater(authToken string) error {
	err := redis.RedisClient.Set("LastUpdateTime", time.Now().Format(time.RFC3339), 0).Err()
	if err != nil {
		return err
	}

	const (
		defaultPage     = 1
		defaultPageSize = 999
		shuttle_type    = -10
	)

	log.Printf("Starting bus info update with page=%d, pageSize=%d", defaultPage, defaultPageSize)

	// 获取 bus 信息
	busInfoList, err := busService.FetchBusInfo(authToken, defaultPage, defaultPageSize)
	if err != nil {
		log.Printf("Error: failed to get bus info: %v", err)
		redis.RedisClient.Set("LastUpdateStatus", "Wrong", 0)
		return err
	}

	for i := range busInfoList {
		timeList, timeErr := busService.FetchBusTime(authToken, busInfoList[i].ID, shuttle_type)
		if timeErr != nil {
			log.Printf("Error: failed to get bus time for ID %s: %v", busInfoList[i].ID, timeErr)
			redis.RedisClient.Set("LastUpdateStatus", "Wrong", 0)
			continue
		}
		busInfoList[i].BusTimeList = timeList
		// busInfoData, err := json.Marshal(busInfoList[i])
		// if err != nil {
		// 	return err
		// }
		// err = redis.RedisClient.HSet("BusInfo:"+busInfoList[i].ID, "info", busInfoData).Err()
		// if err != nil {
		// 	return err
		// }
	}

	// 更新redis
	// err = redis.RedisClient.HSet("BusInfo:"+)
	err = redis.RedisClient.Del("BusInfo").Err()
	if err != nil {
		redis.RedisClient.Set("LastUpdateStatus", "Wrong", 0)
		return err
	}

	for _, busInfo := range busInfoList {
		busInfoData, err := json.Marshal(busInfo)
		if err != nil {
			redis.RedisClient.Set("LastUpdateStatus", "Wrong", 0)
			return err
		}
		err = redis.RedisClient.RPush("BusInfo", busInfoData).Err()
		if err != nil {
			redis.RedisClient.Set("LastUpdateStatus", "Wrong", 0)
			return err
		}
	}

	err = redis.RedisClient.Set("LastUpdateStatus", "Success", 0).Err()
	if err != nil {
		return err
	}
	return nil
}

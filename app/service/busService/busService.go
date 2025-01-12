package busService

import (
	"EBUSU/app/fetch"
	"EBUSU/config/api"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// GetBusInfo 获取公交信息
func GetBusInfo(authToken string, page int, pageSize int) ([]BusInfo, error) {
	url := api.EBusHost + string(api.BusInfo)

	// 发起请求
	resp, err := fetch.Client.R().
		SetQueryParams(map[string]string{
			"page":     strconv.Itoa(page),
			"page_size": strconv.Itoa(pageSize),
		}).
		SetHeader("Authorization", authToken).
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0").
		SetHeader("Accept", "*/*").
		SetHeader("Referer", "https://h5.pinbayun.com/").
		SetHeader("Origin", "https://h5.pinbayun.com").
		Get(url)

	// 错误处理：请求失败
	if err != nil {
		log.Printf("Error sending request to %s: %v\n", url, err)
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	// 错误处理：HTTP 状态码检查
	if resp.StatusCode() != 200 {
		log.Printf("Received non-OK HTTP status from %s: %d\n", url, resp.StatusCode())
		return nil, fmt.Errorf("received non-OK HTTP status: %d", resp.StatusCode())
	}

	// 响应内容，并提取出result字段
	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		log.Printf("Error unmarshaling JSON response: %v\n", err)
		return nil, fmt.Errorf("error unmarshaling JSON response: %w", err)
	}

	// 提取 results 字段
	results, ok := result["results"].([]interface{})
	if !ok {
		log.Printf("Error: 'results' field not found or is not an array.\n")
		return nil, fmt.Errorf("'results' field not found or is not an array")
	}

	// 将 results 转换为 []BusInfo
	resultBytes, err := json.Marshal(results)
	if err != nil {
		log.Printf("Error marshaling 'results' to JSON: %v\n", err)
		return nil, fmt.Errorf("error marshaling 'results' to JSON: %w", err)
	}

	// 解析为 BusInfo 结构体切片
	var busInfo []BusInfo
	err = json.Unmarshal(resultBytes, &busInfo)
	if err != nil {
		log.Printf("Error unmarshaling 'results' into []BusInfo: %v\n", err)
		return nil, fmt.Errorf("error unmarshaling 'results' into []BusInfo: %w", err)
	}

	// 返回解析后的数据
	return busInfo, nil
}

func GetBusTime(authToken, busID string, shuttle_type int) ([]BusTime, error) {
	url := api.EBusHost + string(api.BusTime)
	url = strings.Replace(string(url), "{id}", busID, 1)

	resp, err := fetch.Client.R().
		SetQueryParams(map[string]string{
			"shuttle_type":     "-10",
		}).
		SetHeader("Authorization", authToken).
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0").
		SetHeader("Accept", "*/*").
		SetHeader("Referer", "https://h5.pinbayun.com/").
		SetHeader("Origin", "https://h5.pinbayun.com").
		Get(url)

	if err != nil {
		log.Printf("Error sending request to %s: %v\n", url, err)
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	
	var busTimes []BusTime
	err = json.Unmarshal(resp.Body(), &busTimes)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	return busTimes, nil
}
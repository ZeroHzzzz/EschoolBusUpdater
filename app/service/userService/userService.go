package userService

import (
	"EBUSU/app/fetch"
	"EBUSU/app/service/busService"
	"EBUSU/config/api"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type UnreadCount struct {
	Count int `json:"msg_count"`
}

func GetUnreadCount(authToken string) (int, error) {
	url := api.EBusHost + string(api.UserUnreadCount)
	resp, err := fetch.Client.R().
		SetHeader("Authorization", authToken).
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0").
		SetHeader("Accept", "*/*").
		SetHeader("Referer", "https://h5.pinbayun.com/").
		SetHeader("Origin", "https://h5.pinbayun.com").
		Get(url)

	if err != nil {
		log.Printf("Error sending request to %s: %v\n", url, err)
		return 0, fmt.Errorf("failed to send request: %w", err)
	}
	
	var unreadCount UnreadCount
	err = json.Unmarshal(resp.Body(), &unreadCount)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
		return 0, nil
	}

	return unreadCount.Count, nil
}

func CheckTokenAlive(authToken string) (bool, error) {
	// 利用未读信息接口做Token存活性检测
	url := api.EBusHost + string(api.UserUnreadCount)
	resp, err := fetch.Client.R().
		SetHeader("Authorization", authToken).
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0").
		SetHeader("Accept", "*/*").
		SetHeader("Referer", "https://h5.pinbayun.com/").
		SetHeader("Origin", "https://h5.pinbayun.com").
		Get(url)

	if err != nil {
		log.Printf("Error sending request to %s: %v\n", url, err)
		return false, fmt.Errorf("failed to send request: %w", err)
	}
	
	if resp.StatusCode() != 200 {
		log.Printf("Received non-OK HTTP status from %s: %d\n", url, resp.StatusCode())
		return false, fmt.Errorf("received non-OK HTTP status: %d", resp.StatusCode())
	}

	return true, nil
}

func GetBookRecords(authToken, page, pageSize, status string) ([]busService.BusInfo, error) {
	// status参数为10或者20返回未开始预约信息，为30或者40返回已结束的成功预定信息，为0返回所有预定记录（包括已经取消）
	url := api.EBusHost + string(api.BusInfo)

	// 发起请求
	resp, err := fetch.Client.R().
		SetQueryParams(map[string]string{
			"page":     page,
			"page_size": pageSize,
			"status": status,
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
	var busInfo []busService.BusInfo
	err = json.Unmarshal(resultBytes, &busInfo)
	if err != nil {
		log.Printf("Error unmarshaling 'results' into []BusInfo: %v\n", err)
		return nil, fmt.Errorf("error unmarshaling 'results' into []BusInfo: %w", err)
	}

	// 返回解析后的数据
	return busInfo, nil
}

type QrcodeString struct {
	QrcodeString string `json:"qrcode"`
}
func GetQrcode(authToken string) (string, error){
	url := api.EBusHost + string(api.UserQrcode)
	resp, err := fetch.Client.R().
		SetHeader("Authorization", authToken).
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0").
		SetHeader("Accept", "*/*").
		SetHeader("Referer", "https://h5.pinbayun.com/").
		SetHeader("Origin", "https://h5.pinbayun.com").
		Get(url)

	if err != nil {
		log.Printf("Error sending request to %s: %v\n", url, err)
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	
	var qrcodeString QrcodeString
	err = json.Unmarshal(resp.Body(), &qrcodeString)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
		return "", nil
	}

	return qrcodeString.QrcodeString, nil
}


func GetNotice(authToken, page, pageSize, numPages string) ([]Notice, error) {
	// status参数为10或者20返回未开始预约信息，为30或者40返回已结束的成功预定信息，为0返回所有预定记录（包括已经取消）
	url := api.EBusHost + string(api.UserNotice)

	// 发起请求
	resp, err := fetch.Client.R().
		SetQueryParams(map[string]string{
			"page":     page,
			"page_size": pageSize,
			"num_pages": numPages,
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

	resultBytes, err := json.Marshal(results)
	if err != nil {
		log.Printf("Error marshaling 'results' to JSON: %v\n", err)
		return nil, fmt.Errorf("error marshaling 'results' to JSON: %w", err)
	}

	// 解析为 BusInfo 结构体切片
	var noticeList []Notice
	err = json.Unmarshal(resultBytes, &noticeList)
	if err != nil {
		log.Printf("Error unmarshaling 'results' into []BusInfo: %v\n", err)
		return nil, fmt.Errorf("error unmarshaling 'results' into []BusInfo: %w", err)
	}

	// 返回解析后的数据
	return noticeList, nil
}

func MarkReaded(authToken, noticeID string) error {
	// 标记已读信息
	url := api.EBusHost + string(api.UserReaded)
	url = strings.Replace(string(url), "{id}", noticeID, 1)

	_, err := fetch.Client.R().
		SetHeader("Authorization", authToken).
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0").
		SetHeader("Accept", "*/*").
		SetHeader("Referer", "https://h5.pinbayun.com/").
		SetHeader("Origin", "https://h5.pinbayun.com").
		Get(url)

	if err != nil {
		log.Printf("Error sending request to %s: %v\n", url, err)
		return fmt.Errorf("failed to send request: %w", err)
	}
	return nil
}
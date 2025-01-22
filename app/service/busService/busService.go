package busService

import (
	"EBUSU/app/apiException"
	"EBUSU/app/fetch"
	"EBUSU/app/models"
	"EBUSU/config/api"
	"EBUSU/config/redis"
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

// GetBusInfo 获取公交信息
func FetchBusInfo(authToken string, page int, pageSize int) ([]models.BusInfo, error) {
	url := api.EBusHost + string(api.BusInfo)

	// 发起请求
	resp, err := fetch.Client.R().
		SetQueryParams(map[string]string{
			"page":      strconv.Itoa(page),
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
		return nil, apiException.RequestError
	}

	// 错误处理：HTTP 状态码检查
	if resp.StatusCode() == 400 {
		log.Printf("Received non-OK HTTP status from %s: %d\n", url, resp.StatusCode())
		return []models.BusInfo{}, apiException.AuthWrong
	} else if resp.StatusCode() == 500 {
		return []models.BusInfo{}, apiException.ServerError
	}

	// 响应内容，并提取出result字段
	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		log.Printf("Error unmarshaling JSON response: %v\n", err)
		return nil, apiException.ResponseError
	}

	// 提取 results 字段
	results, ok := result["results"].([]interface{})
	if !ok {
		log.Printf("Error: 'results' field not found or is not an array.\n")
		return nil, apiException.ResponseError
	}

	// 将 results 转换为 []models.BusInfo
	resultBytes, err := json.Marshal(results)
	if err != nil {
		log.Printf("Error marshaling 'results' to JSON: %v\n", err)
		return nil, apiException.ResponseError
	}

	// 解析为 models.BusInfo 结构体切片
	var busInfo []models.BusInfo
	err = json.Unmarshal(resultBytes, &busInfo)
	if err != nil {
		log.Printf("Error unmarshaling 'results' into []models.BusInfo: %v\n", err)
		return nil, apiException.ResponseError
	}

	// 返回解析后的数据
	return busInfo, nil
}

func FetchBusTime(authToken, busID string, shuttle_type int) ([]models.BusTime, error) {
	url := api.EBusHost + string(api.BusTime)
	url = strings.Replace(string(url), "{id}", busID, 1)

	resp, err := fetch.Client.R().
		SetQueryParams(map[string]string{
			"shuttle_type": "-10",
		}).
		SetHeader("Authorization", authToken).
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0").
		SetHeader("Accept", "*/*").
		SetHeader("Referer", "https://h5.pinbayun.com/").
		SetHeader("Origin", "https://h5.pinbayun.com").
		Get(url)

	if err != nil {
		log.Printf("Error sending request to %s: %v\n", url, err)
		return nil, apiException.RequestError
	}

	var busTimes []models.BusTime
	err = json.Unmarshal(resp.Body(), &busTimes)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
		return nil, apiException.ResponseError
	}

	return busTimes, nil
}

type responseStruct struct {
	Results []models.BusRecords `json:"results"`
}

func FetchBusRecords(authToken, page, pageSize, status string) ([]models.BusRecords, error) {
	// status参数为10或者20返回未开始预约信息，为30或者40返回已结束的成功预定信息，为0返回所有预定记录（包括已经取消）
	url := api.EBusHost + string(api.BusRecords)

	// 发起请求
	resp, err := fetch.Client.R().
		SetQueryParams(map[string]string{
			"page":      page,
			"page_size": pageSize,
			"status":    status,
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
		return nil, apiException.RequestError
	}

	// 错误处理：HTTP 状态码检查
	if resp.StatusCode() == 400 {
		log.Printf("Received non-OK HTTP status from %s: %d\n", url, resp.StatusCode())
		return nil, apiException.AuthWrong
	} else if resp.StatusCode() == 500 {
		return nil, apiException.ServerError
	}

	// fmt.Println(resp)
	// 响应内容，并提取出result字段
	var result responseStruct
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		log.Printf("Error unmarshaling JSON response: %v\n", err)
		return nil, apiException.ResponseError
	}

	// 提取 results 字段
	// results, ok := result["results"]
	// if !ok {
	// 	log.Printf("Error: 'results' field not found or is not an array.\n")
	// 	return nil, apiException.ResponseError
	// }

	// 将 results 转换为 []models.BusInfo
	// resultBytes, err := json.Marshal(results)
	// if err != nil {
	// 	log.Printf("Error marshaling 'results' to JSON: %v\n", err)
	// 	return nil, apiException.ResponseError
	// }

	// 解析为 models.BusInfo 结构体切片
	// var busInfo []BusRecords
	// err = json.Unmarshal(resultBytes, &busInfo)
	// if err != nil {
	// 	log.Printf("Error unmarshaling 'results' into []models.BusInfo: %v\n", err)
	// 	return nil, apiException.ResponseError
	// }

	// 返回解析后的数据
	return result.Results, nil
}

func GetBusInfo(search string) ([]models.BusInfo, error) {
	var filteredBusInfoList []models.BusInfo
	busInfoListData, err := redis.RedisClient.LRange("BusInfo", 0, -1).Result()

	if err != nil {
		log.Printf("Error: failed to get bus info list from Redis: %v", err)
		return nil, err
	}
	for _, businfo := range busInfoListData {
		if strings.Contains(businfo, search) {
			var tmp models.BusInfo
			err := json.Unmarshal([]byte(businfo), &tmp)
			if err != nil {
				log.Printf("Error: failed to unmarshal bus info: %v", err)
				continue
			}
			filteredBusInfoList = append(filteredBusInfoList, tmp)
		}
	}
	return filteredBusInfoList, nil
}

// func GetCallback() {

// 	// 设置请求头
// 	headers := map[string]string{
// 		"User-Agent": "Mozilla/5.0 (Linux; Android 14; 23013RK75C Build/UKQ1.230804.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/131.0.6778.200 Mobile Safari/537.36 ZJYXYwebviewbroswer ZJYXYAndroid tourCustomer/yunmaapp.NET/6.4.7/ym-e7ff7527536036bd01d60ba8babf01c6",
// 	}

// 	// 构建请求 URL 和参数
// 	baseURL := "https://open.xiaofubao.com/routeauth/auth/route/ua/authorize/getCodeV2"
// 	params := map[string]string{
// 		"ymAppId":     "2011112043190345310",
// 		"callbackUrl": "https://api.pinbayun.com/api/v1/zjgd_interface/?schoolCode=10337",
// 		"authType":    "2",
// 		"authAppid":   "10337",
// 		"unionid":     "2408157831570432101",
// 		"schoolCode":  "10337",
// 	}

// 	// 发起第一个请求
// 	resp, err := fetch.Client.R().
// 		SetHeaders(headers).
// 		SetQueryParams(params).
// 		Get(baseURL)

// 	if err != nil {
// 		log.Fatalf("Failed to send request: %v", err)
// 	}

// 	// 获取 Location 响应头
// 	redirectUrl := resp.Header().Get("Location")
// 	fmt.Println("Redirect URL:", redirectUrl)

// 	// 发起第二个请求，获取最终的重定向 URL
// 	resp2, err := fetch.Client.R().
// 		Get(redirectUrl)

// 	if err != nil {
// 		log.Fatalf("Failed to follow redirect: %v", err)
// 	}

// 	// 获取最终的重定向 URL
// 	finalRedirectUrl := resp2.Header().Get("Location")
// 	fmt.Println("Final Redirect URL:", finalRedirectUrl)
// }

package userService

import (
	"EBUSU/app/apiException"
	"EBUSU/app/fetch"
	"EBUSU/app/models"
	"EBUSU/config/api"
	"encoding/json"
	"log"
	"net/url"
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

	if resp.StatusCode() == 400 {
		log.Printf("Auth Error: %v\n", err)
		return 0, apiException.AuthWrong
	}

	if err != nil {
		log.Printf("Error sending request to %s: %v\n", url, err)
		return 0, apiException.RequestError
	}

	var unreadCount UnreadCount
	err = json.Unmarshal(resp.Body(), &unreadCount)
	if err != nil {
		log.Printf("Error unmarshaling JSON: %v", err)
		return 0, apiException.ResponseError
	}

	return unreadCount.Count, nil
}

func CheckTokenAlive(authToken string) error {
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
		return apiException.RequestError
	}

	if resp.StatusCode() != 200 {
		log.Printf("Received non-OK HTTP status from %s: %d\n", url, resp.StatusCode())
		return apiException.AuthWrong
	}

	return nil
}

type QrcodeString struct {
	QrcodeString string `json:"qrcode"`
}

func GetQrcode(authToken string) (string, error) {
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
		return "", apiException.RequestError
	}

	if resp.StatusCode() == 400 {
		log.Printf("Received non-OK HTTP status from %s: %d\n", url, resp.StatusCode())
		return "", apiException.AuthWrong
	} else if resp.StatusCode() == 500 {
		return "", apiException.ServerError
	}

	if len(resp.Body()) == 0 {
		return "", apiException.ResponseError
	}

	var qrcodeString QrcodeString
	if err = json.Unmarshal(resp.Body(), &qrcodeString); err != nil {
		log.Printf("Error unmarshaling JSON: %v\n", err)
		return "", apiException.ResponseError
	}

	if qrcodeString.QrcodeString == "" {
		return "", apiException.ResponseError
	}

	return qrcodeString.QrcodeString, nil
}

func GetNotice(authToken, page, pageSize, numPages string) ([]models.Notice, error) {
	// status参数为10或者20返回未开始预约信息，为30或者40返回已结束的成功预定信息，为0返回所有预定记录（包括已经取消）
	url := api.EBusHost + string(api.UserNotice)

	// 发起请求
	resp, err := fetch.Client.R().
		SetQueryParams(map[string]string{
			"page":      page,
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
		return nil, apiException.RequestError
	}

	// 错误处理：HTTP 状态码检查
	if resp.StatusCode() == 400 {
		log.Printf("Received non-OK HTTP status from %s: %d\n", url, resp.StatusCode())
		return []models.Notice{}, apiException.AuthWrong
	} else if resp.StatusCode() == 500 {
		return []models.Notice{}, apiException.ServerError
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

	resultBytes, err := json.Marshal(results)
	if err != nil {
		log.Printf("Error marshaling 'results' to JSON: %v\n", err)
		return nil, apiException.ResponseError
	}

	// 解析为 BusInfo 结构体切片
	var noticeList []models.Notice
	err = json.Unmarshal(resultBytes, &noticeList)
	if err != nil {
		log.Printf("Error unmarshaling 'results' into []BusInfo: %v\n", err)
		return nil, apiException.ResponseError
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
		return apiException.RequestError
	}
	return nil
}

func LoginByPhone(phone, password string) (string, error) {
	// 登录接口
	url := api.EBusHost + string(api.UserLoginByPhone)

	// 发起请求
	resp, err := fetch.Client.R().
		SetBody(map[string]string{
			"phone":    phone,
			"password": password,
		}).
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0").
		SetHeader("Accept", "*/*").
		SetHeader("Referer", "https://h5.pinbayun.com/").
		SetHeader("Origin", "https://h5.pinbayun.com").
		Post(url)

	if err != nil {
		return "", apiException.RequestError
	}

	var result map[string]interface{}
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return "", apiException.ResponseError
	}

	token, ok := result["token"].(string)
	if !ok {
		return "", apiException.NoThatPasswordOrWrong
	}

	return token, nil
}

func LoginByYxy(unionid string) (string, error) {
	// Step 1: Get authentication redirect
	token, err := getAuthRedirect(unionid)
	if err != nil {
		return "", err
	}

	// Step 2: Perform YXY login
	return performYxyLogin(token.openid, token.corpcode)
}

type authToken struct {
	openid   string
	corpcode string
}

func getAuthRedirect(unionid string) (*authToken, error) {
	authurl := string(api.UserOauthLogin)
	resp, err := fetch.Client.R().
		SetQueryParams(map[string]string{
			"ymAppId":     "2011112043190345310",
			"callbackUrl": "https://api.pinbayun.com/api/v1/zjgd_interface/?schoolCode=10337",
			"authType":    "2",
			"authAppid":   "10337",
			"unionid":     unionid,
			"schoolCode":  "10337",
		}).
		SetHeader("User-Agent", "Mozilla/5.0 (Linux; Android 14; 23013RK75C Build/UKQ1.230804.001; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/131.0.6778.200 Mobile Safari/537.36 ZJYXYwebviewbroswer ZJYXYAndroid tourCustomer/yunmaapp.NET/6.4.7/ym-e7ff7527536036bd01d60ba8babf01c6").
		Get(authurl)

	if err != nil {
		return nil, apiException.RequestError
	}

	if resp == nil || resp.RawResponse == nil || resp.RawResponse.Request == nil || resp.RawResponse.Request.Response == nil {
		return nil, apiException.ResponseError
	}

	redirectURL := resp.RawResponse.Request.Response.Header.Get("Location")
	if redirectURL == "" {
		return nil, apiException.ResponseError
	}

	parsedURL, err := url.Parse(redirectURL)
	if err != nil {
		return nil, apiException.ResponseError
	}

	queryParams := parsedURL.Query()
	openid := queryParams.Get("openid")
	corpcode := queryParams.Get("corpcode")

	if openid == "" || corpcode == "" {
		return nil, apiException.ResponseError
	}

	return &authToken{openid: openid, corpcode: corpcode}, nil
}

func performYxyLogin(openid, corpcode string) (string, error) {
	loginURL := api.EBusHost + string(api.UserLoginByYxy)
	resp, err := fetch.Client.R().
		SetBody(map[string]string{
			"openid":   openid,
			"corpcode": corpcode,
		}).
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36 Edg/131.0.0.0").
		SetHeader("Accept", "*/*").
		SetHeader("Referer", "https://h5.pinbayun.com/").
		SetHeader("Origin", "https://h5.pinbayun.com").
		Post(loginURL)

	if err != nil {
		return "", apiException.RequestError
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return "", apiException.ResponseError
	}

	token, ok := result["token"].(string)
	if !ok || token == "" {
		return "", apiException.AuthWrong
	}

	return token, nil
}

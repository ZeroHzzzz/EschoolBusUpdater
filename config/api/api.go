package api

import "EBUSU/config/config"

var EBusHost = config.Config.GetString("eBus.host")
var SystemToken = config.Config.GetString("eBus.token")

type EBusApi string
const (
	LoginByPhone EBusApi = "/api/v4/staff/auths/login/"
	LoginByWX EBusApi = "/api/v1/staff/auths/wx_auth/"
	BusInfo EBusApi = "/api/v2/staff/shuttlebus/"
	BusTime EBusApi = "/api/v2/staff/shuttlebus/{id}/bustimes/"
	BusDate EBusApi = "/api/v2/staff/shuttlebus/{id}/dates/"
	BusBook EBusApi = "/api/v3/staff/busorders/"
	BusRecords EBusApi = "/api/v1/staff/busorders/"
	BusCancel EBusApi = "/api/v3/staff/busorders/{id}/cancel/"
	BusBulktask EBusApi = "/api/v4/staff/busorderbulktask/"
	
	UserQrcode EBusApi = "/api/v3/pos/staff_qrcode/"
	UserCompany EBusApi = "/api/v4/staff/auths/get_company/"
	UserUnreadCount EBusApi = "/api/v1/staff/messages/unread_count/"
	UserNotice EBusApi = "/api/v1/staff/messages/"
	UserReaded EBusApi = "/api/v1/staff/messages/{id}/read/" // 标记已读
)

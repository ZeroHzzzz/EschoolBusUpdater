package models

// 整体逻辑： 线路 -> 班次, 然后再通过这两个id去换订购的id

type Station struct {
	ID          string  `json:"id"`
	StationName string  `json:"station_name"`
	StationSeq  int     `json:"station_seq"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
}

type OtherPrice struct {
	Teacher       int `json:"teacher"`
	Student       int `json:"student"`
	StudentTicket int `json:"student_ticket"`
}

type BusInfo struct {
	// 这个是线路信息
	ID          string `json:"id"`
	ShuttleName string `json:"shuttle_name"`
	Seats       int    `json:"seats"`
	// ApplyExpiredMinutes int    `json:"apply_expired_minutes"`
	Price int `json:"price"`
	// OrderMode            int       `json:"order_mode"`
	// CheckMode            int       `json:"check_mode"`
	// OrderModeName        string    `json:"order_mode_name"`
	// CheckModeName        string    `json:"check_mode_name"`
	GoStationsJson []Station `json:"go_stations_json"`
	// ReturnStationsJson []Station `json:"return_stations_json"`
	// AutoDispatch bool `json:"auto_dispatch"`
	// PeopleVehicle        bool      `json:"people_vehicle"`
	// ApplyDispatchMinutes int `json:"apply_dispatch_minutes"`
	// LongDistance         bool `json:"long_distance"`
	// Overtime      bool `json:"overtime"`
	TicketOfBuses bool `json:"ticket_of_buses"`
	ReservedSeat  int  `json:"reserved_seat"`
	// OtherPrice                 OtherPrice `json:"other_price"`
	// BusReminder      int  `json:"bus_reminder"`
	SortNumber       int  `json:"sort_number"`
	OrderLimitOn     bool `json:"order_limit_on"`
	OrderLimitMinute int  `json:"order_limit_minute"`
	// NetMode            int  `json:"net_mode"`
	TicketCheckMinutes int `json:"ticket_check_minutes"`
	// BlukOrder          bool `json:"bluk_order"`
	// IsBlukOrderStaffLevel bool `json:"is_bluk_order_staff_level"`
	// BlukOrderStaffLevels  string `json:"bluk_order_staff_levels"`
	// RemainderRemind int `json:"remainder_remind"`
	// RemindAdmin                string    `json:"remind_admin"`
	// RefundAtDispath            bool `json:"refund_at_dispath"`
	// RefundAtDispathLimitMinute int  `json:"refund_at_dispath_limit_minute"`
	IsActive bool `json:"is_active"`
	// InstanceDays int  `json:"instance_days"`
	// DriverDispatch             bool      `json:"driver_dispatch"`
	Busfavourite bool      `json:"busfavourite"`
	BusTimeList  []BusTime `json:"bus_times"`
}

type BusTime struct {
	// 这个东西只有id是有用的，bustime接口返回的是一个列表，然后每一项是一个班次。
	DepartureTime   string `json:"departure_time"`
	ShuttleType     int    `json:"shuttle_type"`
	ShuttleTypeName string `json:"shuttle_type_name"`
	ID              string `json:"id"`
}

type BusDate struct {
	BusInfoData BusInfo `json:"shuttle_bus_vo"`
	BusTimeData BusTime `json:"shuttle_bus_time_vo"`
}
type BusRecords struct {
	BusDateData   BusDate `json:"shuttle_bus_date_vo"`
	PayTime       string  `json:"pay_time"`
	DepartureTime string  `json:"departure_datetime"`
}

type BusStatus struct {
	Status       string `json:"status"`
	LastUpdateAt string `json:"lastUpdateAt"`
}

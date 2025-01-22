package models

type ExtraParam struct {
	ActionType int         `json:"action_type"`
	To         interface{} `json:"to"`
}

// Notice 结构体
type Notice struct {
	ID         string     `json:"id"`
	MsgType    string     `json:"msg_type"`
	MsgID      string     `json:"msg_id"`
	Title      string     `json:"title"`
	ExtraParam ExtraParam `json:"extra_param"`
	IsRead     int        `json:"is_read"`
	Content    string     `json:"content"`
	HTML       string     `json:"html"`
	Img        string     `json:"img"`
	Author     string     `json:"author"`
}

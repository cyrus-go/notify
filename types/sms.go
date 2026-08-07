package types

// YPSendSmsResp 云片发送短信返回
type YPSendSmsResp struct {
	Code   int     `json:"code"`
	Msg    string  `json:"msg"`
	Count  int     `json:"count"`
	Fee    float64 `json:"fee"`
	Unit   string  `json:"unit"`
	Mobile string  `json:"mobile"`
	SID    int64   `json:"SID"`
}

// YmrtSendSmsResp 亿美软通发送短信返回
type YmrtSendSmsResp struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

// BeanliSendSmsResp 滨礼网络发送短信返回。
type BeanliSendSmsResp struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Data    BeanliSendSmsData `json:"data"`
}

// BeanliSendSmsData 滨礼网络短信任务信息。
type BeanliSendSmsData struct {
	TaskID     string `json:"taskId"`
	CreateDate string `json:"createDate"`
}

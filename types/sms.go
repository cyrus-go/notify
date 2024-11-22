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

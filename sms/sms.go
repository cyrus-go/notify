package sms

import (
	"github.com/cyrus-go/notify/pkg/tools"
	"log"
	"net/url"
)

// todo: 临时存放，后续放到配置文件中
const apiKey = "14b1bbedae4ba4403778f551948e8026"

func getSmsInfo() (err error) {
	// 获取user信息url
	urlGetUser := "https://sms.yunpian.com/v2/user/get.json"
	dataGetUser := url.Values{"apikey": {apiKey}}
	err = tools.PostForm(urlGetUser, dataGetUser)
	if err != nil {
		log.Printf("get user info failed, err: %v", err)
		return
	}
	return
}

// 发送短信（可群发，群发时手机号用逗号隔开）
func sendSms(mobile, tplId string, tplContent url.Values) (err error) {
	// 发送模板内容编译
	tplValue := tplContent.Encode()
	// 群发模板 url
	urlTplSms := "https://sms.yunpian.com/v2/sms/tpl_batch_send.json"

	dataTplSms := url.Values{"apikey": {apiKey}, "mobile": {mobile}, "tpl_id": {tplId}, "tpl_value": {tplValue}}

	err = tools.PostForm(urlTplSms, dataTplSms)
	if err != nil {
		log.Printf("send tpl sms failed, err: %v", err)
		return
	}
	return
}

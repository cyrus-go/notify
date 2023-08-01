package sms

import (
	"github.com/cyrus-go/notify/pkg/tools"
	"log"
	"net/url"
)

func getSmsInfo(apikey string) (err error) {
	// 获取user信息url
	urlGetUser := "https://sms.yunpian.com/v2/user/get.json"
	dataGetUser := url.Values{"apikey": {apikey}}
	err = tools.PostForm(urlGetUser, dataGetUser)
	if err != nil {
		log.Printf("get user info failed, err: %v", err)
		return
	}
	return
}

// 发送短信（可群发，群发时手机号用逗号隔开）
func sendSms(apikey, mobile, tplId string, tplContent url.Values) (err error) {
	// 发送模板内容编译
	tplValue := tplContent.Encode()
	// 群发模板 url
	urlTplSms := "https://sms.yunpian.com/v2/sms/tpl_batch_send.json"

	dataTplSms := url.Values{"apikey": {apikey}, "mobile": {mobile}, "tpl_id": {tplId}, "tpl_value": {tplValue}}

	err = tools.PostForm(urlTplSms, dataTplSms)
	if err != nil {
		log.Printf("send tpl sms failed, err: %v", err)
		return
	}
	return
}

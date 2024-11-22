package sms

import (
	"encoding/json"
	"fmt"
	"github.com/cyrus-go/notify/pkg/tools"
	"github.com/cyrus-go/notify/types"
	"github.com/pkg/errors"
	"log"
	"net/url"
)

func getSmsInfo(apikey string) error {
	// 获取user信息url
	urlGetUser := "https://sms.yunpian.com/v2/user/get.json"
	dataGetUser := url.Values{"apikey": {apikey}}
	bytes, err := tools.PostForm(urlGetUser, dataGetUser)
	if err != nil {
		log.Printf("get user info failed, err: %v", err)
		return errors.Errorf("get user info failed, err: %v", err)
	}
	fmt.Printf("get user info success, resp: %s\n", string(bytes))
	return nil
}

// 发送短信（可群发，群发时手机号用逗号隔开）
func sendSms(apikey, mobile, tplId string, tplContent url.Values) error {
	// 发送模板内容编译
	tplValue := tplContent.Encode()
	// 群发模板 url
	urlTplSms := "https://sms.yunpian.com/v2/sms/tpl_single_send.json"

	dataTplSms := url.Values{"apikey": {apikey}, "mobile": {mobile}, "tpl_id": {tplId}, "tpl_value": {tplValue}}

	body, err := tools.PostForm(urlTplSms, dataTplSms)
	if err != nil {
		log.Printf("send tpl sms failed, err: %v", err)
		return errors.Errorf("send tpl sms failed, err: %v", err)
	}

	// 解析返回的json数据
	var smsResp types.YPSendSmsResp
	if err = json.Unmarshal(body, &smsResp); err != nil {
		return errors.Errorf("unmarshal send sms response failed, err: %v", err)
	}

	if smsResp.Code != 0 {
		return errors.Errorf("send sms failed, err: %v", smsResp)
	}

	return nil
}

// 亿美软通短信
package sms

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/cyrus-go/notify/pkg/tools"
	"github.com/cyrus-go/notify/types"
	"github.com/pkg/errors"
	"net/url"
	"time"
)

func SendSmsYmrt(appID, secretKey, phone, content string) error {
	timestamp := time.Now().Format("20060102150405")
	// 创建 MD5 哈希
	hash := md5.New()
	hash.Write([]byte(appID + secretKey + timestamp))

	// 获取 MD5 结果并转换为 32 位字符串
	sign := hex.EncodeToString(hash.Sum(nil))

	urlSendSms := fmt.Sprintf("http://www.btom.cn:8080/simpleinter/sendSMS?appId=%s&timestamp=%s&sign=%s&mobiles=%s&content=%s", appID, timestamp, sign, phone, url.QueryEscape(content))
	bytes, err := tools.GetUrl(urlSendSms)
	if err != nil {
		return errors.Errorf("send sms failed, err: %v", err)
	}

	fmt.Println("send sms resp is: ", string(bytes))
	// 解析返回的 json 数据
	var smsResp types.YmrtSendSmsResp
	if err = json.Unmarshal(bytes, &smsResp); err != nil {
		return errors.Errorf("unmarshal send sms response failed, err: %v", err)
	}

	if smsResp.Code != "SUCCESS" {
		return errors.Errorf("send sms failed, err code: %v", smsResp.Code)
	}

	return nil
}

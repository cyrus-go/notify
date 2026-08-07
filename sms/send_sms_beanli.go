// 滨礼网络短信
package sms

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/cyrus-go/notify/types"
	"github.com/pkg/errors"
)

const beanliSmsEndpoint = "http://sms.beanli.com:7891/api/v1/send"

var beanliHTTPClient = &http.Client{Timeout: 10 * time.Second}

type beanliSendSmsRequest struct {
	Account string `json:"account"`
	Pwd     string `json:"pwd"`
	Mobiles string `json:"mobiles"`
	Content string `json:"content"`
}

// SendSmsBeanli 通过滨礼网络发送文字短信。
// password 为管理员提供的接口密码，发送前会按文档要求转换为 32 位小写 MD5。
func SendSmsBeanli(account, password, mobile, content string) error {
	return sendSmsBeanli(beanliHTTPClient, beanliSmsEndpoint, account, password, mobile, content)
}

func sendSmsBeanli(client *http.Client, endpoint, account, password, mobile, content string) error {
	// 滨礼网络要求 pwd 使用接口密码的 32 位 MD5 值。
	passwordMD5 := md5.Sum([]byte(password))
	requestBody, err := json.Marshal(beanliSendSmsRequest{
		Account: account,
		Pwd:     hex.EncodeToString(passwordMD5[:]),
		Mobiles: mobile,
		Content: content,
	})
	if err != nil {
		return errors.Wrap(err, "marshal beanli send sms request failed")
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(requestBody))
	if err != nil {
		return errors.Wrap(err, "create beanli send sms request failed")
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "send beanli sms request failed")
	}
	defer resp.Body.Close()

	// 限制响应体大小，避免异常响应占用过多内存。
	responseBody, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return errors.Wrap(err, "read beanli send sms response failed")
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return errors.Errorf("send beanli sms failed, http status: %d", resp.StatusCode)
	}

	var smsResp types.BeanliSendSmsResp
	if err = json.Unmarshal(responseBody, &smsResp); err != nil {
		return errors.Wrap(err, "unmarshal beanli send sms response failed")
	}
	if smsResp.Code != "00" {
		return errors.Errorf("send beanli sms failed, code: %s, message: %s", smsResp.Code, smsResp.Message)
	}

	return nil
}

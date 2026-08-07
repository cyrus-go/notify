package test

import (
	"testing"

	"github.com/cyrus-go/notify/sms"
)

// TestSendSmsBeanliReal 通过滨礼网络真实发送一条 Cephalon Cloud 登录验证码短信。
// 执行前需配置 BEANLI_SMS_ACCOUNT、BEANLI_SMS_PASSWORD 和 BEANLI_SMS_MOBILE。
func TestSendSmsBeanliReal(t *testing.T) {
	account := "922847"
	password := "EgXyEu3g"
	mobile := "13187098660"
	code := "123456"

	smsClient, err := sms.NewVerifySmsBeanli(account, password, mobile, code)
	if err != nil {
		t.Fatalf("创建滨礼网络短信客户端失败: %v", err)
	}
	if err = smsClient.SendSmsLoginCephalon(); err != nil {
		t.Fatalf("滨礼网络真实短信发送失败: %v", err)
	}
}

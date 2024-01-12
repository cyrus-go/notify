package test

import (
	"fmt"
	"github.com/cyrus-go/notify/sms"
	"testing"
)

func TestSmsTencentCephalonCoreRegister(t *testing.T) {
	phones := []string{"13888888888"}
	code := "123456"
	sdkAppId := "your_sdk_app_id"
	secretId := "your_secret_id"
	secretKey := "your_secret_key"
	smsCli, err := sms.NewVerifySmsTencent(sdkAppId, secretId, secretKey, code, phones)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	respBytes, err := smsCli.SendTencentSmsCephalonCoreRegister()
	if err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
	fmt.Printf("%s\n", respBytes)
}

func TestSmsTencentCephalonCoreLogin(t *testing.T) {
	phones := []string{"+8613888888888"}
	code := "234567"
	sdkAppId := "your_sdk_app_id"
	secretId := "your_secret_id"
	secretKey := "your_secret_key"
	smsCli, err := sms.NewVerifySmsTencent(sdkAppId, secretId, secretKey, code, phones)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	respBytes, err := smsCli.SendTencentSmsCephalonCoreLogin()
	if err != nil {
		t.Fatalf("send login cephalon sms failed, err: %v", err)
	}
	fmt.Printf("%s\n", respBytes)
}

func TestSmsTencentCephalonCoreModifyPwd(t *testing.T) {
	phones := []string{"+8613888888888"}
	code := "666666"
	sdkAppId := "your_sdk_app_id"
	secretId := "your_secret_id"
	secretKey := "your_secret_key"
	smsCli, err := sms.NewVerifySmsTencent(sdkAppId, secretId, secretKey, code, phones)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	respBytes, err := smsCli.SendTencentSmsCephalonCoreModifyPwd()
	if err != nil {
		t.Fatalf("send modify pwd cephalon sms failed, err: %v", err)
	}
	fmt.Printf("%s\n", respBytes)
}

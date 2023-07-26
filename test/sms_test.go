package test

import (
	"github.com/cyrus-go/notify/sms"
	"testing"
)

func TestSmsRegisterCephalon(t *testing.T) {
	mobile := "13187098660"
	code := "666666"
	smsCli := sms.NewVerifySms(mobile, code)
	// 发送模板内容
	err := smsCli.SendSmsRegisterCephalon()
	if err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

func TestSmsLoginCephalon(t *testing.T) {
	mobile := "13187098660"
	code := "777777"
	smsCli := sms.NewVerifySms(mobile, code)
	// 发送模板内容
	err := smsCli.SendSmsLoginCephalon()
	if err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

func TestSmsModifyPwdCephalon(t *testing.T) {
	mobile := "13187098660"
	code := "888888"
	smsCli := sms.NewVerifySms(mobile, code)
	// 发送模板内容
	err := smsCli.SendSmsModifyPwdCephalon()
	if err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

func TestSmsRegisterYuanHui(t *testing.T) {
	mobile := "13187098660"
	code := "999999"
	smsCli := sms.NewVerifySms(mobile, code)
	// 发送模板内容
	err := smsCli.SendSmsRegisterYuanHui()
	if err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

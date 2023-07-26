package test

import (
	"github.com/cyrus-go/notify/sms"
	"testing"
)

func TestSmsRegisterCephalon(t *testing.T) {
	mobile := "13187098660"
	code := "666666"
	smsCli, err := sms.NewVerifySms(mobile, code)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsRegisterCephalon(); err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

func TestSmsLoginCephalon(t *testing.T) {
	mobile := "13187098660"
	code := "777777"
	smsCli, err := sms.NewVerifySms(mobile, code)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsLoginCephalon(); err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

func TestSmsModifyPwdCephalon(t *testing.T) {
	mobile := "13187098660"
	code := "888888"
	smsCli, err := sms.NewVerifySms(mobile, code)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsModifyPwdCephalon(); err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

func TestSmsRegisterYuanHui(t *testing.T) {
	mobile := "13187098660"
	code := "999999"
	smsCli, err := sms.NewVerifySms(mobile, code)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsRegisterYuanHui(); err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

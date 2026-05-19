package test

import (
	"github.com/cyrus-go/notify/sms"
	"testing"
)

func TestSmsRegisterCephalonCore(t *testing.T) {
	mobile := "13187098660"
	code := "666666"
	smsCli, err := sms.NewVerifySms("your_apikey", mobile, code)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsRegisterCephalonCore(); err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

func TestSmsLoginCephalonCore(t *testing.T) {
	mobile := "13187098660"
	code := "777777"
	smsCli, err := sms.NewVerifySms("your_apikey", mobile, code)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsLoginCephalonCore(); err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

func TestSmsModifyPwdCephalonCore(t *testing.T) {
	mobile := "13187098660"
	code := "888888"
	smsCli, err := sms.NewVerifySms("your_apikey", mobile, code)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsModifyPwdCephalonCore(); err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

func TestSmsRegisterYuanHui(t *testing.T) {
	mobile := "13187098660"
	code := "999999"
	smsCli, err := sms.NewVerifySms("your_apikey", mobile, code)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsRegisterYuanHui(); err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

func TestSmsRegisterCephalon(t *testing.T) {
	mobile := "13187098660"
	code := "111111"
	smsCli, err := sms.NewVerifySms("your_apikey", mobile, code)
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
	code := "222222"
	smsCli, err := sms.NewVerifySms("your_apikey", mobile, code)
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
	code := "333333"
	smsCli, err := sms.NewVerifySms("your_apikey", mobile, code)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsModifyPwdCephalon(); err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

func TestInterfaceSmsLoginCephalon(t *testing.T) {
	mobile := "13187098660"
	code := "222222"

	var (
		smsCli sms.Verify
		err    error
	)

	smsCli, err = sms.NewVerifySms("your_apikey", mobile, code)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsLoginCephalon(); err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}

	smsCli, err = sms.NewVerifySmsYmrt("your_app_id", "your_secret_key", mobile, code)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsLoginCephalon(); err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

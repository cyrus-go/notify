package test

import (
	"github.com/cyrus-go/notify/sms"
	"testing"
)

const (
	appID     = "your_appid"
	secretKey = "your_secret_key"
)

func TestSmsYmrtRegisterCephalon(t *testing.T) {
	mobile := "13187098660"
	code := "register"
	smsCli, err := sms.NewVerifySmsYmrt(appID, secretKey, mobile, code)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsRegisterCephalon(); err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

func TestSmsYmrtLoginCephalon(t *testing.T) {
	mobile := "13187098660"
	code := "login"
	smsCli, err := sms.NewVerifySmsYmrt(appID, secretKey, mobile, code)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsLoginCephalon(); err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

func TestSmsYmrtModifyPwdCephalon(t *testing.T) {
	mobile := "13187098660"
	code := "pwd"
	smsCli, err := sms.NewVerifySmsYmrt(appID, secretKey, mobile, code)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsModifyPwdCephalon(); err != nil {
		t.Fatalf("send register cephalon sms failed, err: %v", err)
	}
}

func TestSmsYmrtNotify(t *testing.T) {
	mobile := "13187098660"
	content := "【端脑科技】温馨提示：您正在使用的 AI 应用即将到期，系统显示您的余额不足以进行自动续费，如需继续使用请及时充值噢！到期后应用会自动停用并清除现有设置，为了方便您的使用，请确保在服务到期前进行充值或备份数据。\n    端脑云感谢您的使用，祝您生活愉快！"
	smsCli, err := sms.NewNotifySmsYmrt(appID, secretKey, mobile)
	if err != nil {
		t.Logf("init verify sms failed, err: %v", err)
		return
	}
	// 发送模板内容
	if err = smsCli.SendSmsCustom(content); err != nil {
		t.Fatalf("send notify cephalon sms failed, err: %v", err)
	}
}

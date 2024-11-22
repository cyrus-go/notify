package sms

import (
	"fmt"
)

const (
	contentCephalonCloudRegister   = "【端脑科技】欢迎注册 Cephalon Cloud，本次注册的验证码为 %s，5分钟内有效"
	contentCephalonCloudLogin      = "【端脑科技】欢迎登陆 Cephalon Cloud，本次登陆的验证码为 %s，5分钟内有效"
	contentCephalonCloudModifyPwd  = "【端脑科技】您正在修改 Cephalon Cloud 密码，本次验证码为 %s，5分钟内有效"
	contentCephalonCloudBindWechat = "【端脑科技】您正在使用微信绑定 Cephalon Cloud，本次验证码为 %s，5分钟内有效"

	contentCephalonCoreRegister  = "【端脑科技】欢迎注册 Cephalon Core，本次注册的验证码为 %s，5分钟内有效"
	contentCephalonCoreLogin     = "【端脑科技】欢迎登陆 Cephalon Core，本次登陆的验证码为 %s，5分钟内有效"
	contentCephalonCoreModifyPwd = "【端脑科技】您正在修改 Cephalon Core 平台密码，本次验证码为 %s，5分钟内有效"
)

// SendSmsRegisterCephalon Cephalon 注册短信验证
func (v *VerifySmsYmrt) SendSmsRegisterCephalon() (err error) {
	return SendSmsYmrt(v.AppID, v.SecretKey, v.Mobile, fmt.Sprintf(contentCephalonCloudRegister, v.Code))
}

// SendSmsLoginCephalon Cephalon 登录短信验证
func (v *VerifySmsYmrt) SendSmsLoginCephalon() (err error) {
	return SendSmsYmrt(v.AppID, v.SecretKey, v.Mobile, fmt.Sprintf(contentCephalonCloudLogin, v.Code))
}

// SendSmsModifyPwdCephalon Cephalon 修改密码短信验证
func (v *VerifySmsYmrt) SendSmsModifyPwdCephalon() (err error) {
	return SendSmsYmrt(v.AppID, v.SecretKey, v.Mobile, fmt.Sprintf(contentCephalonCloudModifyPwd, v.Code))
}

// SendSmsBindWechatCephalon Cephalon 绑定微信短信验证
func (v *VerifySmsYmrt) SendSmsBindWechatCephalon() (err error) {
	return SendSmsYmrt(v.AppID, v.SecretKey, v.Mobile, fmt.Sprintf(contentCephalonCloudBindWechat, v.Code))
}

// SendSmsRegisterCephalonCore Cephalon core 注册短信验证
func (v *VerifySmsYmrt) SendSmsRegisterCephalonCore() (err error) {
	return SendSmsYmrt(v.AppID, v.SecretKey, v.Mobile, fmt.Sprintf(contentCephalonCoreRegister, v.Code))
}

// SendSmsLoginCephalonCore Cephalon core 登录短信验证
func (v *VerifySmsYmrt) SendSmsLoginCephalonCore() (err error) {
	return SendSmsYmrt(v.AppID, v.SecretKey, v.Mobile, fmt.Sprintf(contentCephalonCoreLogin, v.Code))
}

// SendSmsModifyPwdCephalonCore Cephalon core 修改密码短信验证
func (v *VerifySmsYmrt) SendSmsModifyPwdCephalonCore() (err error) {
	return SendSmsYmrt(v.AppID, v.SecretKey, v.Mobile, fmt.Sprintf(contentCephalonCoreModifyPwd, v.Code))
}

func (v *NotifySmsYmrt) SendSmsCustom(content string) (err error) {
	return SendSmsYmrt(v.AppID, v.SecretKey, v.Mobile, content)
}

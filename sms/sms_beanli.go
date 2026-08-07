package sms

import "fmt"

// 滨礼网络使用自定义文字短信，内容与亿美软通保持一致。
const (
	contentBeanliCephalonCloudRegister   = "【端脑科技】欢迎注册 Cephalon Cloud，本次注册的验证码为 %s，5分钟内有效"
	contentBeanliCephalonCloudLogin      = "【端脑科技】欢迎登陆 Cephalon Cloud，本次登陆的验证码为 %s，5分钟内有效"
	contentBeanliCephalonCloudModifyPwd  = "【端脑科技】您正在修改 Cephalon Cloud 密码，本次验证码为 %s，5分钟内有效"
	contentBeanliCephalonCloudBindWechat = "【端脑科技】您正在使用微信绑定 Cephalon Cloud，本次验证码为 %s，5分钟内有效"

	contentBeanliCephalonCoreRegister  = "【端脑科技】欢迎注册 Cephalon Core，本次注册的验证码为 %s，5分钟内有效"
	contentBeanliCephalonCoreLogin     = "【端脑科技】欢迎登陆 Cephalon Core，本次登陆的验证码为 %s，5分钟内有效"
	contentBeanliCephalonCoreModifyPwd = "【端脑科技】您正在修改 Cephalon Core 平台密码，本次验证码为 %s，5分钟内有效"
)

// SendSmsRegisterCephalon 发送 Cephalon Cloud 注册验证码短信。
func (v *VerifySmsBeanli) SendSmsRegisterCephalon() error {
	return SendSmsBeanli(v.Account, v.Password, v.Mobile, fmt.Sprintf(contentBeanliCephalonCloudRegister, v.Code))
}

// SendSmsLoginCephalon 发送 Cephalon Cloud 登录验证码短信。
func (v *VerifySmsBeanli) SendSmsLoginCephalon() error {
	return SendSmsBeanli(v.Account, v.Password, v.Mobile, fmt.Sprintf(contentBeanliCephalonCloudLogin, v.Code))
}

// SendSmsModifyPwdCephalon 发送 Cephalon Cloud 修改密码验证码短信。
func (v *VerifySmsBeanli) SendSmsModifyPwdCephalon() error {
	return SendSmsBeanli(v.Account, v.Password, v.Mobile, fmt.Sprintf(contentBeanliCephalonCloudModifyPwd, v.Code))
}

// SendSmsBindWechatCephalon 发送 Cephalon Cloud 绑定微信验证码短信。
func (v *VerifySmsBeanli) SendSmsBindWechatCephalon() error {
	return SendSmsBeanli(v.Account, v.Password, v.Mobile, fmt.Sprintf(contentBeanliCephalonCloudBindWechat, v.Code))
}

// SendSmsRegisterCephalonCore 发送 Cephalon Core 注册验证码短信。
func (v *VerifySmsBeanli) SendSmsRegisterCephalonCore() error {
	return SendSmsBeanli(v.Account, v.Password, v.Mobile, fmt.Sprintf(contentBeanliCephalonCoreRegister, v.Code))
}

// SendSmsLoginCephalonCore 发送 Cephalon Core 登录验证码短信。
func (v *VerifySmsBeanli) SendSmsLoginCephalonCore() error {
	return SendSmsBeanli(v.Account, v.Password, v.Mobile, fmt.Sprintf(contentBeanliCephalonCoreLogin, v.Code))
}

// SendSmsModifyPwdCephalonCore 发送 Cephalon Core 修改密码验证码短信。
func (v *VerifySmsBeanli) SendSmsModifyPwdCephalonCore() error {
	return SendSmsBeanli(v.Account, v.Password, v.Mobile, fmt.Sprintf(contentBeanliCephalonCoreModifyPwd, v.Code))
}

// SendSmsCustom 通过滨礼网络发送自定义短信内容。
func (v *NotifySmsBeanli) SendSmsCustom(content string) error {
	return SendSmsBeanli(v.Account, v.Password, v.Mobile, content)
}

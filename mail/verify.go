package mail

import (
	"fmt"
	"github.com/cyrus-go/notify/types"
)

const (
	titleCephalonCloudRegister   = "端脑云注册验证"
	titleCephalonCloudLogin      = "端脑云登录验证"
	titleCephalonCloudModifyPwd  = "端脑云修改密码验证"
	titleCephalonCloudBindWechat = "端脑云绑定微信验证"

	titleCephalonCoreRegister  = "端脑小程序注册验证"
	titleCephalonCoreLogin     = "端脑小程序登录验证"
	titleCephalonCoreModifyPwd = "端脑小程序修改密码验证"

	contextCephalonCloudRegister   = "【端脑科技】欢迎注册 Cephalon Cloud，本次注册的验证码为 %s，5分钟内有效"
	contextCephalonCloudLogin      = "【端脑科技】欢迎登陆 Cephalon Cloud，本次登陆的验证码为 %s，5分钟内有效"
	contextCephalonCloudModifyPwd  = "【端脑科技】您正在修改 Cephalon Cloud 密码，本次验证码为 %s，5分钟内有效"
	contextCephalonCloudBindWechat = "【端脑科技】您正在使用微信绑定 Cephalon Cloud，本次验证码为 %s，5分钟内有效"

	contextCephalonCoreRegister  = "【端脑科技】欢迎注册 Cephalon Core，本次注册的验证码为 %s，5分钟内有效"
	contextCephalonCoreLogin     = "【端脑科技】欢迎登陆 Cephalon Core，本次登陆的验证码为 %s，5分钟内有效"
	contextCephalonCoreModifyPwd = "【端脑科技】您正在修改 Cephalon Core 平台密码，本次验证码为 %s，5分钟内有效"
)

type VerifyMail struct {
	Username    string
	Password    string
	SendMail    string
	ReceiveMail string
	Code        string
}

func NewVerifyMail(username, password, sendMail, receiveMail, code string) *VerifyMail {
	return &VerifyMail{
		Username:    username,
		Password:    password,
		SendMail:    sendMail,
		ReceiveMail: receiveMail,
		Code:        code,
	}
}

// SendMailCephalonCloudRegister Cephalon Cloud 注册
func (v *VerifyMail) SendMailCephalonCloudRegister() error {
	return SendMailAli(v.Username, v.Password, v.SendMail, v.ReceiveMail, types.SendNameCephalon, titleCephalonCloudRegister, fmt.Sprintf(contextCephalonCloudRegister, v.Code))
}

// SendMailCephalonCloudLogin Cephalon Cloud 登录
func (v *VerifyMail) SendMailCephalonCloudLogin() error {
	return SendMailAli(v.Username, v.Password, v.SendMail, v.ReceiveMail, types.SendNameCephalon, titleCephalonCloudLogin, fmt.Sprintf(contextCephalonCloudLogin, v.Code))
}

// SendMailCephalonCloudModifyPwd Cephalon Cloud 修改密码
func (v *VerifyMail) SendMailCephalonCloudModifyPwd() error {
	return SendMailAli(v.Username, v.Password, v.SendMail, v.ReceiveMail, types.SendNameCephalon, titleCephalonCloudModifyPwd, fmt.Sprintf(contextCephalonCloudModifyPwd, v.Code))
}

// SendMailCephalonCloudBindWechat Cephalon Cloud 绑定微信
func (v *VerifyMail) SendMailCephalonCloudBindWechat() error {
	return SendMailAli(v.Username, v.Password, v.SendMail, v.ReceiveMail, types.SendNameCephalon, titleCephalonCloudBindWechat, fmt.Sprintf(contextCephalonCloudBindWechat, v.Code))
}

// SendMailCephalonCoreRegister Cephalon Core 注册
func (v *VerifyMail) SendMailCephalonCoreRegister() error {
	return SendMailAli(v.Username, v.Password, v.SendMail, v.ReceiveMail, types.SendNameCephalonCore, titleCephalonCoreRegister, fmt.Sprintf(contextCephalonCoreRegister, v.Code))
}

// SendMailCephalonCoreLogin Cephalon Core 登录
func (v *VerifyMail) SendMailCephalonCoreLogin() error {
	return SendMailAli(v.Username, v.Password, v.SendMail, v.ReceiveMail, types.SendNameCephalonCore, titleCephalonCoreLogin, fmt.Sprintf(contextCephalonCoreLogin, v.Code))
}

// SendMailCephalonCoreModifyPwd Cephalon Core 修改密码
func (v *VerifyMail) SendMailCephalonCoreModifyPwd() error {
	return SendMailAli(v.Username, v.Password, v.SendMail, v.ReceiveMail, types.SendNameCephalonCore, titleCephalonCoreModifyPwd, fmt.Sprintf(contextCephalonCoreModifyPwd, v.Code))
}

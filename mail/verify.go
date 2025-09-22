package mail

import (
	"fmt"
	"github.com/cyrus-go/notify/types"
)

const (
	//titleCephalonCloudRegister   = "端脑云注册验证"
	//titleCephalonCloudLogin      = "端脑云登录验证"
	//titleCephalonCloudModifyPwd  = "端脑云修改密码验证"
	//titleCephalonCloudBindWechat = "端脑云绑定微信验证"
	titleCephalonCloudRegister   = "Welcome to Cephalon AI"
	titleCephalonCloudLogin      = "Cephalon AI Login Verification"
	titleCephalonCloudModifyPwd  = "Cephalon AI Email Verification"
	titleCephalonCloudBindWechat = "Cephalon AI Email Verification"

	titleCephalonCoreRegister  = "端脑小程序注册验证"
	titleCephalonCoreLogin     = "端脑小程序登录验证"
	titleCephalonCoreModifyPwd = "端脑小程序修改密码验证"

	//contextCephalonCloudRegister   = "【端脑科技】欢迎注册 Cephalon Cloud，本次注册的验证码为 %s，5分钟内有效"
	//contextCephalonCloudLogin      = "【端脑科技】欢迎登陆 Cephalon Cloud，本次登陆的验证码为 %s，5分钟内有效"
	//contextCephalonCloudModifyPwd  = "【端脑科技】您正在修改 Cephalon Cloud 密码，本次验证码为 %s，5分钟内有效"
	//contextCephalonCloudBindWechat = "【端脑科技】您正在使用微信绑定 Cephalon Cloud，本次验证码为 %s，5分钟内有效"
	contextCephalonCloudRegister   = "【Cephalon AI】Welcome to Cephalon AI registration! Your verification code is %s, valid for 5 minutes."
	contextCephalonCloudLogin      = "【Cephalon AI】Welcome to Cephalon AI! Your login verification code is %s, valid for 5 minutes."
	contextCephalonCloudModifyPwd  = "【Cephalon AI】You are changing your Cephalon AI password. Your verification code is %s, valid for 5 minutes."
	contextCephalonCloudBindWechat = "【Cephalon AI】You are binding Cephalon AI with WeChat. Your verification code is %s, valid for 5 minutes."

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

type VerifyServiceMail struct {
	Username       string
	Password       string
	Host           string
	To             string
	Cc             string
	Bcc            string
	Date           string
	MailType       string
	ReplyToAddress string
	Code           string
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

func NewVerifyServiceMail(username, password, to, date, code string) *VerifyServiceMail {
	return &VerifyServiceMail{
		Username: username,
		Password: password,
		To:       to,
		Date:     date,
		Code:     code,
		Host:     "smtpdm.aliyun.com:80", // 阿里云固定 host
		//Cc:             cc,  // 抄送
		//Bcc:            bcc, // 密送
		//ReplyToAddress: replyToAddress,
		//MailType: mailType,
	}
}

func (v *VerifyMail) SendMailCephalonCloudRegisterTencent() error {
	return SendMailTencent(v.Username, v.Password, v.SendMail, v.ReceiveMail, types.SendNameCephalon, titleCephalonCloudRegister, fmt.Sprintf(contextCephalonCloudRegister, v.Code))
}

func (v *VerifyMail) SendMailCephalonCloudLoginTencent() error {
	return SendMailTencent(v.Username, v.Password, v.SendMail, v.ReceiveMail, types.SendNameCephalon, titleCephalonCloudLogin, fmt.Sprintf(contextCephalonCloudLogin, v.Code))
}

func (v *VerifyMail) SendMailCephalonCloudModifyPwdTencent() error {
	return SendMailTencent(v.Username, v.Password, v.SendMail, v.ReceiveMail, types.SendNameCephalon, titleCephalonCloudModifyPwd, fmt.Sprintf(contextCephalonCloudModifyPwd, v.Code))
}

func (v *VerifyMail) SendMailCephalonCloudBindWechatTencent() error {

	return SendMailTencent(v.Username, v.Password, v.SendMail, v.ReceiveMail, types.SendNameCephalon, titleCephalonCloudBindWechat, fmt.Sprintf(contextCephalonCloudBindWechat, v.Code))
}

func (v *VerifyMail) SendMailCephalonCoreRegisterTencent() error {
	return SendMailTencent(v.Username, v.Password, v.SendMail, v.ReceiveMail, types.SendNameCephalonCore, titleCephalonCoreRegister, fmt.Sprintf(contextCephalonCoreRegister, v.Code))
}

func (v *VerifyMail) SendMailCephalonCoreLoginTencent() error {
	return SendMailTencent(v.Username, v.Password, v.SendMail, v.ReceiveMail, types.SendNameCephalonCore, titleCephalonCoreLogin, fmt.Sprintf(contextCephalonCoreLogin, v.Code))
}

func (v *VerifyMail) SendMailCephalonCoreModifyPwdTencent() error {
	return SendMailTencent(v.Username, v.Password, v.SendMail, v.ReceiveMail, types.SendNameCephalonCore, titleCephalonCoreModifyPwd, fmt.Sprintf(contextCephalonCoreModifyPwd, v.Code))
}

func (v *VerifyServiceMail) SendMailCephalonCloudRegisterTencent() error {
	return SendMailServiceTencent(v.Username, v.Password, v.Host, titleCephalonCloudRegister, v.Date, fmt.Sprintf(contextCephalonCloudRegister, v.Code), "text", "", []string{v.To}, []string{types.CcEmailTencent}, []string{})
}

func (v *VerifyServiceMail) SendMailCephalonCloudLoginTencent() error {
	return SendMailServiceTencent(v.Username, v.Password, v.Host, titleCephalonCloudLogin, v.Date, fmt.Sprintf(contextCephalonCloudLogin, v.Code), "text", "", []string{v.To}, []string{types.CcEmailTencent}, []string{})
}

func (v *VerifyServiceMail) SendMailCephalonCloudModifyPwdTencent() error {
	return SendMailServiceTencent(v.Username, v.Password, v.Host, titleCephalonCloudModifyPwd, v.Date, fmt.Sprintf(contextCephalonCloudModifyPwd, v.Code), "text", "", []string{v.To}, []string{types.CcEmailTencent}, []string{})
}

func (v *VerifyServiceMail) SendMailCephalonCloudBindWechatTencent() error {
	return SendMailServiceTencent(v.Username, v.Password, v.Host, titleCephalonCloudBindWechat, v.Date, fmt.Sprintf(contextCephalonCloudBindWechat, v.Code), "text", "", []string{v.To}, []string{types.CcEmailTencent}, []string{})
}

func (v *VerifyServiceMail) SendMailCephalonCoreRegisterTencent() error {
	return SendMailServiceTencent(v.Username, v.Password, v.Host, titleCephalonCoreRegister, v.Date, fmt.Sprintf(contextCephalonCoreRegister, v.Code), "text", "", []string{v.To}, []string{types.CcEmailTencent}, []string{})
}

func (v *VerifyServiceMail) SendMailCephalonCoreLoginTencent() error {
	return SendMailServiceTencent(v.Username, v.Password, v.Host, titleCephalonCoreLogin, v.Date, fmt.Sprintf(contextCephalonCoreLogin, v.Code), "text", "", []string{v.To}, []string{types.CcEmailTencent}, []string{})
}

func (v *VerifyServiceMail) SendMailCephalonCoreModifyPwdTencent() error {
	return SendMailServiceTencent(v.Username, v.Password, v.Host, titleCephalonCoreModifyPwd, v.Date, fmt.Sprintf(contextCephalonCoreModifyPwd, v.Code), "text", "", []string{v.To}, []string{types.CcEmailTencent}, []string{})
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

// SendMailCephalonCloudRegister Cephalon Cloud 注册
func (v *VerifyServiceMail) SendMailCephalonCloudRegister() error {
	return SendMailServiceAli(v.Username, v.Password, v.Host, titleCephalonCloudRegister, v.Date, fmt.Sprintf(contextCephalonCloudRegister, v.Code), "text", "", []string{v.To}, []string{types.CcEmail}, []string{})
}

// SendMailCephalonCloudLogin Cephalon Cloud 登录
func (v *VerifyServiceMail) SendMailCephalonCloudLogin() error {
	return SendMailServiceAli(v.Username, v.Password, v.Host, titleCephalonCloudLogin, v.Date, fmt.Sprintf(contextCephalonCloudLogin, v.Code), "text", "", []string{v.To}, []string{types.CcEmail}, []string{})
}

// SendMailCephalonCloudModifyPwd Cephalon Cloud 修改密码
func (v *VerifyServiceMail) SendMailCephalonCloudModifyPwd() error {
	return SendMailServiceAli(v.Username, v.Password, v.Host, titleCephalonCloudModifyPwd, v.Date, fmt.Sprintf(contextCephalonCloudModifyPwd, v.Code), "text", "", []string{v.To}, []string{types.CcEmail}, []string{})
}

// SendMailCephalonCloudBindWechat Cephalon Cloud 绑定微信
func (v *VerifyServiceMail) SendMailCephalonCloudBindWechat() error {
	return SendMailServiceAli(v.Username, v.Password, v.Host, titleCephalonCloudBindWechat, v.Date, fmt.Sprintf(contextCephalonCloudBindWechat, v.Code), "text", "", []string{v.To}, []string{types.CcEmail}, []string{})
}

// SendMailCephalonCoreRegister Cephalon Core 注册
func (v *VerifyServiceMail) SendMailCephalonCoreRegister() error {
	return SendMailServiceAli(v.Username, v.Password, v.Host, titleCephalonCoreRegister, v.Date, fmt.Sprintf(contextCephalonCoreRegister, v.Code), "text", "", []string{v.To}, []string{types.CcEmail}, []string{})
}

// SendMailCephalonCoreLogin Cephalon Core 登录
func (v *VerifyServiceMail) SendMailCephalonCoreLogin() error {
	return SendMailServiceAli(v.Username, v.Password, v.Host, titleCephalonCoreLogin, v.Date, fmt.Sprintf(contextCephalonCoreLogin, v.Code), "text", "", []string{v.To}, []string{types.CcEmail}, []string{})
}

// SendMailCephalonCoreModifyPwd Cephalon Core 修改密码
func (v *VerifyServiceMail) SendMailCephalonCoreModifyPwd() error {
	return SendMailServiceAli(v.Username, v.Password, v.Host, titleCephalonCoreModifyPwd, v.Date, fmt.Sprintf(contextCephalonCoreModifyPwd, v.Code), "text", "", []string{v.To}, []string{types.CcEmail}, []string{})
}

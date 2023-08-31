package sms

import (
	"github.com/cyrus-go/notify/pkg/tools"
	"github.com/pkg/errors"
	"net/url"
)

const (
	TplIdForCephalonCoreRegister  = "5706256" // Cephalon Core 注册模板编号
	TplIdForCephalonCoreLogin     = "5706258" // Cephalon Core 登录模板编号
	TplIdForCephalonCoreModify    = "5706266" // Cephalon Core 修改密码模板编号
	TplIdForYuanHuiRegister       = "5706320" // 元绘注册模板编号
	TplIdForCephalonCloudModify   = "5720264" // Cephalon Cloud 修改密码模板编号
	TplIdForCephalonCloudLogin    = "5720262" // Cephalon Cloud 登录模板编号
	TplIdForCephalonCloudRegister = "5720260" // Cephalon Cloud 注册模板编号
	TplIdForCephalonBindWechat    = "5749864" // Cephalon Cloud 绑定微信模板编号
	//apikey                   = "14b1bbedae4ba4403778f551948e8026"
)

type VerifySms struct {
	APIKey string
	Mobile string
	Code   string
}

func NewVerifySms(apikey, mobile, code string) (*VerifySms, error) {
	// 手机号验证
	if !tools.VerifyMobile(mobile) {
		return nil, errors.New("invalid phone num")
	}
	return &VerifySms{
		APIKey: apikey,
		Mobile: mobile,
		Code:   code,
	}, nil
}

// SendSmsRegisterCephalonCore Cephalon core 注册短信验证
func (v *VerifySms) SendSmsRegisterCephalonCore() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdForCephalonCoreRegister, tplValue)
}

// SendSmsLoginCephalonCore Cephalon core 登录短信验证
func (v *VerifySms) SendSmsLoginCephalonCore() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdForCephalonCoreLogin, tplValue)
}

// SendSmsModifyPwdCephalonCore Cephalon core 修改密码短信验证
func (v *VerifySms) SendSmsModifyPwdCephalonCore() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdForCephalonCoreModify, tplValue)
}

// SendSmsRegisterYuanHui 元绘 登录短信验证
func (v *VerifySms) SendSmsRegisterYuanHui() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdForYuanHuiRegister, tplValue)
}

// SendSmsRegisterCephalon Cephalon 注册短信验证
func (v *VerifySms) SendSmsRegisterCephalon() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdForCephalonCloudRegister, tplValue)
}

// SendSmsLoginCephalon Cephalon 登录短信验证
func (v *VerifySms) SendSmsLoginCephalon() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdForCephalonCloudLogin, tplValue)
}

// SendSmsModifyPwdCephalon Cephalon 修改密码短信验证
func (v *VerifySms) SendSmsModifyPwdCephalon() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdForCephalonCloudModify, tplValue)
}

// SendSmsBindWechatCephalon Cephalon 绑定微信短信验证
func (v *VerifySms) SendSmsBindWechatCephalon() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdForCephalonBindWechat, tplValue)
}

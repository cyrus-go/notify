package sms

import (
	"github.com/cyrus-go/notify/pkg/tools"
	"github.com/pkg/errors"
	"net/url"
)

const (
	TplIdForCephalonRegister = "5706256" // Cephalon 注册模板编号
	TplIdForCephalonLogin    = "5706258" // Cephalon 登录模板编号
	TplIdForCephalonModify   = "5706266" // Cephalon 修改密码模板编号
	TplIdForYuanHuiRegister  = "5706320" // 元绘注册模板编号
	//apikey                   = "14b1bbedae4ba4403778f551948e8026"
)

type VerifySms struct {
	Mobile string
	Code   string
}

func NewVerifySms(mobile, code string) (*VerifySms, error) {
	// 手机号验证
	if !tools.VerifyMobile(mobile) {
		return nil, errors.New("invalid phone num")
	}
	return &VerifySms{
		Mobile: mobile,
		Code:   code,
	}, nil
}

// SendSmsRegisterCephalon Cephalon 注册短信验证
func (v *VerifySms) SendSmsRegisterCephalon() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.Mobile, TplIdForCephalonRegister, tplValue)
}

// SendSmsLoginCephalon Cephalon 登录短信验证
func (v *VerifySms) SendSmsLoginCephalon() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.Mobile, TplIdForCephalonLogin, tplValue)
}

// SendSmsModifyPwdCephalon Cephalon 修改密码短信验证
func (v *VerifySms) SendSmsModifyPwdCephalon() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.Mobile, TplIdForCephalonModify, tplValue)
}

// SendSmsRegisterYuanHui 元绘 登录短信验证
func (v *VerifySms) SendSmsRegisterYuanHui() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.Mobile, TplIdForYuanHuiRegister, tplValue)
}

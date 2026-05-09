package sms

type Verify interface {
	SendSmsRegisterCephalon() (smsID string, err error)
	SendSmsLoginCephalon() (smsID string, err error)
	SendSmsModifyPwdCephalon() (smsID string, err error)
	SendSmsBindWechatCephalon() (smsID string, err error)
	SendSmsRegisterCephalonCore() (smsID string, err error)
	SendSmsLoginCephalonCore() (smsID string, err error)
	SendSmsModifyPwdCephalonCore() (smsID string, err error)
}

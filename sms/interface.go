package sms

type Verify interface {
	SendSmsRegisterCephalon() (err error)
	SendSmsLoginCephalon() (err error)
	SendSmsModifyPwdCephalon() (err error)
	SendSmsBindWechatCephalon() (err error)
	SendSmsRegisterCephalonCore() (err error)
	SendSmsLoginCephalonCore() (err error)
	SendSmsModifyPwdCephalonCore() (err error)
}

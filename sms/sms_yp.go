package sms

import (
	"net/url"
)

// SendSmsRegisterCephalonCore Cephalon core 注册短信验证
func (v *VerifySms) SendSmsRegisterCephalonCore() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdYPForCephalonCoreRegister, tplValue)
}

// SendSmsLoginCephalonCore Cephalon core 登录短信验证
func (v *VerifySms) SendSmsLoginCephalonCore() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdYPForCephalonCoreLogin, tplValue)
}

// SendSmsModifyPwdCephalonCore Cephalon core 修改密码短信验证
func (v *VerifySms) SendSmsModifyPwdCephalonCore() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdYPForCephalonCoreModify, tplValue)
}

// SendSmsRegisterYuanHui 元绘 注册短信验证
func (v *VerifySms) SendSmsRegisterYuanHui() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdYPForYuanHuiRegister, tplValue)
}

// SendSmsYuanHuiLogin 元绘 登录短信验证
func (v *VerifySms) SendSmsYuanHuiLogin() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdYPForYuanHuiLogin, tplValue)
}

// SendSmsYuanHuiBindWechat 元绘 绑定微信短信验证
func (v *VerifySms) SendSmsYuanHuiBindWechat() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdYPForYuanHuiBindWechat, tplValue)
}

// SendSmsRegisterCephalon Cephalon 注册短信验证
func (v *VerifySms) SendSmsRegisterCephalon() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdYPForCephalonCloudRegister, tplValue)
}

// SendSmsLoginCephalon Cephalon 登录短信验证
func (v *VerifySms) SendSmsLoginCephalon() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdYPForCephalonCloudLogin, tplValue)
}

// SendSmsModifyPwdCephalon Cephalon 修改密码短信验证
func (v *VerifySms) SendSmsModifyPwdCephalon() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdYPForCephalonCloudModify, tplValue)
}

// SendSmsBindWechatCephalon Cephalon 绑定微信短信验证
func (v *VerifySms) SendSmsBindWechatCephalon() (err error) {
	tplValue := url.Values{"#code#": {v.Code}}
	return sendSms(v.APIKey, v.Mobile, TplIdYPForCephalonBindWechat, tplValue)
}

// SendSmsCephalonRenewalBalanceNotEnough Cephalon 自动续费余额不足短信
func (v *NotifySms) SendSmsCephalonRenewalBalanceNotEnough() (err error) {
	return sendSms(v.APIKey, v.Mobile, TplIdYPForCephalonRenewalBalanceNotEnough, nil)
}

// SendSmsCephalonMissionExpired Cephalon 任务到期提醒短信
func (v *NotifySms) SendSmsCephalonMissionExpired() (err error) {
	return sendSms(v.APIKey, v.Mobile, TplIdYPForCephalonMissionExpired, nil)
}

// SendSmsCephalonMissionRunningTimeWarningSixHour Cephalon 任务运行时长提醒 6 小时
func (v *NotifySms) SendSmsCephalonMissionRunningTimeWarningSixHour() (err error) {
	tplValue := url.Values{"#time#": {WarningTimeParamSixHour}}
	return sendSms(v.APIKey, v.Mobile, TplIdForCephalonMissionRunningTimeWarning, tplValue)
}

// SendSmsCephalonMissionRunningTimeWarningTwelveHour Cephalon 任务运行时长提醒 12 小时
func (v *NotifySms) SendSmsCephalonMissionRunningTimeWarningTwelveHour() (err error) {
	tplValue := url.Values{"#time#": {WarningTimeParamTwelveHour}}
	return sendSms(v.APIKey, v.Mobile, TplIdForCephalonMissionRunningTimeWarning, tplValue)
}

// SendSmsCephalonMissionRunningTimeWarningTwentyFourHour Cephalon 任务运行时长提醒 24 小时
func (v *NotifySms) SendSmsCephalonMissionRunningTimeWarningTwentyFourHour() (err error) {
	tplValue := url.Values{"#time#": {WarningTimeParamTwentyFourHour}}
	return sendSms(v.APIKey, v.Mobile, TplIdForCephalonMissionRunningTimeWarning, tplValue)
}

func (v *NotifySms) SendSmsCustom(templateID string, tplValue url.Values) (err error) {
	return sendSms(v.APIKey, v.Mobile, templateID, tplValue)
}

// SendSmsCephalonDeviceDriveAbnormal Cephalon 节点设备驱动异常提醒
func (v *NotifySmsDeviceDriveAbnormal) SendSmsCephalonDeviceDriveAbnormal() (err error) {
	tplValue := url.Values{"#name#": {v.DeviceName}, "#time#": {v.Time}}
	return sendSms(v.APIKey, v.Mobile, TplIdForCephalonDeviceDriveAbnormal, tplValue)
}

package sms

// SendTencentSmsCephalonCoreRegister Cephalon core 注册短信验证 - 腾讯云短信
func (v *VerifySmsTencent) SendTencentSmsCephalonCoreRegister() ([]byte, error) {
	return sendSmsTencent(v.SecretId, v.SecretKey, v.SdkAppId, TplIdTencentForCephalonCoreRegister, v.Phones, []string{v.Code})
}

// SendTencentSmsCephalonCoreLogin Cephalon core 登录短信验证 - 腾讯云短信
func (v *VerifySmsTencent) SendTencentSmsCephalonCoreLogin() ([]byte, error) {
	return sendSmsTencent(v.SecretId, v.SecretKey, v.SdkAppId, TplIdTencentForCephalonCoreLogin, v.Phones, []string{v.Code})
}

// SendTencentSmsCephalonCoreModifyPwd Cephalon core 修改密码短信验证 - 腾讯云短信
func (v *VerifySmsTencent) SendTencentSmsCephalonCoreModifyPwd() ([]byte, error) {
	return sendSmsTencent(v.SecretId, v.SecretKey, v.SdkAppId, TplIdTencentForCephalonCoreModify, v.Phones, []string{v.Code})
}

// SendTencentSmsCephalonRegister Cephalon 注册短信验证 - 腾讯云短信
func (v *VerifySmsTencent) SendTencentSmsCephalonRegister() ([]byte, error) {
	return sendSmsTencent(v.SecretId, v.SecretKey, v.SdkAppId, TplIdTencentForCephalonCloudRegister, v.Phones, []string{v.Code})
}

// SendTencentSmsCephalonLogin Cephalon 登录短信验证 - 腾讯云短信
func (v *VerifySmsTencent) SendTencentSmsCephalonLogin() ([]byte, error) {
	return sendSmsTencent(v.SecretId, v.SecretKey, v.SdkAppId, TplIdTencentForCephalonCloudLogin, v.Phones, []string{v.Code})
}

// SendTencentSmsCephalonModifyPwd Cephalon 修改密码短信验证 - 腾讯云短信
func (v *VerifySmsTencent) SendTencentSmsCephalonModifyPwd() ([]byte, error) {
	return sendSmsTencent(v.SecretId, v.SecretKey, v.SdkAppId, TplIdTencentForCephalonCloudModify, v.Phones, []string{v.Code})
}

// SendTencentSmsCephalonBindWechat Cephalon 绑定微信短信验证 - 腾讯云短信
func (v *VerifySmsTencent) SendTencentSmsCephalonBindWechat() ([]byte, error) {
	return sendSmsTencent(v.SecretId, v.SecretKey, v.SdkAppId, TplIdTencentForCephalonBindWechat, v.Phones, []string{v.Code})
}

// SendTencentSmsCephalonRenewalBalanceNotEnough Cephalon 自动续费余额不足短信 - 腾讯云短信
func (v *NotifySmsTencent) SendTencentSmsCephalonRenewalBalanceNotEnough() ([]byte, error) {
	return sendSmsTencent(v.SecretId, v.SecretKey, v.SdkAppId, TplIdTencentForCephalonRenewalBalanceNotEnough, v.Phones, []string{})
}

// SendTencentSmsCephalonMissionExpired Cephalon 任务到期提醒短信 - 腾讯云短信
func (v *NotifySmsTencent) SendTencentSmsCephalonMissionExpired() ([]byte, error) {
	return sendSmsTencent(v.SecretId, v.SecretKey, v.SdkAppId, TplIdTencentForCephalonMissionExpired, v.Phones, []string{})
}

// SendTencentSmsCephalonMissionRunningTimeWarning Cephalon 任务运行时长提醒 - 腾讯云短信
func (v *NotifySmsWithParamTencent) SendTencentSmsCephalonMissionRunningTimeWarning() ([]byte, error) {
	return sendSmsTencent(v.SecretId, v.SecretKey, v.SdkAppId, TplIdTencentForCephalonMissionRunningTimeWarning, v.Phones, v.Params)
}

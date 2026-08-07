package sms

import (
	"github.com/cyrus-go/notify/pkg/tools"
	"github.com/pkg/errors"
)

type VerifySms struct {
	APIKey string
	Mobile string
	Code   string
}

type NotifySms struct {
	APIKey string
	Mobile string
}

type NotifySmsDeviceDriveAbnormal struct {
	APIKey     string
	Mobile     string
	DeviceName string
	Time       string
}

type VerifySmsTencent struct {
	SdkAppId  string
	SecretId  string
	SecretKey string
	Phones    []string
	Code      string
}

type NotifySmsTencent struct {
	SdkAppId  string
	SecretId  string
	SecretKey string
	Phones    []string
}

// VerifySmsYmrt 亿美软通
type VerifySmsYmrt struct {
	AppID     string
	SecretKey string
	Mobile    string
	Code      string
}

type NotifySmsYmrt struct {
	AppID     string
	SecretKey string
	Mobile    string
}

// VerifySmsBeanli 滨礼网络验证码短信客户端。
type VerifySmsBeanli struct {
	Account  string
	Password string
	Mobile   string
	Code     string
}

// NotifySmsBeanli 滨礼网络通知短信客户端。
type NotifySmsBeanli struct {
	Account  string
	Password string
	Mobile   string
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

func NewNotifySms(apikey, mobile string) (*NotifySms, error) {
	// 手机号验证
	if !tools.VerifyMobile(mobile) {
		return nil, errors.New("invalid phone num")
	}
	return &NotifySms{
		APIKey: apikey,
		Mobile: mobile,
	}, nil
}

func NewNotifySmsDeviceDriveAbnormal(apikey, mobile, deviceName, abnormalTime string) (*NotifySmsDeviceDriveAbnormal, error) {
	// 手机号验证
	if !tools.VerifyMobile(mobile) {
		return nil, errors.New("invalid phone num")
	}
	return &NotifySmsDeviceDriveAbnormal{
		APIKey:     apikey,
		Mobile:     mobile,
		DeviceName: deviceName,
		Time:       abnormalTime,
	}, nil
}

func NewVerifySmsTencent(sdkAppId, secretId, secretKey, code string, phones []string) (*VerifySmsTencent, error) {
	if len(phones) <= 0 {
		return nil, errors.New("phones is a blank list")
	}
	// 手机号验证
	for _, phone := range phones {
		if !tools.VerifyMobile(phone) {
			return nil, errors.New("invalid phone num")
		}
	}
	return &VerifySmsTencent{
		SdkAppId:  sdkAppId,
		SecretId:  secretId,
		SecretKey: secretKey,
		Phones:    phones,
		Code:      code,
	}, nil
}

func NewNotifySmsTencent(sdkAppId, secretId, secretKey string, phones []string) (*NotifySmsTencent, error) {
	if len(phones) <= 0 {
		return nil, errors.New("phones is a blank list")
	}
	// 手机号验证
	for _, phone := range phones {
		if !tools.VerifyMobile(phone) {
			return nil, errors.New("invalid phone num")
		}
	}
	return &NotifySmsTencent{
		SdkAppId:  sdkAppId,
		SecretId:  secretId,
		SecretKey: secretKey,
		Phones:    phones,
	}, nil
}

func NewVerifySmsYmrt(appID, secretKey, mobile, code string) (*VerifySmsYmrt, error) {
	// 手机号验证
	if !tools.VerifyMobile(mobile) {
		return nil, errors.New("invalid phone num")
	}
	return &VerifySmsYmrt{
		AppID:     appID,
		SecretKey: secretKey,
		Mobile:    mobile,
		Code:      code,
	}, nil
}

func NewNotifySmsYmrt(appID, secretKey, mobile string) (*NotifySmsYmrt, error) {
	// 手机号验证
	if !tools.VerifyMobile(mobile) {
		return nil, errors.New("invalid phone num")
	}
	return &NotifySmsYmrt{
		AppID:     appID,
		SecretKey: secretKey,
		Mobile:    mobile,
	}, nil
}

// NewVerifySmsBeanli 创建滨礼网络验证码短信客户端。
func NewVerifySmsBeanli(account, password, mobile, code string) (*VerifySmsBeanli, error) {
	// 滨礼网络文字短信接口接收单个或多个手机号，此客户端与亿美软通保持一致，仅接收单个手机号。
	if !tools.VerifyMobile(mobile) {
		return nil, errors.New("invalid phone num")
	}
	return &VerifySmsBeanli{
		Account:  account,
		Password: password,
		Mobile:   mobile,
		Code:     code,
	}, nil
}

// NewNotifySmsBeanli 创建滨礼网络通知短信客户端。
func NewNotifySmsBeanli(account, password, mobile string) (*NotifySmsBeanli, error) {
	// 手机号验证。
	if !tools.VerifyMobile(mobile) {
		return nil, errors.New("invalid phone num")
	}
	return &NotifySmsBeanli{
		Account:  account,
		Password: password,
		Mobile:   mobile,
	}, nil
}

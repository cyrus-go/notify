package mail

import (
	"github.com/cyrus-go/notify/types"
)

const (
	titleNotifyBalanceWarning    = "余额预警"
	titleNotifyAppExpiredWarning = "应用到期预警"

	contextNotifyBalanceWarning    = "【端脑科技】温馨提示：您正在使用的 AI 应用即将到期，系统显示您的余额不足以进行自动续费，如需继续使用请及时充值噢！到期后应用会自动停用并清除现有设置，为了方便您的使用，请确保在服务到期前进行充值或备份数据。\n    端脑云感谢您的使用，祝您生活愉快！"
	contextNotifyAppExpiredWarning = "【端脑科技】温馨提示：您正在使用的 AI 应用即将到期，如需继续使用请及时续费噢！到期后应用会自动停用并清除现有设置，为了方便您的使用，请确保在服务到期前进行续费或备份数据。\n    端脑云感谢您的使用，祝您生活愉快！"
)

type NotifyMail struct {
	Username    string
	Password    string
	SendMail    string
	ReceiveMail string
}

func NewNotifyMail(username, password, sendMail, receiveMail string) *NotifyMail {
	return &NotifyMail{
		Username:    username,
		Password:    password,
		SendMail:    sendMail,
		ReceiveMail: receiveMail,
	}
}

func (n *NotifyMail) SendEmailWarningBalance() error {
	return SendMailAli(n.Username, n.Password, n.SendMail, n.ReceiveMail, types.SendNameCephalon, titleNotifyBalanceWarning, contextNotifyBalanceWarning)
}

func (n *NotifyMail) SendEmailWarningAppExpired() error {
	return SendMailAli(n.Username, n.Password, n.SendMail, n.ReceiveMail, types.SendNameCephalon, titleNotifyAppExpiredWarning, contextNotifyAppExpiredWarning)
}

// SendEmailCustom 发送自定义邮件
func (n *NotifyMail) SendEmailCustom(sendName, title, content string) error {
	return SendMailAli(n.Username, n.Password, n.SendMail, n.ReceiveMail, sendName, title, content)
}

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

type NotifyServiceMail struct {
	Username       string
	Password       string
	Host           string
	To             string
	Cc             string
	Bcc            string
	Date           string
	MailType       string
	ReplyToAddress string
}

func NewNotifyMail(username, password, sendMail, receiveMail string) *NotifyMail {
	return &NotifyMail{
		Username:    username,
		Password:    password,
		SendMail:    sendMail,
		ReceiveMail: receiveMail,
	}
}

func NewNotifyServiceMail(username, password, to, date string) *NotifyServiceMail {
	return &NotifyServiceMail{
		Username: username,
		Password: password,
		To:       to,
		Date:     date,
		Host:     "smtpdm.aliyun.com:80", // 阿里云固定 host
		//Cc:             cc,  // 抄送
		//Bcc:            bcc, // 密送
		//ReplyToAddress: replyToAddress,
		//MailType: mailType,
	}
}
func NewNotifyServiceMailTencent(username, password, to, date string) *NotifyServiceMail {
	return &NotifyServiceMail{
		Username: username,
		Password: password,
		To:       to,
		Date:     date,
		Host:     "gz-smtp.qcloudmail.com:465", // 阿里云固定 host
		//Cc:             cc,  // 抄送
		//Bcc:            bcc, // 密送
		//ReplyToAddress: replyToAddress,
		//MailType: mailType,
	}
}
func NewNotifyServiceMailSelf(username, password, to, date string) *NotifyServiceMail {
	return &NotifyServiceMail{
		Username: username,
		Password: password,
		To:       to,
		Date:     date,
		Host:     "mail.zk-agent.ai:587", // 阿里云固定 host
		//Cc:             cc,  // 抄送
		//Bcc:            bcc, // 密送
		//ReplyToAddress: replyToAddress,
		//MailType: mailType,
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

func (n *NotifyServiceMail) SendEmailWarningBalance() error {
	return SendMailServiceAli(n.Username, n.Password, n.Host, titleNotifyBalanceWarning, n.Date, contextNotifyBalanceWarning, "text", "", []string{n.To}, []string{types.CcEmail}, []string{})
}

func (n *NotifyServiceMail) SendEmailWarningAppExpired() error {
	return SendMailServiceAli(n.Username, n.Password, n.Host, titleNotifyAppExpiredWarning, n.Date, contextNotifyAppExpiredWarning, "text", "", []string{n.To}, []string{types.CcEmail}, []string{})
}

// SendEmailCustom 发送自定义邮件
func (n *NotifyServiceMail) SendEmailCustom(title, content string) error {
	return SendMailServiceAli(n.Username, n.Password, n.Host, title, n.Date, content, "text", "", []string{n.To}, []string{types.CcEmail}, []string{})
}

func (n *NotifyMail) SendEmailWarningBalanceTencent() error {
	return SendMailTencent(n.Username, n.Password, n.SendMail, n.ReceiveMail, types.SendNameCephalon, titleNotifyBalanceWarning, contextNotifyBalanceWarning)
}

func (n *NotifyMail) SendEmailWarningAppExpiredTencent() error {
	return SendMailTencent(n.Username, n.Password, n.SendMail, n.ReceiveMail, types.SendNameCephalon, titleNotifyAppExpiredWarning, contextNotifyAppExpiredWarning)
}

// SendEmailCustom 发送自定义邮件
func (n *NotifyMail) SendEmailCustomTencent(sendName, title, content string) error {
	return SendMailTencent(n.Username, n.Password, n.SendMail, n.ReceiveMail, sendName, title, content)
}

func (n *NotifyServiceMail) SendEmailWarningBalanceTencent() error {
	return SendMailServiceTencent(n.Username, n.Password, n.Host, titleNotifyBalanceWarning, n.Date, contextNotifyBalanceWarning, "text", "", []string{n.To}, []string{types.CcEmailTencent}, []string{})
}

func (n *NotifyServiceMail) SendEmailWarningAppExpiredTencent() error {
	return SendMailServiceTencent(n.Username, n.Password, n.Host, titleNotifyAppExpiredWarning, n.Date, contextNotifyAppExpiredWarning, "text", "", []string{n.To}, []string{types.CcEmailTencent}, []string{})
}

// SendEmailCustom 发送自定义邮件
func (n *NotifyServiceMail) SendEmailCustomTencent(title, content string) error {
	return SendMailServiceTencent(n.Username, n.Password, n.Host, title, n.Date, content, "text", "", []string{n.To}, []string{types.CcEmailTencent}, []string{})
}

func (n *NotifyServiceMail) SendEmailCustomSelf(title, content string) error {
	return SendMailServiceSelf(n.Username, n.Password, n.Host, title, n.Date, content, "text", "", []string{n.To}, []string{types.CcSelfMail}, []string{})
}

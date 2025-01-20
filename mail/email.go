package mail

import (
	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
	"log"
	"mime"
	"net/smtp"
	"strings"
)

const (
	EmailAliHost = "smtp.mxhichina.com"
	EmailAliPort = 465
)

/*
   QQ 邮箱：
   SMTP 服务器地址：smtp.qq.com（SSL协议端口：465/994 | 非SSL协议端口：25）
   163 邮箱：
   SMTP 服务器地址：smtp.163.com（端口：25）
   阿里邮箱：
   SMTP 服务器地址：smtp.mxhichina.com（端口：465）
*/

// SendMailAli 用阿里邮箱发送邮件(mailFrom: 发件人, mailTo: 收件人, title: 标题, content: 内容)
func SendMailAli(username, password, mailFrom, mailTo, sendName, title, content string) error {
	// 初始化
	var conn = &gomail.Dialer{}
	// 创建一个message
	m := gomail.NewMessage()

	conn = gomail.NewDialer(EmailAliHost, EmailAliPort, username, password)
	// 添加发件人信息
	m.SetHeader("From", mime.QEncoding.Encode("UTF-8", sendName)+"<"+mailFrom+">")

	// 构建邮件信息
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", content)

	// 发送
	if err := conn.DialAndSend(m); err != nil {
		return errors.Errorf("To: %s ## Send Email Failed! Err: %v", mailTo, err)
	}

	log.Println("To:", mailTo, "##", "Send Email Successfully!")
	return nil
}

// SendMailServiceAli 发送邮件服务
// user: 用户名, password: 密码, host: 邮箱服务器地址, subject: 主题, date: 发送时间, body: 内容, mailType: html or text, replyToAddress: 回复地址, to: 收件人, cc: 抄送人, bcc: 密送人
func SendMailServiceAli(user, password, host, subject, date, body, mailType, replyToAddress string, to, cc, bcc []string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	if mailType == "html" {
		contentType = "Content-Type: text/" + mailType + "; charset=UTF-8"
	} else {
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	ccAddress := strings.Join(cc, ";")
	bccAddress := strings.Join(bcc, ";")
	toAddress := strings.Join(to, ";")
	msg := []byte("To: " + toAddress + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\nDate: " + date + "\r\nReply-To: " + replyToAddress + "\r\nCc: " + ccAddress + "\r\nBcc: " + bccAddress + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := MergeSlice(to, cc)
	sendTo = MergeSlice(sendTo, bcc)
	return smtp.SendMail(host, auth, user, sendTo, msg)
}

func MergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}

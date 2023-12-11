package mail

import (
	"github.com/pkg/errors"
	"gopkg.in/gomail.v2"
	"log"
	"mime"
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

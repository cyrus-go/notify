package test

import (
	"fmt"
	"github.com/cyrus-go/notify/mail"
	"testing"
	"time"
)

func TestMailRegisterCephalonCloud(t *testing.T) {
	//username := "your@mail.com"
	//password := "your_pwd"
	//to := "to@mail.com"
	username := "youruser"
	password := "yourpasswd"
	to := "touser"
	date := fmt.Sprintf("%s", time.Now().Format(time.RFC1123Z))
	code := "223344"
	verifyMail := mail.NewVerifyServiceMail(username, password, to, date, code)
	if err := verifyMail.SendMailCephalonCloudRegister(); err != nil {
		t.Fatalf("send register cephalon mail failed, err: %v", err)
	}
}

func TestMailRegisterCephalonCore(t *testing.T) {
	username := "your@mail.com"
	password := "your_pwd"
	sendMail := "your@mail.com"
	receiveMail := "receive@mail.com"
	code := "223344"
	verifyMail := mail.NewVerifyMail(username, password, sendMail, receiveMail, code)
	if err := verifyMail.SendMailCephalonCoreRegister(); err != nil {
		t.Fatalf("send register cephalon core mail failed, err: %v", err)
	}
}

func TestMailWarningBalance(t *testing.T) {
	//username := "your@mail.com"
	//password := "your_pwd"
	//sendMail := "your@mail.com"
	//receiveMail := "receive@mail.com"
	username := "service@cephalon.tech"
	password := "tWrlNuH5@123"
	to := "hyk130406@icloud.com"
	date := fmt.Sprintf("%s", time.Now().Format(time.RFC1123Z))
	notifyMail := mail.NewNotifyServiceMail(username, password, to, date)
	if err := notifyMail.SendEmailWarningBalance(); err != nil {
		t.Fatalf("send balance warning mail failed, err: %v", err)
	}
}

func TestMailWarningAppExpired(t *testing.T) {
	username := "service@cephalon.tech"
	password := "tWrlNuH5@123"
	sendMail := "service@cephalon.tech"
	receiveMail := "hyk130406@icloud.com"
	notifyMail := mail.NewNotifyMail(username, password, sendMail, receiveMail)
	if err := notifyMail.SendEmailWarningAppExpired(); err != nil {
		t.Fatalf("send app expired warning mail failed, err: %v", err)
	}
}

func TestMailSendSelfMail(t *testing.T) {
	username := "service@cephalon.vip"
	password := "qcLsV2VAXn_MP5Z"
	receiveMail := "2131939742@qq.com"
	notifyMail := mail.NewNotifyServiceMailSelf(username, password, receiveMail, fmt.Sprintf("%s", time.Now().Format(time.RFC1123Z)))
	if err := notifyMail.SendEmailCustomSelf("nihao", "123123"); err != nil {
		t.Fatalf("send self mail failed, err: %v", err)
	}
}

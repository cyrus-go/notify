package test

import (
	"github.com/cyrus-go/notify/mail"
	"testing"
)

func TestMailRegisterCephalonCloud(t *testing.T) {
	username := "your@mail.com"
	password := "your_pwd"
	sendMail := "your@mail.com"
	receiveMail := "receive@mail.com"
	code := 223344
	verifyMail := mail.NewVerifyMail(username, password, sendMail, receiveMail, code)
	if err := verifyMail.SendMailCephalonCloudRegister(); err != nil {
		t.Fatalf("send register cephalon mail failed, err: %v", err)
	}
}

func TestMailRegisterCephalonCore(t *testing.T) {
	username := "your@mail.com"
	password := "your_pwd"
	sendMail := "your@mail.com"
	receiveMail := "receive@mail.com"
	code := 223344
	verifyMail := mail.NewVerifyMail(username, password, sendMail, receiveMail, code)
	if err := verifyMail.SendMailCephalonCoreRegister(); err != nil {
		t.Fatalf("send register cephalon core mail failed, err: %v", err)
	}
}

func TestMailWarningBalance(t *testing.T) {
	username := "your@mail.com"
	password := "your_pwd"
	sendMail := "your@mail.com"
	receiveMail := "receive@mail.com"
	notifyMail := mail.NewNotifyMail(username, password, sendMail, receiveMail)
	if err := notifyMail.SendEmailWarningBalance(); err != nil {
		t.Fatalf("send balance warning mail failed, err: %v", err)
	}
}

func TestMailWarningAppExpired(t *testing.T) {
	username := "your@mail.com"
	password := "your_pwd"
	sendMail := "your@mail.com"
	receiveMail := "receive@mail.com"
	notifyMail := mail.NewNotifyMail(username, password, sendMail, receiveMail)
	if err := notifyMail.SendEmailWarningAppExpired(); err != nil {
		t.Fatalf("send app expired warning mail failed, err: %v", err)
	}
}

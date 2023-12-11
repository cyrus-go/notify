package test

import (
	"github.com/cyrus-go/notify/mail"
	"testing"
)

func TestMailRegisterCephalonCore(t *testing.T) {
	username := "your@mail.com"
	password := "your_pwd"
	sendMail := "your@mail.com"
	receiveMail := "receive@mail.com"
	code := 223344
	verifyMail := mail.NewVerifyMail(username, password, sendMail, receiveMail, code)
	if err := verifyMail.SendMailCephalonRegister(); err != nil {
		t.Fatalf("send register cephalon mail failed, err: %v", err)
	}
}

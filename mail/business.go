package mail

import "fmt"

const (
	cephalonRegisterContext = "欢迎注册 Cephalon Core，本次注册的验证码为 %d，5分钟内有效"

	sendNameCephalonRegister = "端脑云"

	titleCephalonRegister = "端脑云注册验证"
)

type VerifyMail struct {
	Username    string
	Password    string
	SendMail    string
	ReceiveMail string
	Code        int
}

func NewVerifyMail(username, password, sendMail, receiveMail string, code int) *VerifyMail {
	return &VerifyMail{
		Username:    username,
		Password:    password,
		SendMail:    sendMail,
		ReceiveMail: receiveMail,
		Code:        code,
	}
}

// SendMailCephalonRegister 端脑云注册
func (v *VerifyMail) SendMailCephalonRegister() error {
	return SendMailAli(v.Username, v.Password, v.SendMail, v.ReceiveMail, sendNameCephalonRegister, titleCephalonRegister, fmt.Sprintf(cephalonRegisterContext, v.Code))
}
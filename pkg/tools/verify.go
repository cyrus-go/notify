package tools

import "regexp"

// VerifyMobile 验证是否为真实手机号码
func VerifyMobile(mobile string) bool {
	//regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|193|198|199|(147))\\d{8}$"
	regular := "^1[3-9][0-9]\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

// VerifyEmail 验证是否为真实邮箱
func VerifyEmail(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

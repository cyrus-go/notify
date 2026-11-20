# notify

Go 消息通知工具包。当前提供短信和邮件发送能力。

## 30 秒接入

```bash
go get github.com/cyrus-go/notify@main
```

选择供应商后，按 [接入文档](docs/integration.md) 配置凭据并复制对应示例。

```go
package main

import (
	"log"
	"os"

	"github.com/cyrus-go/notify/sms"
)

func main() {
	client, err := sms.NewVerifySmsYmrt(
		os.Getenv("YMRT_APP_ID"),
		os.Getenv("YMRT_SECRET_KEY"),
		"13800138000",
		"123456",
	)
	if err != nil {
		log.Fatal(err)
	}
	if err = client.SendSmsLoginCephalon(); err != nil {
		log.Fatal(err)
	}
}
```

## 能力

| 类型 | 供应商 | 入口 |
| --- | --- | --- |
| 验证码/通知短信 | 亿美软通 | `sms.NewVerifySmsYmrt`、`sms.NewNotifySmsYmrt` |
| 验证码/通知短信 | 云片 | `sms.NewVerifySms`、`sms.NewNotifySms` |
| 验证码/通知短信 | 腾讯云短信 | `sms.NewVerifySmsTencent`、`sms.NewNotifySmsTencent` |
| 验证码/通知短信 | 滨礼网络 | `sms.NewVerifySmsBeanli`、`sms.NewNotifySmsBeanli` |
| 验证码/通知邮件 | 阿里邮箱/邮件推送 | `mail.NewVerifyMail`、`mail.NewNotifyMail` |
| 验证码/通知邮件 | 腾讯企业邮 | `mail.NewVerifyMail`、`mail.NewNotifyMail` |
| 验证码/通知邮件 | 自建 SMTP | `mail.NewVerifyServiceMailSelf`、`mail.NewNotifyServiceMailSelf` |

## Agent 接入入口

把下面地址交给代码 Agent：

```text
https://github.com/cyrus-go/notify/blob/main/docs/integration.md
```

Agent 应按文档完成四步：

1. 安装 `github.com/cyrus-go/notify`。
2. 向用户询问所选供应商要求的凭据和收件手机号/邮箱。
3. 使用对应 `New...` 构造函数；不要把凭据写入源码。
4. 调用对应发送方法，并把错误返回给业务层。

## 重要限制

- 供应商账户、短信签名、短信模板必须先在供应商控制台开通并审核。
- 亿美软通当前使用代码内置的文字内容，不使用供应商模板 ID。
- 腾讯云的地域、签名和模板 ID 当前在 `sms/consts.go` 与 `sms/send_sms_tencent.go` 内固定；接入其他账号前必须确认这些值适用。
- 滨礼网络接口地址当前为 HTTP：`http://sms.beanli.com:7891/api/v1/send`。部署环境必须允许访问该地址。
- 构造函数会校验手机号；失败时返回 `invalid phone num`。
- 仓库中的 `test` 目录包含真实供应商调用示例。不要直接运行其中测试，否则可能发送真实短信或邮件。

完整配置表、调用示例、错误排查和安全要求见 [docs/integration.md](docs/integration.md)。

## 本地校验

只编译库包，不触发供应商调用：

```bash
go test ./sms ./mail ./types ./pkg/...
```

## 许可证

以仓库实际许可证文件为准。

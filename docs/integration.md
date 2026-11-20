# notify 接入文档

本文面向业务开发者和代码 Agent。所有 API 名称、参数顺序和默认值以当前仓库代码为准。

仓库：<https://github.com/cyrus-go/notify>

模块：`github.com/cyrus-go/notify`

## 1. 安装

```bash
go get github.com/cyrus-go/notify@main
go mod tidy
```

项目要求 Go 1.20 或更高版本。

## 2. 通用接入规则

1. 在供应商控制台创建账户，并完成短信签名、短信模板、发送权限审核。
2. 把凭据放入环境变量或密钥管理系统，不要提交到 Git。
3. 在业务代码启动时读取配置，在发送请求时构造客户端。
4. 处理构造错误和发送错误；不要只记录“发送成功”。
5. 先用供应商测试号码验证，再开放真实用户发送。

手机号由 `sms` 包校验。当前校验要求号码中包含中国大陆 `1[3-9]` 开头的 11 位手机号；腾讯云示例使用 E.164 格式，例如 `+8613800138000`。

## 3. Agent 自动接入协议

如果用户把本文地址交给 Agent，Agent 应先询问：

- 供应商：`ymrt`、`yunpian`、`tencent`、`beanli`、`mail`。
- 用途：验证码、登录、注册、修改密码、微信绑定或自定义通知。
- 接收手机号/邮箱。
- 对应凭据是否已经在本地环境变量中。

Agent 应遵守：

- 使用本节后面列出的构造函数和发送方法，不自行拼接供应商请求。
- 凭据只从环境变量、密钥管理器或用户明确提供的运行时配置读取。
- 不把真实凭据写入源码、日志、测试文件或提交记录。
- 只运行 `go test ./sms ./mail ./types ./pkg/...` 做离线编译校验。`test` 目录中的测试可能触发真实外部发送。

## 4. 亿美软通短信（推荐示例）

### 4.1 用户需要准备

在亿美软通控制台准备：

| 配置 | 用途 | 示例环境变量 |
| --- | --- | --- |
| App ID | 亿美应用标识 | `YMRT_APP_ID` |
| Secret Key | 亿美应用密钥 | `YMRT_SECRET_KEY` |
| 手机号 | 接收短信的手机号 | `SMS_MOBILE` |
| 验证码 | 业务生成的验证码 | 业务运行时变量 |

当前实现访问：`http://www.btom.cn:8080/simpleinter/sendSMS`。

签名算法由 SDK 自动完成：`MD5(AppID + SecretKey + yyyyMMddHHmmss)`，结果为 32 位小写十六进制字符串。业务代码不需要自己计算签名。

亿美实现发送完整文字内容，不传模板 ID。内容必须满足亿美账户的签名和内容审核规则。

### 4.2 验证码短信

```go
package notification

import (
	"fmt"
	"os"

	"github.com/cyrus-go/notify/sms"
)

func SendLoginCode(mobile, code string) error {
	client, err := sms.NewVerifySmsYmrt(
		os.Getenv("YMRT_APP_ID"),
		os.Getenv("YMRT_SECRET_KEY"),
		mobile,
		code,
	)
	if err != nil {
		return fmt.Errorf("create ymrt sms client: %w", err)
	}
	if err = client.SendSmsLoginCephalon(); err != nil {
		return fmt.Errorf("send ymrt login code: %w", err)
	}
	return nil
}
```

内置验证码方法和文字内容：

| 方法 | 用途 |
| --- | --- |
| `SendSmsRegisterCephalon` | Cephalon Cloud 注册 |
| `SendSmsLoginCephalon` | Cephalon Cloud 登录 |
| `SendSmsModifyPwdCephalon` | Cephalon Cloud 修改密码 |
| `SendSmsBindWechatCephalon` | Cephalon Cloud 绑定微信 |
| `SendSmsRegisterCephalonCore` | Cephalon Core 注册 |
| `SendSmsLoginCephalonCore` | Cephalon Core 登录 |
| `SendSmsModifyPwdCephalonCore` | Cephalon Core 修改密码 |

### 4.3 自定义通知短信

```go
func SendNotice(mobile, content string) error {
	client, err := sms.NewNotifySmsYmrt(
		os.Getenv("YMRT_APP_ID"),
		os.Getenv("YMRT_SECRET_KEY"),
		mobile,
	)
	if err != nil {
		return fmt.Errorf("create ymrt sms client: %w", err)
	}
	return client.SendSmsCustom(content)
}
```

### 4.4 本地配置

Linux/macOS：

```bash
export YMRT_APP_ID="your_app_id"
export YMRT_SECRET_KEY="your_secret_key"
export SMS_MOBILE="13800138000"
```

PowerShell：

```powershell
$env:YMRT_APP_ID = "your_app_id"
$env:YMRT_SECRET_KEY = "your_secret_key"
$env:SMS_MOBILE = "13800138000"
```

应用读取环境变量时，缺少配置会导致供应商返回错误。当前构造函数只校验手机号，不会提前校验 App ID 或 Secret Key 是否为空。

## 5. 云片短信

### 用户需要准备

- 云片 `APIKey`。
- 已审核的模板 ID。当前仓库把项目模板 ID 固定在 `sms/consts.go`。
- 接收手机号和验证码。

### 验证码示例

```go
client, err := sms.NewVerifySms(os.Getenv("YUNPIAN_API_KEY"), mobile, code)
if err != nil {
	return err
}
if err = client.SendSmsLoginCephalon(); err != nil {
	return err
}
```

自定义模板：

```go
client, err := sms.NewNotifySms(os.Getenv("YUNPIAN_API_KEY"), mobile)
if err != nil {
	return err
}
return client.SendSmsCustom(templateID, url.Values{"#code#": {code}})
```

云片请求使用 `https://sms.yunpian.com/v2/sms/tpl_single_send.json`。

通知短信使用 `sms.NewNotifySms`，再调用 `SendSmsCephalonRenewalBalanceNotEnough`、`SendSmsCephalonMissionExpired` 或对应的时长预警方法。

## 6. 腾讯云短信

### 用户需要准备

| 配置 | 代码参数 |
| --- | --- |
| 短信应用 SDK AppID | `sdkAppId` |
| 腾讯云 SecretId | `secretId` |
| 腾讯云 SecretKey | `secretKey` |
| 已审核模板 ID | 当前固定在 `sms/consts.go` |
| 已审核签名 | 当前固定为 `深圳端脑科技` |
| 手机号 | `phones []string`，推荐 `+8613800138000` |

### 验证码示例

```go
client, err := sms.NewVerifySmsTencent(
	os.Getenv("TENCENTCLOUD_SMS_SDK_APP_ID"),
	os.Getenv("TENCENTCLOUD_SECRET_ID"),
	os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	code,
	[]string{"+8613800138000"},
)
if err != nil {
	return err
}
_, err = client.SendTencentSmsCephalonCoreLogin()
return err
```

当前 SDK 固定使用地域 `ap-guangzhou`、域名 `sms.tencentcloudapi.com` 和签名算法 `HmacSHA1`。如果接入其他腾讯云账号，先确认签名和模板属于该账号；不匹配时需要修改仓库常量后再部署。

通知短信使用 `sms.NewNotifySmsTencent`，再调用 `SendTencentSmsCephalonRenewalBalanceNotEnough`、`SendTencentSmsCephalonMissionExpired` 或对应的时长预警方法。

## 7. 滨礼网络短信

### 用户需要准备

- 滨礼网络账号 `account`。
- 管理员提供的接口密码 `password`。
- 接收手机号。

代码会在请求前把原始密码计算为 32 位小写 MD5，并以 JSON 字段 `pwd` 发送。不要提前 MD5 两次。

### 验证码示例

```go
client, err := sms.NewVerifySmsBeanli(
	os.Getenv("BEANLI_SMS_ACCOUNT"),
	os.Getenv("BEANLI_SMS_PASSWORD"),
	mobile,
	code,
)
if err != nil {
	return err
}
return client.SendSmsLoginCephalon()
```

当前请求地址为 `http://sms.beanli.com:7891/api/v1/send`，请求字段为 `account`、`pwd`、`mobiles`、`content`。部署环境必须允许访问该 HTTP 地址；生产环境应结合供应商能力评估传输安全。

通知短信使用 `sms.NewNotifySmsBeanli`，再调用 `SendSmsCustom(content)` 发送自定义文字。

## 8. 邮件

### 阿里邮箱

```go
client := mail.NewVerifyMail(
	os.Getenv("SMTP_USERNAME"),
	os.Getenv("SMTP_PASSWORD"),
	os.Getenv("MAIL_FROM"),
	receiver,
	code,
)
return client.SendMailCephalonCloudLogin()
```

`NewVerifyMail` 使用阿里邮箱 SMTP：`smtp.mxhichina.com:465`。

### 腾讯企业邮

```go
client := mail.NewVerifyMail(
	os.Getenv("SMTP_USERNAME"),
	os.Getenv("SMTP_PASSWORD"),
	os.Getenv("MAIL_FROM"),
	receiver,
	code,
)
return client.SendMailCephalonCloudLoginTencent()
```

直接发送自定义邮件：

```go
return mail.SendMailAli(
	os.Getenv("SMTP_USERNAME"),
	os.Getenv("SMTP_PASSWORD"),
	os.Getenv("MAIL_FROM"),
	receiver,
	"Your App",
	"Login code",
	content,
)
```

服务型 SMTP 使用 `NewVerifyServiceMail`、`NewNotifyServiceMail`，或对应的 Tencent/Self 构造函数。自建 SMTP 的默认地址为 `mail.zk-agent.ai:587`，腾讯企业邮默认地址为 `gz-smtp.qcloudmail.com:465`。

普通通知邮件使用 `mail.NewNotifyMail`；自建 SMTP 验证/通知邮件分别使用 `mail.NewVerifyServiceMailSelf` 和 `mail.NewNotifyServiceMailSelf`。

## 9. 配置清单

最小运行配置按供应商选择：

| 供应商 | 必填配置 |
| --- | --- |
| 亿美软通 | `YMRT_APP_ID`、`YMRT_SECRET_KEY`、手机号 |
| 云片 | `YUNPIAN_API_KEY`、手机号、已审核模板 ID |
| 腾讯云 | `TENCENTCLOUD_SMS_SDK_APP_ID`、`TENCENTCLOUD_SECRET_ID`、`TENCENTCLOUD_SECRET_KEY`、手机号 |
| 滨礼网络 | `BEANLI_SMS_ACCOUNT`、`BEANLI_SMS_PASSWORD`、手机号 |
| 阿里邮箱 | SMTP 用户名、SMTP 密码、发件人、收件人 |
| 腾讯企业邮 | SMTP 用户名、SMTP 密码、发件人、收件人 |
| 自建 SMTP | 用户名、密码、SMTP 主机/端口、收件人 |

## 10. 错误与排查

| 错误现象 | 优先检查 |
| --- | --- |
| `invalid phone num` | 手机号格式；腾讯云使用 `+86` 国家码格式 |
| 亿美返回非 `SUCCESS` | App ID、Secret Key、账户余额、签名/内容审核、出口访问 `www.btom.cn:8080` |
| 云片返回非零 `code` | APIKey、模板 ID、模板参数名、手机号和账户余额 |
| 腾讯云鉴权或模板错误 | SecretId/SecretKey、SDK AppID、签名 `深圳端脑科技`、模板 ID、地域 |
| 滨礼返回非 `00` | account、原始接口密码、手机号、HTTP 地址和账户余额 |
| 邮件 SMTP 认证失败 | 用户名、SMTP 专用密码、主机、端口和 SSL 设置 |

错误应向业务层返回并记录请求追踪 ID。不要记录 SecretKey、Secret、SMTP 密码或完整短信内容。

## 11. 离线校验与真实发送

离线编译库包：

```bash
go test ./sms ./mail ./types ./pkg/...
```

真实发送需要供应商凭据和可发送手机号/邮箱。仓库 `test` 目录中的测试包含外部调用，不属于离线测试；运行前确认测试代码、凭据和目标号码。

## 12. 接入边界

- 本库是发送 SDK，不负责验证码生成、缓存、过期、重试、频控、幂等或审计。
- 业务层应限制验证码尝试次数，并为供应商调用设置业务级超时和告警。
- 模板、签名、账户余额和供应商合规审核属于使用方责任。
- 供应商接口地址、签名、模板 ID 或协议变化时，应先检查本仓库实现再升级依赖。

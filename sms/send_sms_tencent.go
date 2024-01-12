package sms

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	tecentErrors "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
	"net/http"
)

// SendSmsTencent 发送短信 - 腾讯云
// phones: 手机号数组（使用国家地区，如：+8613888888888）
// tplParams: 模板参数
func sendSmsTencent(secretId, secretKey, sdkAppId, templateId string, phones, tplParams []string) ([]byte, error) {
	// 实例化一个认证对象，入参需要传入腾讯云账户密钥对 secretId，secretKey
	credential := common.NewCredential(
		// os.Getenv("TENCENTCLOUD_SECRET_ID"),
		// os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		secretId,
		secretKey,
	)
	// 非必要步骤: 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	// SDK 默认使用 POST 方法，如果一定要使用 GET 方法，可以在这里设置。GET方法无法处理一些较大的请求
	cpf.HttpProfile.ReqMethod = http.MethodPost
	// 指定接入地域域名
	cpf.HttpProfile.Endpoint = tencentEndpoint
	// SDK默认用TC3-HMAC-SHA256进行签名，非必要请不要修改这个字段
	cpf.SignMethod = "HmacSHA1"

	// 实例化要请求产品(以sms为例)的client对象,第二个参数是地域信息，可以直接填写字符串ap-guangzhou，支持的地域列表参考 https://cloud.tencent.com/document/api/382/52071#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	request := sms.NewSendSmsRequest()
	// 短信应用ID
	request.SmsSdkAppId = common.StringPtr(sdkAppId)
	// 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名
	request.SignName = common.StringPtr(tencentSignName)
	// 模板 ID: 必须填写已审核通过的模板 ID
	request.TemplateId = common.StringPtr(templateId)
	// 模板参数: 模板参数的个数需要与 TemplateId 对应模板的变量个数保持一致，若无模板参数，则设置为空
	request.TemplateParamSet = common.StringPtrs(tplParams)
	// 下发手机号码，采用 E.164 标准，+[国家或地区码][手机号]
	request.PhoneNumberSet = common.StringPtrs(phones)
	/* 用户的 session 内容（无需要可忽略）: 可以携带用户侧 ID 等上下文信息，server 会原样返回 */
	request.SessionContext = common.StringPtr("")
	/* 短信码号扩展号（无需要可忽略）: 默认未开通，如需开通请联系 [腾讯云短信小助手] */
	request.ExtendCode = common.StringPtr("")
	/* 国内短信无需填写该项；国际/港澳台短信已申请独立 SenderId 需要填写该字段，默认使用公共 SenderId，无需填写该字段。注：月度使用量达到指定量级可申请独立 SenderId 使用，详情请联系 [腾讯云短信小助手](https://cloud.tencent.com/document/product/382/3773#.E6.8A.80.E6.9C.AF.E4.BA.A4.E6.B5.81)。 */
	request.SenderId = common.StringPtr("")

	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.SendSms(request)
	// 处理异常
	var tencentCloudSDKError *tecentErrors.TencentCloudSDKError
	if errors.As(err, &tencentCloudSDKError) {
		return nil, errors.Errorf("client send sms failed, err: %v", err)
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}

	return json.Marshal(response.Response)
}

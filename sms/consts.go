package sms

const (
	// 默认就近地域接入域名为 sms.tencentcloudapi.com ，也支持指定地域域名访问，例如广州地域的域名为 sms.ap-guangzhou.tencentcloudapi.com
	tencentEndpoint = "sms.tencentcloudapi.com"
	tencentSdkAppId = "1400882834"
	tencentSignName = "深圳端脑科技"
)

const (
	TplIdYPForCephalonCoreRegister            = "5706256" // Cephalon Core 注册模板编号
	TplIdYPForCephalonCoreLogin               = "5706258" // Cephalon Core 登录模板编号
	TplIdYPForCephalonCoreModify              = "5706266" // Cephalon Core 修改密码模板编号
	TplIdYPForYuanHuiRegister                 = "5706320" // 元绘注册模板编号
	TplIdYPForYuanHuiLogin                    = "5755196" // 元绘登录模板编号
	TplIdYPForYuanHuiBindWechat               = "5755198" // 元绘微信绑定模板编号
	TplIdYPForCephalonCloudModify             = "5720264" // Cephalon Cloud 修改密码模板编号
	TplIdYPForCephalonCloudLogin              = "5720262" // Cephalon Cloud 登录模板编号
	TplIdYPForCephalonCloudRegister           = "5720260" // Cephalon Cloud 注册模板编号
	TplIdYPForCephalonBindWechat              = "5749864" // Cephalon Cloud 绑定微信模板编号
	TplIdYPForCephalonRenewalBalanceNotEnough = "5783976" // Cephalon Cloud 自动续费余额不足通知模板编号
	TplIdYPForCephalonMissionExpired          = "5783972" // Cephalon Cloud 任务到期提醒模板编号
	TplIdForCephalonMissionRunningTimeWarning = "5872698" // Cephalon Cloud 任务运行时长提醒模板编号
)

const (
	TplIdTencentForCephalonCoreRegister              = "2047320" // Cephalon Core 注册模板编号
	TplIdTencentForCephalonCoreLogin                 = "2047321" // Cephalon Core 登录模板编号
	TplIdTencentForCephalonCoreModify                = "2047323" // Cephalon Core 修改密码模板编号
	TplIdTencentForCephalonCloudModify               = "2047324" // Cephalon Cloud 修改密码模板编号
	TplIdTencentForCephalonCloudLogin                = "2047326" // Cephalon Cloud 登录模板编号
	TplIdTencentForCephalonCloudRegister             = "2047327" // Cephalon Cloud 注册模板编号
	TplIdTencentForCephalonBindWechat                = "2047328" // Cephalon Cloud 绑定微信模板编号
	TplIdTencentForCephalonRenewalBalanceNotEnough   = "2047330" // Cephalon Cloud 自动续费余额不足通知模板编号
	TplIdTencentForCephalonMissionExpired            = "2047331" // Cephalon Cloud 任务到期提醒模板编号
	TplIdTencentForCephalonMissionRunningTimeWarning = "2058052" // Cephalon Cloud 任务运行时长提醒模板编号
)

const (
	WarningTimeParamSixHour        = "6"
	WarningTimeParamTwelveHour     = "12"
	WarningTimeParamTwentyFourHour = "24"
)

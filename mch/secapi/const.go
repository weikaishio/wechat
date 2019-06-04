package secapi

const (
	ReceiverCustomRelation_SERVICE_PROVIDER = "SERVICE_PROVIDER" // 服务商
	ReceiverCustomRelation_STORE            = "STORE"            // 门店
	ReceiverCustomRelation_STAFF            = "STAFF"            // 员工
	ReceiverCustomRelation_STORE_OWNER      = "STORE_OWNER"      // 店主
	ReceiverCustomRelation_PARTNER          = "PARTNER"          // 合作伙伴
	ReceiverCustomRelation_HEADQUARTER      = "HEADQUARTER"      // 总部
	ReceiverCustomRelation_BRAND            = "BRAND"            // 品牌方
	ReceiverCustomRelation_DISTRIBUTOR      = "DISTRIBUTOR"      // 分销商
	ReceiverCustomRelation_USER             = "USER"             // 用户
	ReceiverCustomRelation_SUPPLIER         = "SUPPLIER"         // 供应商
	ReceiverCustomRelation_CUSTOM           = "CUSTOM"           // 自定义
)

const (
	ProfitSharingReceiverType_MERCHANT_ID         = "MERCHANT_ID"         // 商户ID
	ProfitSharingReceiverType_PERSONAL_WECHATID   = "PERSONAL_WECHATID"   // 个人微信号
	ProfitSharingReceiverType_PERSONAL_OPENID     = "PERSONAL_OPENID"     // 个人openid（由父商户APPID转换得到
	ProfitSharingReceiverType_PERSONAL_SUB_OPENID = "PERSONAL_SUB_OPENID" // 个人sub_openid（由子商户APPID转换得到）
)

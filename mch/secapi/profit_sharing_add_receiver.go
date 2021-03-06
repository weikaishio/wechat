package secapi

import (
	"errors"
	"github.com/weikaishio/wechat/mch/core"
)

/*
https://pay.weixin.qq.com/wiki/doc/api/allocation_sl.php?chapter=25_3&index=4
*/
func ProfitSharingAddReceiver(clt *core.Client, req *ProfitSharingAddReceiverRequest) (resp *ProfitSharingAddReceiverResponse, err error) {
	m1 := RequestToMap(req)

	if clt == nil {
		return nil, errors.New("core.Client is nil")
	}
	m1["sign_type"] = core.SignType_HMAC_SHA256
	m2, err := clt.PostXML(core.APIBaseURL()+"/pay/profitsharingaddreceiver", m1)
	if err != nil {
		return nil, err
	}

	resp = &ProfitSharingAddReceiverResponse{}
	ResponseFromMap(m2, resp)
	return resp, nil
}

type ProfitSharingReceiver struct {
	Type           string `json:"type,omitempty"`            // MERCHANT_ID：商户ID PERSONAL_WECHATID：个人微信号PERSONAL_OPENID：个人openid（由父商户APPID转换得到）PERSONAL_SUB_OPENID: 个人sub_openid（由子商户APPID转换得到）
	Account        string `json:"account,omitempty"`         // 类型是MERCHANT_ID时，是商户ID  类型是PERSONAL_WECHATID时，是个人微信号  类型是PERSONAL_OPENID时，是个人openid  类型是PERSONAL_SUB_OPENID时，是个人sub_openid
	Name           string `json:"name,omitempty"`            // 分账接收方类型是MERCHANT_ID时，是商户全称（必传） 分账接收方类型是PERSONAL_NAME 时，是个人姓名（必传） 分账接收方类型是PERSONAL_OPENID时，是个人姓名（选传，传则校验） 分账接收方类型是PERSONAL_SUB_OPENID时，是个人姓名（选传，传则校验）
	Amount         int    `json:"amount,omitempty"`          // 分账金额，单位为分，只能为整数，不能超过原订单支付金额及最大分账比例金额
	Description    string `json:"description,omitempty"`     // 分账的原因描述，分账账单中需要体现
	RelationType   string `json:"relation_type,omitempty"`   // 子商户与接收方的关系。 本字段值为枚举： SERVICE_PROVIDER：服务商 STORE：门店 STAFF：员工 STORE_OWNER：店主 PARTNER：合作伙伴 HEADQUARTER：总部 BRAND：品牌方 DISTRIBUTOR：分销商 USER：用户 SUPPLIER：供应商 CUSTOM：自定义
	CustomRelation string `json:"custom_relation,omitempty"` // 子商户与接收方具体的关系，本字段最多10个字。 当字段relation_type的值为CUSTOM时，本字段必填 当字段relation_type的值不为CUSTOM时，本字段无需填写
}

type ProfitSharingAddReceiverRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	MchId    string `xml:"mch_id"`     // 微信支付分配的商户号
	SubMchId string `xml:"sub_mch_id"` // 微信支付分配的子商户号
	Appid    string `xml:"appid"`      // 微信分配的公众账号ID
	SubAppid string `xml:"sub_appid"`  // 微信分配的子商户公众账号ID *非必填
	NonceStr string `xml:"nonce_str"`  // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType string `xml:"sign_type"`  // 签名类型，默认为MD5，支持HMAC-SHA256和MD5。

	Receiver string `xml:"receiver"` // 分账接收方对象，json格式
}

type ProfitSharingAddReceiverResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	ReturnCode string `xml:"return_code"` // SUCCESS/FAIL 此字段是通信标识，非交易标识
	ReturnMsg  string `xml:"return_msg"`  // 返回信息，如非空，为错误原因

	// 以下字段在return_code为SUCCESS的时候有返回
	ResultCode string `xml:"result_code"`  // SUCCESS：分账申请接收成功，结果通过分账查询接口查询 FAIL ：提交业务失败
	ErrCode    string `xml:"err_code"`     // 列表详见错误码列表
	ErrCodeDes string `xml:"err_code_des"` // 结果信息描述

	// 以下字段在return_code和result_code都为SUCCESS的时候返回
	MchId    string `xml:"mch_id"`     // 调用接口时提供的商户号
	SubMchId string `xml:"sub_mch_id"` // 微信支付分配的子商户号
	Appid    string `xml:"appid"`      // 微信分配的公众账号ID
	SubAppid string `xml:"sub_appid"`  // 微信分配的子商户公众账号ID
	Receiver string `xml:"receiver"`   // 分账接收方对象，json格式
	NonceStr string `xml:"nonce_str"`  // 微信返回的随机字符串
	Sign     string `xml:"sign"`       // 微信返回的签名
}

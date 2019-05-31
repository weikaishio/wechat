package secapi

import (
	"errors"
	"github.com/weikaishio/wechat/mch/core"
)

/*
https://pay.weixin.qq.com/wiki/doc/api/allocation_sl.php?chapter=25_6&index=2
*/
func MultiProfitSharing(clt *core.Client, req *MultiProfitSharingRequest) (resp *MultiProfitSharingResponse, err error) {
	m1 := RequestToMap(req)

	if clt == nil {
		return nil, errors.New("core.Client is nil")
	}
	m2, err := clt.PostXML(core.APIBaseURL()+"/secapi/pay/multiprofitsharing", m1)
	if err != nil {
		return nil, err
	}

	resp = &MultiProfitSharingResponse{}
	ResponseFromMap(m2, resp)
	return resp, nil
}

type MultiProfitSharingRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	MchId         string `xml:"mch_id"`         // 微信支付分配的商户号
	SubMchId      string `xml:"sub_mch_id"`     // 微信支付分配的子商户号
	Appid         string `xml:"appid"`          // 微信分配的公众账号ID
	SubAppid      string `xml:"sub_appid"`      // 微信分配的子商户公众账号ID *非必填
	NonceStr      string `xml:"nonce_str"`      // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType      string `xml:"sign_type"`      // 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
	TransactionId string `xml:"transaction_id"` // 微信支付订单号
	OutOrderNo    string `xml:"out_order_no"`   // 商户系统内部的分账单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一分账单号多次请求等同一次。
	Receivers     string `xml:"receivers"`      // 分账接收方列表，不超过50个json对象，可以设置出资子商户作为分账接受方
	Type          string `xml:"type"`           // MERCHANT_ID：商户ID PERSONAL_WECHATID：个人微信号PERSONAL_OPENID：个人openid（由父商户APPID转换得到）PERSONAL_SUB_OPENID: 个人sub_openid（由子商户APPID转换得到）
	Account       string `xml:"account"`        // 类型是MERCHANT_ID时，是商户ID 类型是PERSONAL_WECHATID时，是个人微信号 类型是PERSONAL_OPENID时，是个人openid 类型是PERSONAL_SUB_OPENID时，是个人sub_openid
	Amount        int    `xml:"amount"`         // 分账金额，单位为分，只能为整数，不能超过原订单支付金额及最大分账比例金额
	Description   string `xml:"description"`    // 分账的原因描述，分账账单中需要体现
}

type MultiProfitSharingResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	ReturnCode string `xml:"return_code"` // SUCCESS/FAIL 此字段是通信标识，非交易标识
	ReturnMsg  string `xml:"return_msg"`  // 返回信息，如非空，为错误原因

	// 以下字段在return_code为SUCCESS的时候有返回
	ResultCode string `xml:"result_code"`  // SUCCESS：分账申请接收成功，结果通过分账查询接口查询 FAIL ：提交业务失败
	ErrCode    string `xml:"err_code"`     // 列表详见错误码列表
	ErrCodeDes string `xml:"err_code_des"` // 结果信息描述

	MchId    string `xml:"mch_id"`     // 调用接口时提供的商户号
	SubMchId string `xml:"sub_mch_id"` // 微信支付分配的子商户号
	Appid    string `xml:"appid"`      // 微信分配的公众账号ID
	SubAppid string `xml:"sub_appid"`  // 微信分配的子商户公众账号ID
	NonceStr string `xml:"nonce_str"`  // 微信返回的随机字符串
	Sign     string `xml:"sign"`       // 微信返回的签名

	// 以下字段在return_code和result_code都为SUCCESS的时候返回
	TransactionId string `xml:"transaction_id"` // 微信支付订单号
	OutOrderNo    string `xml:"out_order_no"`   // 调用接口提供的商户系统内部的分账单号
	OrderId       string `xml:"order_id"`       // 微信分账单号，微信系统返回的唯一标识
}

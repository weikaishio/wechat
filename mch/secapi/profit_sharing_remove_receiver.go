package secapi

import (
	"errors"
	"github.com/weikaishio/wechat/mch/core"
)

/*
https://pay.weixin.qq.com/wiki/doc/api/allocation_sl.php?chapter=25_4&index=5
*/
func ProfitSharingRemoveReceiver(clt *core.Client, req *ProfitSharingRemoveReceiverRequest) (resp *ProfitSharingRemoveReceiverResponse, err error) {
	m1 := RequestToMap(req)

	if clt == nil {
		return nil, errors.New("core.Client is nil")
	}
	m1["sign_type"] = core.SignType_HMAC_SHA256
	m2, err := clt.PostXML(core.APIBaseURL()+"/pay/profitsharingremovereceiver", m1)
	if err != nil {
		return nil, err
	}

	resp = &ProfitSharingRemoveReceiverResponse{}
	ResponseFromMap(m2, resp)
	return resp, nil
}

type ProfitSharingRemoveReceiverRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	MchId    string `xml:"mch_id"`     // 微信支付分配的商户号
	SubMchId string `xml:"sub_mch_id"` // 微信支付分配的子商户号
	Appid    string `xml:"appid"`      // 微信分配的公众账号ID
	SubAppid string `xml:"sub_appid"`  // 微信分配的子商户公众账号ID *非必填
	NonceStr string `xml:"nonce_str"`  // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType string `xml:"sign_type"`  // 签名类型，默认为MD5，支持HMAC-SHA256和MD5。

	Receiver string `xml:"receiver"` // 分账接收方对象，json格式
}

type ProfitSharingRemoveReceiverResponse struct {
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

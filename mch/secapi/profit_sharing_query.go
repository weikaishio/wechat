package secapi

import (
	"errors"
	"github.com/weikaishio/wechat/mch/core"
)

/*
https://pay.weixin.qq.com/wiki/doc/api/allocation_sl.php?chapter=25_2&index=3
*/
func ProfitSharingQuery(clt *core.Client, req *ProfitSharingQueryRequest) (resp *ProfitSharingQueryResponse, err error) {
	m1 := RequestToMap(req)

	if clt == nil {
		return nil, errors.New("core.Client is nil")
	}
	m2, err := clt.PostXML(core.APIBaseURL()+"/pay/profitsharingquery", m1)
	if err != nil {
		return nil, err
	}

	resp = &ProfitSharingQueryResponse{}
	ResponseFromMap(m2, resp)
	return resp, nil
}

type ProfitSharingQueryRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	MchId         string `xml:"mch_id"`         // 微信支付分配的商户号
	SubMchId      string `xml:"sub_mch_id"`     // 微信支付分配的子商户号
	TransactionId string `xml:"transaction_id"` // 微信支付订单号
	OutOrderNo    string `xml:"out_order_no"`   // 商户系统内部的分账单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一分账单号多次请求等同一次。
	NonceStr      string `xml:"nonce_str"`      // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType      string `xml:"sign_type"`      // 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
}

type ProfitSharingQueryResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	ReturnCode string `xml:"return_code"` // SUCCESS/FAIL 此字段是通信标识，非交易标识
	ReturnMsg  string `xml:"return_msg"`  // 返回信息，如非空，为错误原因

	// 以下字段在return_code为SUCCESS的时候有返回
	ResultCode string `xml:"result_code"`  // SUCCESS：分账申请接收成功，结果通过分账查询接口查询 FAIL ：提交业务失败
	ErrCode    string `xml:"err_code"`     // 列表详见错误码列表
	ErrCodeDes string `xml:"err_code_des"` // 结果信息描述

	// 以下字段在return_code和result_code都为SUCCESS的时候返回
	TransactionId string `xml:"transaction_id"` // 微信支付订单号
	OutOrderNo    string `xml:"out_order_no"`   // 调用接口提供的商户系统内部的分账单号
	OrderId       string `xml:"order_id"`       // 微信分账单号，微信系统返回的唯一标识
	Status        string `xml:"status"`         // 分账单状态： ACCEPTED—受理成功 PROCESSING—处理中 FINISHED—处理完成 CLOSED—处理失败，已关单
	CloseReason   string `xml:"close_reason"`   // NO_AUTH:分账授权已解除
	Receivers     string `xml:"receivers"`      // 分账接收方列表，json对象详细说明见下文，仅当查询分账请求结果时，存在本字段
	Amount        int    `xml:"amount"`         // 分账完结的分账金额，单位为分， 仅当查询分账完结的执行结果时，存在本字段
	Description   string `xml:"description"`    // 分账完结的原因描述，仅当查询分账完结的执行结果时，存在本字段
}

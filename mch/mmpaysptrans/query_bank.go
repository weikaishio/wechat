package mmpaysptrans

import (
	"errors"
	"github.com/weikaishio/wechat/mch/core"
	"github.com/weikaishio/wechat/mch/secapi"
)

func QueryBank(clt *core.Client, req *QueryBankRequest) (resp *QueryBankResponse, err error) {
	m1 := secapi.RequestToMap(req)

	if clt == nil {
		return nil, errors.New("core.Client is nil")
	}
	m2, err := clt.PostXML(core.APIBaseURL()+"/mmpaysptrans/query_bank", m1)
	if err != nil {
		return nil, err
	}

	resp = &QueryBankResponse{}
	secapi.ResponseFromMap(m2, resp)
	return resp, nil
}

type QueryBankRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	MchId          string `xml:"mch_id"`           // 微信支付分配的商户号
	PartnerTradeNo string `xml:"partner_trade_no"` // 商户企业付款单号 商户订单号，需保持唯一（只允许数字[0~9]或字母[A~Z]和[a~z]，最短8位，最长32位）
	NonceStr       string `xml:"nonce_str"`        // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType       string `xml:"sign_type"`        // 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
}

type QueryBankResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	ReturnCode string `xml:"return_code"` // SUCCESS/FAIL 此字段是通信标识，非交易标识
	ReturnMsg  string `xml:"return_msg"`  // 返回信息，如非空，为错误原因

	// 以下字段在return_code为SUCCESS的时候有返回
	ResultCode string `xml:"result_code"`  // SUCCESS：分账申请接收成功，结果通过分账查询接口查询 FAIL ：提交业务失败
	ErrCode    string `xml:"err_code"`     // 列表详见错误码列表
	ErrCodeDes string `xml:"err_code_des"` // 结果信息描述

	MchId          string `xml:"mch_id"`           // 调用接口时提供的商户号
	PartnerTradeNo string `xml:"partner_trade_no"` // 商户企业付款单号
	PaymentNo      string `xml:"payment_no"`       // 微信企业付款单号 代付成功后，返回的内部业务单号
	BankNoMd5      string `xml:"bank_no_md5"`      // 收款用户银行卡号(MD5加密)
	TrueNameMd5    string `xml:"true_name_md5"`    // 收款人真实姓名（MD5加密）
	Amount         int    `xml:"amount"`           // 代付金额
	Status         string `xml:"status"`           // 代付订单状态： PROCESSING（处理中，如有明确失败，则返回额外失败原因；否则没有错误原因） SUCCESS（付款成功） FAILED（付款失败,需要替换付款单号重新发起付款） BANK_FAIL（银行退票，订单状态由付款成功流转至退票,退票时付款金额和手续费会自动退还）
	CmmsAmt        int    `xml:"cmms_amt"`         // 手续费金额 手续费金额 RMB：分
	CreateTime     string `xml:"create_time"`      // 微信侧订单创建时间
	PaySuccTime    string `xml:"pay_succ_time"`    // 微信侧付款成功时间（依赖银行的处理进度，可能出现延迟返回，甚至被银行退票的情况）
	Reason         string `xml:"reason"`           // 订单失败原因（如：余额不足）
}

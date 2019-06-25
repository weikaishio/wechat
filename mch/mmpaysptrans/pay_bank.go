package mmpaysptrans

import (
	"errors"
	"github.com/weikaishio/wechat/mch/core"
	"github.com/weikaishio/wechat/mch/secapi"
)

func PayBank(clt *core.Client, req *PayBankRequest) (resp *PayBankResponse, err error) {
	m1 := secapi.RequestToMap(req)

	if clt == nil {
		return nil, errors.New("core.Client is nil")
	}
	m2, err := clt.PostXML(core.APIBaseURL()+"/mmpaysptrans/pay_bank", m1)
	if err != nil {
		return nil, err
	}

	resp = &PayBankResponse{}
	secapi.ResponseFromMap(m2, resp)
	return resp, nil
}

type PayBankRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	MchId          string `xml:"mch_id"`           // 微信支付分配的商户号
	PartnerTradeNo string `xml:"partner_trade_no"` // 商户企业付款单号 商户订单号，需保持唯一（只允许数字[0~9]或字母[A~Z]和[a~z]，最短8位，最长32位）
	NonceStr       string `xml:"nonce_str"`        // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType       string `xml:"sign_type"`        // 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
	EncBankNo      string `xml:"enc_bank_no"`      // 收款方银行卡号 收款方银行卡号（采用标准RSA算法，公钥由微信侧提供）,详见获取RSA加密公钥API
	EncTrueName    string `xml:"enc_true_name"`    // 收款方用户名 收款方用户名（采用标准RSA算法，公钥由微信侧提供）详见获取RSA加密公钥API
	BankCode       string `xml:"bank_code"`        // 收款方开户行 银行卡所在开户行编号,详见银行编号列表
	Amount         int    `xml:"amount"`           // 付款金额付款金额：RMB分（支付总额，不含手续费） 注：大于0的整数
	Desc           string `xml:"desc"`             // 付款说明 企业付款到银行卡付款说明,即订单备注（UTF8编码，允许100个字符以内）
}

type PayBankResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	ReturnCode string `xml:"return_code"` // SUCCESS/FAIL 此字段是通信标识，非交易标识
	ReturnMsg  string `xml:"return_msg"`  // 返回信息，如非空，为错误原因

	// 以下字段在return_code为SUCCESS的时候有返回
	ResultCode string `xml:"result_code"`  // SUCCESS：分账申请接收成功，结果通过分账查询接口查询 FAIL ：提交业务失败
	ErrCode    string `xml:"err_code"`     // 列表详见错误码列表
	ErrCodeDes string `xml:"err_code_des"` // 结果信息描述

	MchId          string `xml:"mch_id"`           // 调用接口时提供的商户号
	PartnerTradeNo string `xml:"partner_trade_no"` // 商户企业付款单号
	Amount         int    `xml:"amount"`           // 代付金额
	PaymentNo      string `xml:"payment_no"`       // 微信企业付款单号 代付成功后，返回的内部业务单号
	CmmsAmt        int    `xml:"cmms_amt"`         // 手续费金额 手续费金额 RMB：分
}

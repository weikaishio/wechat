package mmpaysptrans

import (
	"github.com/weikaishio/wechat/mch/core"
	"github.com/weikaishio/wechat/mch/secapi"
	"errors"
)

func GetPublicKey(clt *core.Client, req *GetPublicKeyRequest) (resp *GetPublicKeyResponse, err error) {
	m1 := secapi.RequestToMap(req)

	if clt == nil {
		return nil, errors.New("core.Client is nil")
	}
	m2, err := clt.PostXML("https://fraud.mch.weixin.qq.com/risk/getpublickey", m1)
	if err != nil {
		return nil, err
	}

	resp = &GetPublicKeyResponse{}
	secapi.ResponseFromMap(m2, resp)
	return resp, nil
}

type GetPublicKeyRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	MchId    string `xml:"mch_id"`    // 微信支付分配的商户号
	NonceStr string `xml:"nonce_str"` // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType string `xml:"sign_type"` // 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
}

type GetPublicKeyResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	ReturnCode string `xml:"return_code"` // SUCCESS/FAIL 此字段是通信标识，非交易标识
	ReturnMsg  string `xml:"return_msg"`  // 返回信息，如非空，为错误原因

	// 以下字段在return_code为SUCCESS的时候有返回
	ResultCode string `xml:"result_code"`  // SUCCESS：分账申请接收成功，结果通过分账查询接口查询 FAIL ：提交业务失败
	ErrCode    string `xml:"err_code"`     // 列表详见错误码列表
	ErrCodeDes string `xml:"err_code_des"` // 结果信息描述

	MchId  string `xml:"mch_id"`  // 调用接口时提供的商户号
	PubKey string `xml:"pub_key"` //RSA 公钥
}

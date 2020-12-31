package wxmch_api

import (
	"errors"
	"fmt"
)

/*
	普通支付通知
	退款通知
	分账动账通知
 */
// 通知类型枚举
type EventTypeEnum string

// 支付成功
const EVENTTYPE_TRANSACTION_SUCCESS EventTypeEnum = "TRANSACTION.SUCCESS"
// 退款成功
const EVENTTYPE_REFUND_SUCCESS EventTypeEnum = "REFUND.SUCCESS"
// 退款异常
const EVENTTYPE_REFUND_ABNORMAL EventTypeEnum = "REFUND.ABNORMAL"
// 退款关闭
const EVENTTYPE_REFUND_CLOSED EventTypeEnum = "REFUND.CLOSED"
// 分账回退
const EVENTTYPE_TRANSACTION_RETURN EventTypeEnum = "TRANSACTION.RETURN"


// 通知报文
type Notification struct {
	// 通知ID
	ID string `json:"id"`
	// 通知创建时间
	CreateTime string `json:"create_time"`
	// 通知类型
	EventType string `json:"event_type"`
	// 通知数据类型
	ResourceType string `json:"resource_type"`
	// 通知数据(需要解密)
	Resource CipherBlockResource `json:"resource"`
	// 回调摘要
	Summary string `json:"summary"`
}

// 加密块
type CipherBlockResource struct{
	// 加密算法类型
	Algorithm string `json:"algorithm"`
	// 数据密文
	Ciphertext string `json:"ciphertext"`
	// 附加数据
	AssociatedData string `json:"associated_data"`
	// 随机串
	Nonce string `json:"nonce"`
}

func (c MerchantApiClient) GetResourcePlainText(r CipherBlockResource) (plainText []byte, err error) {
	switch r.Algorithm {
	case "AEAD_AES_256_GCM":
		plainText = decryptCiphertext(r.AssociatedData, r.Nonce, r.Ciphertext, c.apiCert)
		break
	default:
		err = errors.New(fmt.Sprintf("algorithm:%s not supported", r.Algorithm))
	}
	return
}

// 支付成功通知参数
type PayNotification struct {
	// 服务商公众号ID
	SpAppID string `json:"sp_appid"`
	// 服务商户号
	SpMchID string `json:"sp_mchid"`
	// 二级商户公众号ID
	SubAppID string `json:"sub_appid"`
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 商户订单号
	OutTradeNo string `json:"out_trade_no"`
	// 微信支付订单号
	TransactionID string `json:"transaction_id"`
	// 交易类型
	TradeType string `json:"trade_type"`
	// 交易状态
	TradeState string `json:"trade_state"`
	// 交易状态描述
	TradeStateDesc string `json:"trade_state_desc"`
	// 付款银行
	BankType string `json:"bank_type"`
	// 附加数据
	Attach string `json:"attach"`
	// 支付完成时间
	SuccessTime string `json:"success_time"`
	// 支付者
	Payer struct{
		// 用户服务标识
		SpOpenID string `json:"sp_openid"`
		// 用户子标识
		SubOpenID string `json:"sub_openid"`
	} `json:"payer"`
	// 订单金额
	Amount struct{
		// 总金额
		Total uint `json:"total"`
		// 用户支付金额
		PayerTotal uint `json:"payer_total"`
		// 货币类型
		Currency string `json:"currency"`
		// 用户支付币种
		PayerCurrency string `json:"payer_currency"`
	} `json:"amount"`
	// 场景信息
	SceneInfo struct{
		// 商户端设备号
		DeviceID string `json:"device_id"`
	} `json:"scene_info"`
	// 优惠功能
	PromotionDetail struct{
		// 券ID
		CouponID string `json:"coupon_id"`
		// 优惠名称
		Name string `json:"name"`
		// 优惠范围
		Scope string `json:"scope"`
		// 优惠类型
		Type string `json:"type"`
		// 优惠券面额
		Amount uint `json:"amount"`
		// 活动ID
		StockID string `json:"stock_id"`
		// 微信出资
		WechatpayContribute uint `json:"wechatpay_contribute"`
		// 商户出资
		MerchantContribute uint `json:"merchant_contribute"`
		// 其他出资
		OtherContribute uint `json:"other_contribute"`
		// 优惠币种
		Currency string `json:"currency"`
		// 单品列表
		GoodsDetail []struct{
			// 商品编码
			GoodsID string `json:"goods_id"`
			// 商品数量
			Quantity uint `json:"quantity"`
			// 商品单价
			UnitPrice uint `json:"unit_price"`
			// 商品优惠金额
			DiscountAmount uint `json:"discount_amount"`
			// 商品备注
			GoodsRemark string `json:"goods_remark"`
		} `json:"goods_detail"`
	} `json:"promotion_detail"`
}

// 退款通知参数
type RefundNotification struct {
	// 服务商户号
	SpMchID string `json:"sp_mchid"`
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 商户订单号
	OutTradeNo string `json:"out_trade_no"`
	// 微信支付订单号
	TransactionID string `json:"transaction_id"`
	// 商户退款单号
	OutRefundNo string `json:"out_refund_no"`
	// 微信退款单号
	RefundID string `json:"refund_id"`
	// 退款状态
	RefundStatus string `json:"refund_status"`
	// 退款成功时间
	SuccessTime string `json:"success_time"`
	// 退款入账账户
	UserReceivedAccount string `json:"user_received_account"`
	// 金额信息
	Amount struct{
		// 订单金额
		Total uint `json:"total"`
		// 退款金额
		Refund uint `json:"refund"`
		// 用户支付金额
		PayerTotal uint `json:"payer_total"`
		// 用户退款金额	
		PayerRefund uint `json:"payer_refund"`
	} `json:"amount"`
}

// 分账动账通知
type ProfitSharingNotification struct {
	// 直连商户号
	MchID string `json:"mchid"`
	// 服务商户号
	SpMchID string `json:"sp_mchid"`
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 微信订单号
	TransactionID string `json:"transaction_id"`
	// 微信分账/回退单号
	OrderID string `json:"order_id"`
	// 商户分账/回退单号
	OutOrderNo string `json:"out_order_no"`
	// 分账接收方列表	
	Receivers []struct{
		// 分账接收方类型
		Type string `json:"type"`
		// 分账接收方帐号
		Account string `json:"account"`
		// 分账动账金额
		Amount uint `json:"amount"`
		// 分账/回退描述
		Description string `json:"description"`
	} `json:"receivers"`
	// 成功时间	
	SuccessTime string `json:"success_time"`

}
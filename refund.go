package wxmch_api

import (
	"context"
	"encoding/json"
	"fmt"
)

/*
	电商收付通提交退款
	电商收付通查询退款
*/

type RefundRequest struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 电商平台APPID
	SpAppID string `json:"sp_appid"`
	// 二级商户APPID
	SubAppID string `json:"sub_appid,omitempty"`
	// 微信订单号
	TransactionID string `json:"transaction_id,omitempty"`
	// 商户订单号
	OutTradeNo string `json:"out_trade_no,omitempty"`
	// 商户退款单号
	OutRefundNo string `json:"out_refund_no"`
	// 退款原因
	Reason string `json:"reason,omitempty"`
	//订单金额
	Amount struct {
		// 退款金额
		Refund uint `json:"refund"`
		// 原订单金额
		Total uint `json:"total"`
		// 退款币种
		Currency string `json:"currency"`
	} `json:"amount"`
	// 退款结果回调url
	NotifyUrl string `json:"notify_url"`
}

type RefundResponse struct {
	// 微信退款单号
	RefundID string `json:"refund_id"`
	// 商户退款单号
	OutRefundNo string `json:"out_refund_no"`
	// 退款创建时间
	CreateTime string `json:"create_time"`
	// 金额信息
	Amount struct {
		// 退款金额
		Refund uint `json:"refund"`
		// 用户退款金额
		PayerRefund uint `json:"payer_refund"`
		// 优惠退款金额
		DiscountRefund uint `json:"discount_refund"`
		// 退款币种
		Currency string `json:"currency"`
	} `json:"amount"`
	// 优惠退款详情
	PromotionDetail []struct {
		// 券ID
		PromotionID string `json:"promotion_id"`
		// 优惠范围
		Scope string `json:"scope"`
		// 优惠类型
		Type string `json:"type"`
		// 优惠券面额
		Amount uint `json:"amount"`
		// 优惠退款金额
		RefundAmount uint `json:"refund_amount"`
	} `json:"promotion_detail"`
}

// 申请退款
func (c MerchantApiClient) RefundApply(ctx context.Context, req RefundRequest) (resp *RefundResponse, err error) {
	url := "/v3/ecommerce/refunds/apply"
	body, _ := json.Marshal(&req)
	res, err := c.doRequestAndVerifySignature(ctx, "POST", url, nil, body)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

type QueryRefundByIDRequest struct {
	// 二级商户号
	SubMchID string
	// 微信退款单号
	RefundID string
}

type QueryRefundByOutRefundNoRequest struct {
	// 二级商户号
	SubMchID string
	// 商户退款单号
	OutRefundNo string
}

type QueryRefundResponse struct {
	// 微信退款单号
	RefundID string `json:"refund_id"`
	// 商户退款单号
	OutRefundNo string `json:"out_refund_no"`
	// 微信订单号
	TransactionID string `json:"transaction_id"`
	// 商户订单号
	OutTradeNo string `json:"out_trade_no"`
	// 退款渠道
	Channel string `json:"channel"`
	// 退款入账账户
	UserReceivedAccount string `json:"user_received_account"`
	// 退款成功时间
	SuccessTime string `json:"success_time"`
	// 退款创建时间
	CreateTime string `json:"create_time"`
	// 退款状态
	Status string `json:"status"`
	// 金额信息
	Amount struct {
		// 退款金额
		Refund uint `json:"refund"`
		// 用户退款金额
		PayerRefund uint `json:"payer_refund"`
		// 优惠退款金额
		DiscountRefund uint `json:"discount_refund"`
		// 退款币种
		Currency string `json:"currency"`
	} `json:"amount"`
	// 优惠退款详情
	PromotionDetail []struct {
		// 券ID
		PromotionID string `json:"promotion_id"`
		// 优惠范围
		Scope string `json:"scope"`
		// 优惠类型
		Type string `json:"type"`
		// 优惠券面额
		Amount uint `json:"amount"`
		// 优惠退款金额
		RefundAmount uint `json:"refund_amount"`
	} `json:"promotion_detail"`
}

// 通过微信支付退款单号查询退款
func (c MerchantApiClient) QueryRefundByID(ctx context.Context, req QueryRefundByIDRequest) (resp *QueryRefundResponse, err error) {
	url := fmt.Sprintf("/v3/ecommerce/refunds/id/%s", req.RefundID)
	qm := map[string]string{"sub_mchid": req.SubMchID}
	res, err := c.doRequestAndVerifySignature(ctx, "GET", url, qm, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

// 通过商户退款单号查询退款
func (c MerchantApiClient) QueryRefundByOutRefundNo(ctx context.Context, req QueryRefundByOutRefundNoRequest) (resp *QueryRefundResponse, err error) {
	url := fmt.Sprintf("/v3/ecommerce/refunds/out-refund-no/%s", req.OutRefundNo)
	qm := map[string]string{"sub_mchid": req.SubMchID}
	res, err := c.doRequestAndVerifySignature(ctx, "GET", url, qm, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

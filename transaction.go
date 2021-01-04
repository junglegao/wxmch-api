package wxmch_api

import (
	"context"
	"encoding/json"
	"fmt"
)

/*
	电商收付通普通支付
 */

type JsApiPrepayRequest struct {
	// 服务商公众号ID
	SpAppID string `json:"sp_appid"`
	// 服务商户号
	SpMchID string `json:"sp_mchid"`
	// 二级商户公众号ID
	SubAppID string `json:"sub_appid"`
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 商品描述
	Description string `json:"description"`
	// 商户订单号
	OutTradeNo string `json:"out_trade_no"`
	// 交易结束时间
	TimeExpire string `json:"time_expire"`
	// 附加数据
	Attach string `json:"attach"`
	// 通知地址
	NotifyUrl string `json:"notify_url"`
	// 订单优惠标记
	GoodsTag string `json:"goods_tag"`
	// 结算信息
	SettleInfo struct{
		// 是否指定分账
		ProfitSharing string `json:"profit_sharing"`
		// 补差金额
		SubsidyAmount string `json:"subsidy_amount"`
	} `json:"settle_info"`
	// 订单金额
	Amount struct{
		// 总金额
		Total uint `json:"total"`
		// 货币类型
		Currency string `json:"currency"`
	} `json:"amount"`
	// 支付者
	Payer struct{
		// 用户服务标识
		SpOpenID string `json:"sp_openid"`
		// 用户子标识
		SubOpenID string `json:"sub_openid"`
	} `json:"payer"`
	// 优惠功能
	Detail struct{
		// 订单原价
		CostPrice uint `json:"cost_price"`
		// 商品小票ID
		InvoiceID string `json:"invoice_id"`
		// 单品列表
		GoodsDetail []struct{
			// 商户侧商品编码
			MerchantGoodsID string `json:"merchant_goods_id"`
			// 微信侧商品编码
			WechatpayGoodsID string `json:"wechatpay_goods_id"`
			// 商品名称
			GoodsName string `json:"goods_name"`
			// 商品数量
			Quantity uint `json:"quantity"`
			// 商品单价	
			UnitPrice uint `json:"unit_price"`
		} `json:"goods_detail"`
	} `json:"detail"`
	// 场景信息
	SceneInfo struct{
		// 用户终端IP
		PayerClientIP string `json:"payer_client_ip"`
		// 商户端设备号
		DeviceID string `json:"device_id"`
		// 商户门店信息
		StoreInfo struct{
			// 门店编号
			ID string `json:"id"`
			// 门店名称
			Name string `json:"name"`
			// 地区编码
			AreaCode string `json:"area_code"`
			// 详细地址
			Address string `json:"address"`
 		} `json:"store_info"`
	} `json:"scene_info"`
}

type PrepayPayResponse struct {
	// 预支付交易会话标识
	PrepayID string `json:"prepay_id"`
}

// JSAPI下单API
func (c MerchantApiClient) JsApiPrepay(req JsApiPrepayRequest) (resp *PrepayPayResponse, err error) {
	url := "/v3/pay/partner/transactions/jsapi"
	body, _ := json.Marshal(&req)
	res, err := c.doRequestAndVerifySignature(context.Background(), "POST", url, "", body)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

type QueryPayResultByTransactionIDRequest struct {
	// 服务商户号
	SpMchID string
	// 二级商户号
	SubMchID string
	// 微信支付订单号
	TransactionID string
}

type QueryPayResultByOutRequestNoRequest struct {
	// 服务商户号
	SpMchID string
	// 二级商户号
	SubMchID string
	// 商户订单号
	OutTradeNo string
}

type QueryPayResultResponse struct {
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

// 微信支付订单号查询交易结果
func (c MerchantApiClient) QueryPayResultByTransactionID(req QueryPayResultByTransactionIDRequest) (resp *QueryPayResultResponse, err error) {
	url := fmt.Sprintf("/v3/pay/partner/transactions/id/%s", req.TransactionID)
	query := fmt.Sprintf("sp_mchid=%s&sub_mchid=%s", req.SpMchID, req.SpMchID)
	res, err := c.doRequestAndVerifySignature(context.Background(), "GET", url, query, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

// 商户订单号查询交易结果
func (c MerchantApiClient) QueryPayResultByOutRequestNo(req QueryPayResultByOutRequestNoRequest) (resp *QueryPayResultResponse, err error) {
	url := fmt.Sprintf("/v3/pay/partner/transactions/out-trade-no/%s", req.OutTradeNo)
	query := fmt.Sprintf("sp_mchid=%s&sub_mchid=%s", req.SpMchID, req.SpMchID)
	res, err := c.doRequestAndVerifySignature(context.Background(), "GET", url, query, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

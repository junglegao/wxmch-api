package wxmch_api

import (
	"context"
	"encoding/json"
	"fmt"
)

/*
	电商收付通分账接口：
		请求分账
		查询分账结果
		分账回退
		查询分账回退结果
		查询剩余待分账金额
		分账完结
*/

// 请求分账参数
type ProfitShareApplyRequest struct {
	// 服务商appid
	SpAppID string `json:"appid"`
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 微信订单号
	TransactionID string `json:"transaction_id"`
	//商户分账单号
	OutOrderNo string `json:"out_order_no"`
	// 分账接收方列表
	Receivers []struct {
		// 分账接收方类型
		Type string `json:"type"`
		// 分账接收方帐号
		Account string `json:"receiver_account"`
		// 分账动账金额
		Amount uint `json:"amount"`
		// 分账/回退描述
		Description string `json:"description"`
		// 分账接受方姓名
		ReceiverName string `json:"receiver_name"`
	} `json:"receivers"`
	// 是否分账完成
	Finish bool `json:"finish"`
}

type ProfitShareApplyResponse struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 微信订单号
	TransactionID string `json:"transaction_id"`
	// 商户分账单号
	OutOrderNo string `json:"out_order_no"`
	// 微信分账单号
	OrderID string `json:"order_id"`
}

// 请求分账API
func (c MerchantApiClient) ProfitShareApply(ctx context.Context, req ProfitShareApplyRequest) (resp *ProfitShareApplyResponse, err error) {
	url := "/v3/ecommerce/profitsharing/orders"
	body, _ := json.Marshal(&req)
	res, err := c.doRequestAndVerifySignature(ctx, "POST", url, "", body)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

type ProfitShareQueryRequest struct {
	// 二级商户号
	SubMchID string
	// 微信订单号
	TransactionID string
	// 商户分账单号
	OutOrderNo string
}

type ProfitShareQueryResponse struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 微信订单号
	TransactionID string `json:"transaction_id"`
	// 商户分账单号
	OutOrderNo string `json:"out_order_no"`
	// 微信分账单号
	OrderID string `json:"order_id"`
	// 分账单状态
	Status string `json:"status"`
	// 分账接收方列表
	Receivers []struct {
		// 分账接收商户号
		ReceiverMchID string `json:"receiver_mchid"`
		// 分账金额
		Amount uint `json:"amount"`
		// 分账描述
		Description string `json:"description"`
		// 分账结果
		Result string `json:"result"`
		// 完成时间
		FinishTime string `json:"finish_time"`
		// 分账失败原因
		FailReason string `json:"fail_reason"`
		// 分账接收方类型
		Type string `json:"type"`
		// 分账接收方帐号
		Account string `json:"receiver_account"`
	} `json:"receivers"`
	// 关单原因
	CloseReason string `json:"close_reason"`
	// 分账完结金额
	FinishAmount uint `json:"finish_amount"`
	// 分账完结描述
	FinishDescription uint `json:"finish_description"`
}

// 查询分账结果API
func (c MerchantApiClient) ProfitShareQuery(ctx context.Context, req ProfitShareQueryRequest) (resp *ProfitShareQueryResponse, err error) {
	url := "/v3/ecommerce/profitsharing/orders"
	query := fmt.Sprintf("sub_mchid=%s&transaction_id=%s&out_order_no=%s", req.SubMchID, req.TransactionID, req.OutOrderNo)
	res, err := c.doRequestAndVerifySignature(ctx, "GET", url, query, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

// 分账回退参数
type ProfitReturnApplyRequest struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 微信分账单号
	OrderID string `json:"order_id"`
	// 商户分账单号
	OutOrderNo string `json:"out_order_no"`
	// 商户回退单号
	OutReturnNo string `json:"out_return_no"`
	// 回退商户号
	ReturnMchID string `json:"return_mchid"`
	// 回退金额
	Amount uint `json:"amount"`
	// 回退描述
	Description string `json:"description"`
}

// 分账回退返回
type ProfitReturnApplyResponse struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 微信分账单号
	OrderID string `json:"order_id"`
	// 商户分账单号
	OutOrderNo string `json:"out_order_no"`
	// 商户回退单号
	OutReturnNo string `json:"out_return_no"`
	// 回退商户号
	ReturnMchID string `json:"return_mchid"`
	// 回退金额
	Amount uint `json:"amount"`
	// 微信回退单号
	ReturnNo string `json:"return_no"`
	// 回退结果
	Result string `json:"result"`
	// 失败原因
	FailReason string `json:"fail_reason"`
	// 分账回退完成时间
	FinishTime string `json:"finish_time"`
}

// 请求分账回退API
func (c MerchantApiClient) ProfitReturnApply(ctx context.Context, req ProfitReturnApplyRequest) (resp *ProfitReturnApplyResponse, err error) {
	url := "/v3/ecommerce/profitsharing/returnorders"
	body, _ := json.Marshal(&req)
	res, err := c.doRequestAndVerifySignature(ctx, "POST", url, "", body)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

// 查询分账回退结果参数
type ProfitReturnQueryRequest struct {
	// 二级商户号
	SubMchID string
	// 微信分账单号
	OrderID string
	// 商户分账单号
	OutOrderNo string
	// 商户回退单号
	OutReturnNo string
}

type ProfitReturnQueryResponse struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 微信分账单号
	OrderID string `json:"order_id"`
	// 商户分账单号
	OutOrderNo string `json:"out_order_no"`
	// 商户回退单号
	OutReturnNo string `json:"out_return_no"`
	// 回退商户号
	ReturnMchID string `json:"return_mchid"`
	// 回退金额
	Amount uint `json:"amount"`
	// 微信回退单号
	ReturnNo string `json:"return_no"`
	// 回退结果
	Result string `json:"result"`
	// 失败原因
	FailReason string `json:"fail_reason"`
	// 分账回退完成时间
	FinishTime string `json:"finish_time"`
}

// 查询分账回退结果API
func (c MerchantApiClient) ProfitReturnQuery(ctx context.Context, req ProfitReturnQueryRequest) (resp ProfitReturnQueryResponse, err error) {
	url := "/v3/ecommerce/profitsharing/returnorders"
	query := fmt.Sprintf("sub_mchid=%s&out_return_no=%s", req.SubMchID, req.OutReturnNo)
	if req.OutOrderNo == "" && req.OrderID == "" {
		err = fmt.Errorf("out_order_no和order_id不能同时为空")
		return
	}
	if req.OutOrderNo != "" {
		query += fmt.Sprintf("&out_order_no=%s", req.OutOrderNo)
	}
	if req.OrderID != "" {
		query += fmt.Sprintf("&order_id=%s", req.OrderID)
	}
	res, err := c.doRequestAndVerifySignature(ctx, "GET", url, query, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

// 查询订单剩余待分账金额参数
type ProfitShareUnSplitAmountQueryRequest struct {
	// 微信订单号
	TransactionID string `json:"transaction_id"`
}

// 查询订单剩余待分账金额返回
type ProfitShareUnSplitAmountQueryResponse struct {
	// 微信订单号
	TransactionID string `json:"transaction_id"`
	// 订单剩余待分金额
	UnSplitAmount uint `json:"unsplit_amount"`
}

// 查询订单剩余待分金额API
func (c MerchantApiClient) ProfitShareUnSplitAmountQuery(ctx context.Context, req ProfitShareUnSplitAmountQueryRequest) (resp *ProfitShareUnSplitAmountQueryResponse, err error) {
	url := fmt.Sprintf("/v3/ecommerce/profitsharing/orders/%s/amounts", req.TransactionID)
	res, err := c.doRequestAndVerifySignature(ctx, "GET", url, "query", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

// 完结分账参数
type ProfitShareFinishRequest struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 微信订单号
	TransactionID string `json:"transaction_id"`
	// 商户分账单号
	OutOrderNo string `json:"out_order_no"`
	// 分账描述
	Description string `json:"description"`
}

// 完结分账返回
type ProfitShareFinishResponse struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 微信订单号
	TransactionID string `json:"transaction_id"`
	// 商户分账单号
	OutOrderNo string `json:"out_order_no"`
	// 微信分账单号
	OrderID string `json:"order_id"`
}

// 完结分账API
func (c MerchantApiClient) ProfitShareFinish(ctx context.Context, req ProfitShareFinishRequest) (resp *ProfitShareFinishResponse, err error) {
	url := "/v3/ecommerce/profitsharing/finish-order"
	body, _ := json.Marshal(&req)
	res, err := c.doRequestAndVerifySignature(ctx, "POST", url, "query", body)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

package wxmch_api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// 二级商户账户实时余额查询
type SubMchBalanceQueryRequest struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 账户类型 BASIC：基本账户 OPERATION：运营账户 FEES：手续费账户
	AccountType string `json:"account_type"`
}

type SubMchBalanceQueryResponse struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 账户类型
	AccountType string `json:"account_type"`
	// 可用余额
	AvailableAmount int64 `json:"available_amount"`
	// 不可用余额
	PendingAmount int64 `json:"pending_amount"`
}

// 二级商户账户实时余额查询
func (c MerchantApiClient) SubMchBalanceQuery(ctx context.Context, req SubMchBalanceQueryRequest) (resp *SubMchBalanceQueryResponse, err error) {
	rUrl := fmt.Sprintf("/v3/ecommerce/fund/balance/%s", req.SubMchID)
	query := fmt.Sprintf("account_type=%s", url.QueryEscape(req.AccountType))
	res, err := c.doRequestAndVerifySignature(ctx, "GET", rUrl, query, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

type SubMchEndDayBalanceQueryRequest struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 日期 示例值：2019-08-17
	Date string `json:"date"`
}

type SubMchEndDayBalanceQueryResponse struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 可用余额
	AvailableAmount int64 `json:"available_amount"`
	// 不可用余额
	PendingAmount int64 `json:"pending_amount"`
}

// 二级商户账户日终余额
func (c MerchantApiClient) SubMchEndDayBalanceQuery(ctx context.Context, req SubMchEndDayBalanceQueryRequest) (resp *SubMchEndDayBalanceQueryResponse, err error) {
	rUrl := fmt.Sprintf("/v3/ecommerce/fund/enddaybalance/%s", req.SubMchID)
	query := fmt.Sprintf("date=%s", url.QueryEscape(req.Date))
	res, err := c.doRequestAndVerifySignature(ctx, "GET", rUrl, query, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

// 电商平台账户实时余额查询请求
type PlatformBalanceQueryRequest struct {
	// 账户类型 BASIC：基本账户 OPERATION：运营账户 FEES：手续费账户
	AccountType string `json:"account_type"`
}

// 电商平台账户实时余额查询返回
type PlatformBalanceQueryResponse struct {
	// 可用余额
	AvailableAmount int64 `json:"available_amount"`
	// 不可用余额
	PendingAmount int64 `json:"pending_amount"`
}

// 电商平台账户实时余额查询
func (c MerchantApiClient) PlatformBalanceQuery(ctx context.Context, req PlatformBalanceQueryRequest) (resp *PlatformBalanceQueryResponse, err error) {
	rUrl := fmt.Sprintf("/v3/merchant/fund/balance/%s", req.AccountType)
	res, err := c.doRequestAndVerifySignature(ctx, "GET", rUrl, "", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

// 电商平台账户日终余额查询请求
type PlatformEndDayBalanceQueryRequest struct {
	// 账户类型 BASIC：基本账户 OPERATION：运营账户 FEES：手续费账户
	AccountType string `json:"account_type"`
	// 日期 示例值：2019-08-17
	Date string `json:"date"`
}

// 电商平台账户日终余额查询返回
type PlatformEndDayBalanceQueryResponse struct {
	// 可用余额
	AvailableAmount int64 `json:"available_amount"`
	// 不可用余额
	PendingAmount int64 `json:"pending_amount"`
}

// 电商平台账户日终余额查询
func (c MerchantApiClient) PlatformEndDayBalanceQuery(ctx context.Context, req PlatformEndDayBalanceQueryRequest) (resp *PlatformEndDayBalanceQueryResponse, err error) {
	rUrl := fmt.Sprintf("/v3/merchant/fund/dayendbalance/%s", req.AccountType)
	query := fmt.Sprintf("date=%s", url.QueryEscape(req.Date))
	res, err := c.doRequestAndVerifySignature(ctx, "GET", rUrl, query, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

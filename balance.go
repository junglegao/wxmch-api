package wxmch_api

import (
	"context"
	"encoding/json"
	"fmt"
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
	qm := map[string]string {"account_type": req.AccountType}
	res, err := c.doRequestAndVerifySignature(ctx, "GET", rUrl, qm, nil)
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
	qm := map[string]string{"date": req.Date}
	res, err := c.doRequestAndVerifySignature(ctx, "GET", rUrl, qm, nil)
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
	res, err := c.doRequestAndVerifySignature(ctx, "GET", rUrl, nil, nil)
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
	qm := map[string]string{"date": req.Date}
	res, err := c.doRequestAndVerifySignature(ctx, "GET", rUrl, qm, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

// 二级商户提现请求
type SubMchWithdrawRequest struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 商户提现单号
	OutRequestNo string `json:"out_request_no"`
	// 提现金额
	Amount int64 `json:"amount"`
	// 提现备注
	Remark string `json:"remark,omitempty"`
	// 银行附言
	BankMemo string `json:"bank_memo,omitempty"`
}

// 二级商户提现返回
type SubMchWithdrawResponse struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 商户提现单号
	OutRequestNo string `json:"out_request_no"`
	// 微信提现单号
	WithdrawID string `json:"withdraw_id"`
}

// 二级商户提现
func (c MerchantApiClient) SubMchWithdraw(ctx context.Context, req SubMchWithdrawRequest) (resp *SubMchWithdrawResponse, err error) {
	rUrl := "/v3/ecommerce/fund/withdraw"
	body, _ := json.Marshal(&req)
	res, err := c.doRequestAndVerifySignature(ctx, "POST", rUrl, nil, body)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

// 根据微信提现单号查询二级商户提现状态
type SubMchWithdrawQueryByWithdrawIDRequest struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 微信提现单号
	WithdrawID string `json:"withdraw_id"`
}

// 根据商户提现单号查询二级商户提现状态
type SubMchWithdrawQueryByOutRequestNoRequest struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 商户提现单号
	OutRequestNo string `json:"out_request_no"`
}

// 二级商户查询提现状态返回
type SubMchWithdrawQueryResponse struct {
	// 二级商户号
	SubMchID string `json:"sub_mchid"`
	// 电商平台商户号
	SpMchID string `json:"sp_mchid"`
	// 提现单状态
	Status string `json:"status"`
	// 微信提现单号
	WithdrawID string `json:"withdraw_id"`
	// 商户提现单号
	OutRequestNo string `json:"out_request_no"`
	// 提现金额
	Amount int64 `json:"amount"`
	// 发起提现时间
	CreateTime string `json:"create_time"`
	// 提现状态更新时间
	UpdateTime string `json:"update_time"`
	// 失败原因
	Reason string `json:"reason"`
	// 提现备注
	Remark string `json:"remark"`
	// 银行附言
	BankMemo string `json:"bank_memo"`
}

// 根据微信提现单号查询二级商户提现状态
func (c MerchantApiClient) SubMchWithdrawQueryByWithdrawID(ctx context.Context, req SubMchWithdrawQueryByWithdrawIDRequest) (resp *SubMchWithdrawQueryResponse, err error) {
	rUrl := fmt.Sprintf("/v3/ecommerce/fund/withdraw/%s", req.WithdrawID)
	qm := map[string]string{"sub_mchid": req.SubMchID}
	res, err := c.doRequestAndVerifySignature(ctx, "GET", rUrl, qm, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

// 根据商户提现单号查询二级商户提现状态
func (c MerchantApiClient) SubMchWithdrawQueryByOutRequestNo(ctx context.Context, req SubMchWithdrawQueryByOutRequestNoRequest) (resp *SubMchWithdrawQueryResponse, err error) {
	rUrl := fmt.Sprintf("/v3/ecommerce/fund/withdraw/out-request-no/%s", req.OutRequestNo)
	qm := map[string]string{"sub_mchid": req.SubMchID}
	res, err := c.doRequestAndVerifySignature(ctx, "GET", rUrl, qm, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}
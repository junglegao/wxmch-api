package wxmch_api

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

/*
	(直连)批量转账到零钱
*/

type BatchTransferRequest struct {
	// 直连商户的AppID
	AppID string `json:"appid"`
	// 商家批次单号
	OutBatchNo string `json:"out_batch_no"`
	// 批次名称
	BatchName string `json:"batch_name"`
	// 批次备注
	BatchRemark string `json:"batch_remark"`
	// 转账总金额
	TotalAmount uint `json:"total_amount"`
	// 转账总笔数
	TotalNum           uint             `json:"total_num"`
	TransferDetailList []TransferDetail `json:"transfer_detail_list"`
}

type TransferDetail struct {
	// 商家明细单号
	OutDetailNo string `json:"out_detail_no"`
	// 转账金额
	TransferAmount uint `json:"transfer_amount"`
	// 转账备注
	TransferRemark string `json:"transfer_remark"`
	// OpenID
	OpenID string `json:"openid"`
	// 收款用户姓名
	UserName string `json:"user_name"`
	// 收款用户身份证
	UserIDCard string `json:"user_id_card,omitempty"`
}

type BatchTransferResponse struct {
	// 商家批次单号
	OutBatchNo string `json:"out_batch_no"`
	// 微信批次单号
	BatchID string `json:"batch_id"`
	// 批次创建时间
	CreateTime string `json:"create_time"`
}

func (c MerchantApiClient) BatchTransfer(ctx context.Context, req BatchTransferRequest) (resp *BatchTransferResponse, err error) {
	url := "/v3/transfer/batches"
	pubKey := c.getPlatformPublicKey()
	for i := range req.TransferDetailList {
		req.TransferDetailList[i].UserName = encryptCiphertext(req.TransferDetailList[i].UserName, pubKey)
		if req.TransferDetailList[i].UserIDCard != "" {
			req.TransferDetailList[i].UserIDCard = encryptCiphertext(req.TransferDetailList[i].UserIDCard, pubKey)
		}
	}
	body, _ := json.Marshal(&req)

	res, err := c.doRequestAndVerifySignature(ctx, "POST", url, nil, body)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

type BatchTransferQueryByOutNoRequest struct {
	OutBatchNo      string `json:"out_batch_no"`
	NeedQueryDetail bool   `json:"need_query_detail"`
	Offset          int64  `json:"offset,omitempty"`
	Limit           int64  `json:"limit,omitempty"`
	DetailStatus    string `json:"detail_status,omitempty"`
}

type TransferBatch struct {
	// 直连商户号
	MchID string `json:"mch_id"`
	// 商家批次单号
	OutBatchNo string `json:"out_batch_no"`
	// 微信批次单号
	BatchID string `json:"batch_id"`
	// 直连商户AppID
	AppID string `json:"app_id"`
	// 批次状态
	BatchStatus string `json:"batch_status"`
	// 批次类型
	BatchType string `json:"batch_type"`
	// 批次名称
	BatchName string `json:"batch_name"`
	// 批次备注
	BatchRemark string `json:"batch_remark"`
	// 批次关闭原因
	CloseReason string `json:"close_reason"`
	// 转账总金额
	TotalAmount int64 `json:"total_amount"`
	// 转账总笔数
	TotalNum int64 `json:"total_num"`
	// 批次创建时间
	CreateTime string `json:"create_time"`
	// 批次更新时间
	UpdateTime string `json:"update_time"`
	// 成功金额
	SuccessAmount int64 `json:"success_amount"`
	// 成功笔数
	SuccessNum int64 `json:"success_num"`
	// 失败金额
	FailAmount int64 `json:"fail_amount"`
	// 失败笔数
	FailNum int64 `json:"fail_num"`
}

type TransferDetailItem struct {
	DetailID    string `json:"detail_id"`
	OutDetailNo string `json:"out_detail_no"`
	Status      string `json:"detail_status"`
}

type BatchTransferQueryByOutNoResponse struct {
	TransferBatch   TransferBatch        `json:"transfer_batch"`
	TransferDetails []TransferDetailItem `json:"transfer_detail_list"`
	Offset          int64                `json:"offset"`
	Limit           int64                `json:"limit"`
}

func (c MerchantApiClient) BatchTransferQueryByOutNo(ctx context.Context, req BatchTransferQueryByOutNoRequest) (resp *BatchTransferQueryByOutNoResponse, err error) {
	url := fmt.Sprintf("/v3/transfer/batches/out-batch-no/%s", req.OutBatchNo)
	qm := map[string]string{
		"need_query_detail": strconv.FormatBool(req.NeedQueryDetail),
		"offset":            strconv.FormatInt(req.Offset, 10),
		"limit":             strconv.FormatInt(req.Limit, 10),
		"detail_status":     req.DetailStatus,
	}
	res, err := c.doRequestAndVerifySignature(ctx, "GET", url, qm, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

package wxmch_api

import (
	"context"
	"encoding/json"
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

	res, err := c.doRequestAndVerifySignature(ctx, "POST", url, "", body)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

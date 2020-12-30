package wxmch_api

import (
	"context"
	"encoding/json"
)

type SubmitApplymentResp struct {
	//微信支付申请单号
	ApplymentID uint `json:"applyment_id"`
	//业务申请编号
	OutRequestNo string `json:"out_request_no"`
}

type SubmitApplymentRequest struct {
	// 业务申请编号
	OutRequestNo string
	// 主体类型
	OrganizationType string
	// 营业执照/登记证书信息
	BusinessLicenseInfo struct{
		// 证件扫描件 - MediaID
		BusinessLicenseCopy string
		// 证件注册号
		BusinessLicenseNumber string
		// 商户名称
		MerchantName string
		// 经营者/法定代表人姓名
		LegalPerson string
		// 注册地址
		CompanyAddress string
		// 营业期限
		BusinessTime string
	}
}

// 二级商户进件
func (c MerchantApiClient) SubmitApplyment(req SubmitApplymentRequest) (resp *SubmitApplymentResp, err error) {

	res, err := c.doRequest(context.Background(), "GET", "/v3/ecommerce/applyments/", "", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(res), &resp)
	return

}

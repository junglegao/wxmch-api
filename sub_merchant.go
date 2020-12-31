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
	OutRequestNo string `json:"out_request_no"`
	// 主体类型
	OrganizationType string `json:"organization_type"`
	// 营业执照/登记证书信息
	BusinessLicenseInfo struct{
		// 证件扫描件 - MediaID
		BusinessLicenseCopy string `json:"business_license_copy"`
		// 证件注册号
		BusinessLicenseNumber string `json:"business_license_number"`
		// 商户名称
		MerchantName string `json:"merchant_name"`
		// 经营者/法定代表人姓名
		LegalPerson string `json:"legal_person"`
		// 注册地址
		CompanyAddress string `json:"company_address"`
		// 营业期限
		BusinessTime string `json:"business_time"`
	} `json:"business_license_info"`
	// 组织机构代码证信息
	OrganizationCertInfo struct{
		// 组织机构代码证照片
		OrganizationCopy string `json:"organization_copy"`
		// 组织机构代码
		OrganizationNumber string `json:"organization_number"`
		// 组织机构代码有效期限
		OrganizationTime string `json:"organization_time"`
	} `json:"organization_cert_info"`
	// 经营者/法人证件类型
	IDDocType string `json:"id_doc_type"`
	// 经营者/法人身份证信息
	IDCardInfo struct{
		// 身份证人像面照片
		IDCardCopy string `json:"id_card_copy"`
		// 身份证国徽面照片
		IDCardNational string `json:"id_card_national"`
		// 身份证姓名
		IDCardName string `json:"id_card_name"`
		// 身份证号码
		IDCardNumber string `json:"id_card_number"`
		// 身份证有效期限
		IDCardValidTime string `json:"id_card_valid_time"`
	} `json:"id_card_info"`
	// 经营者/法人其他类型证件信息
	IDDocInfo struct{
		// 证件姓名
		IDDocName string `json:"id_doc_name"`
		// 证件号码
		IDDocNumber string `json:"id_doc_number"`
		// 证件照片
		IDDocCopy string `json:"id_doc_copy"`
		// 证件结束日期
		DocPeriodEnd string `json:"doc_period_end"`
	} `json:"id_doc_info"`
	// 是否填写结算银行账户
	NeedAccountInfo bool `json:"need_account_info"`
	// 结算银行账户
	AccountInfo struct{
		// 账户类型
		BankAccountType string `json:"bank_account_type"`
		// 开户银行
		AccountBank string `json:"account_bank"`
		// 开户名称
		AccountName string `json:"account_name"`
		// 开户银行省市编码
		BankAddressCode string `json:"bank_address_code"`
		// 开户银行联行号
		BankBranchID string `json:"bank_branch_id"`
		// 开户银行全称 （含支行）
		BankName string `json:"bank_name"`
		// 银行帐号
		AccountNumber string `json:"account_number"`
	} `json:"account_info"`
	// 超级管理员信息
	ContactInfo struct{
		// 超级管理员类型
		ContactType string `json:"contact_type"`
		// 超级管理员姓名
		ContactName string `json:"contact_name"`
		// 超级管理员身份证件号码
		ContactIDCardNumber string `json:"contact_id_card_number"`
		// 超级管理员手机
		MobilePhone string `json:"mobile_phone"`
		// 超级管理员邮箱
		ContactEmail string `json:"contact_email"`
	} `json:"contact_info"`
	// 店铺信息
	SalesSceneInfo struct{
		// 店铺名称
		StoreName string `json:"store_name"`
		// 店铺链接
		StoreUrl string `json:"store_url"`
		// 店铺二维码
		StoreQrCode string `json:"store_qr_code"`
		// 小程序AppID
		MiniProgramSubAppID string `json:"mini_program_sub_appid"`
	} `json:"sales_scene_info"`
	// 商户简称
	MerchantShortname string `json:"merchant_shortname"`
	// 特殊资质
	Qualifications string `json:"qualifications"`
	// 补充材料
	BusinessAdditionPics string `json:"business_addition_pics"`
	// 补充说明
	BusinessAdditionDesc string `json:"business_addition_desc"`
}

// 二级商户进件
func (c MerchantApiClient) SubmitApplyment(req SubmitApplymentRequest) (resp *SubmitApplymentResp, err error) {
	body, _ := json.Marshal(&req)
	res, err := c.doRequest(context.Background(), "POST", "/v3/ecommerce/applyments/", "", body)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(res), &resp)
	return

}

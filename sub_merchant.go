package wxmch_api

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"fmt"
)

type SubmitApplymentResp struct {
	//微信支付申请单号
	ApplymentID uint `json:"applyment_id"`
	//业务申请编号
	OutRequestNo string `json:"out_request_no"`
}

type BusinessLicenseInfo struct {
	// 证件扫描件 - MediaID
	BusinessLicenseCopy string `json:"business_license_copy"`
	// 证件注册号
	BusinessLicenseNumber string `json:"business_license_number"`
	// 商户名称
	MerchantName string `json:"merchant_name"`
	// 经营者/法定代表人姓名
	LegalPerson string `json:"legal_person"`
	// 注册地址
	CompanyAddress string `json:"company_address,omitempty"`
	// 营业期限
	BusinessTime string `json:"business_time,omitempty"`
}

type OrganizationCertInfo struct {
	// 组织机构代码证照片
	OrganizationCopy string `json:"organization_copy"`
	// 组织机构代码
	OrganizationNumber string `json:"organization_number"`
	// 组织机构代码有效期限
	OrganizationTime string `json:"organization_time"`
}

type IDCardInfo struct {
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
}

type IDDocInfo struct {
	// 证件姓名
	IDDocName string `json:"id_doc_name"`
	// 证件号码
	IDDocNumber string `json:"id_doc_number"`
	// 证件照片
	IDDocCopy string `json:"id_doc_copy"`
	// 证件结束日期
	DocPeriodEnd string `json:"doc_period_end"`
}

type AccountInfo struct {
	// 账户类型
	BankAccountType string `json:"bank_account_type"`
	// 开户银行
	AccountBank string `json:"account_bank"`
	// 开户名称
	AccountName string `json:"account_name"`
	// 开户银行省市编码
	BankAddressCode string `json:"bank_address_code"`
	// 开户银行联行号
	BankBranchID string `json:"bank_branch_id,omitempty"`
	// 开户银行全称 （含支行）
	BankName string `json:"bank_name,omitempty"`
	// 银行帐号
	AccountNumber string `json:"account_number"`
}

type ContactInfo struct {
	// 超级管理员类型
	ContactType string `json:"contact_type"`
	// 超级管理员姓名
	ContactName string `json:"contact_name"`
	// 超级管理员身份证件号码
	ContactIDCardNumber string `json:"contact_id_card_number"`
	// 超级管理员手机
	MobilePhone string `json:"mobile_phone"`
	// 超级管理员邮箱
	ContactEmail string `json:"contact_email,omitempty"`
}

type SalesSceneInfo struct {
	// 店铺名称
	StoreName string `json:"store_name"`
	// 店铺链接
	StoreUrl string `json:"store_url,omitempty"`
	// 店铺二维码
	StoreQrCode string `json:"store_qr_code,omitempty"`
	// 小程序AppID
	MiniProgramSubAppID string `json:"mini_program_sub_appid,omitempty"`
}

type SubmitApplymentRequest struct {
	// 业务申请编号
	OutRequestNo string `json:"out_request_no"`
	// 主体类型
	OrganizationType string `json:"organization_type"`
	// 营业执照/登记证书信息
	BusinessLicenseInfo *BusinessLicenseInfo `json:"business_license_info,omitempty"`
	// 组织机构代码证信息
	OrganizationCertInfo *OrganizationCertInfo `json:"organization_cert_info,omitempty"`
	// 经营者/法人证件类型
	IDDocType string `json:"id_doc_type,omitempty"`
	// 经营者/法人身份证信息
	IDCardInfo *IDCardInfo `json:"id_card_info,omitempty"`
	// 经营者/法人其他类型证件信息
	IDDocInfo *IDDocInfo `json:"id_doc_info,omitempty"`
	// 是否填写结算银行账户
	NeedAccountInfo bool `json:"need_account_info"`
	// 结算银行账户
	AccountInfo *AccountInfo `json:"account_info,omitempty"`
	// 超级管理员信息
	ContactInfo ContactInfo `json:"contact_info"`
	// 店铺信息
	SalesSceneInfo SalesSceneInfo `json:"sales_scene_info"`
	// 商户简称
	MerchantShortname string `json:"merchant_shortname"`
	// 特殊资质
	Qualifications []string `json:"qualifications,omitempty"`
	// 补充材料
	BusinessAdditionPics []string `json:"business_addition_pics,omitempty"`
	// 补充说明
	BusinessAdditionDesc string `json:"business_addition_desc,omitempty"`
}

// 二级商户进件
func (c MerchantApiClient) ApplymentSubmit(ctx context.Context, req SubmitApplymentRequest) (resp *SubmitApplymentResp, err error) {
	// 法人身份证姓名和号码需要加密
	pubKey := c.getPlatformPublicKey()
	req.IDCardInfo.IDCardName = encryptCiphertext(req.IDCardInfo.IDCardName, pubKey)
	req.IDCardInfo.IDCardNumber = encryptCiphertext(req.IDCardInfo.IDCardNumber, pubKey)
	// 超级管理员姓名、身份证、手机号、邮箱需要加密
	req.ContactInfo.ContactName = encryptCiphertext(req.ContactInfo.ContactName, pubKey)
	req.ContactInfo.ContactIDCardNumber = encryptCiphertext(req.ContactInfo.ContactIDCardNumber, pubKey)
	req.ContactInfo.MobilePhone = encryptCiphertext(req.ContactInfo.MobilePhone, pubKey)
	if req.ContactInfo.ContactEmail != "" {
		req.ContactInfo.ContactEmail = encryptCiphertext(req.ContactInfo.ContactEmail, pubKey)
	}
	// 法人其他证件信息（选传）如果有，需要加密
	if req.IDDocInfo != nil && req.IDDocInfo.IDDocName != "" {
		req.IDDocInfo.IDDocName = encryptCiphertext(req.IDDocInfo.IDDocName, pubKey)
	}
	if req.IDDocInfo != nil && req.IDDocInfo.IDDocNumber != "" {
		req.IDDocInfo.IDDocNumber = encryptCiphertext(req.IDDocInfo.IDDocNumber, pubKey)
	}
	// 结算银行账户 如果有，需要加密
	if req.NeedAccountInfo {
		req.AccountInfo.AccountName = encryptCiphertext(req.AccountInfo.AccountName, pubKey)
		req.AccountInfo.AccountNumber = encryptCiphertext(req.AccountInfo.AccountNumber, pubKey)
	}

	body, _ := json.Marshal(&req)
	res, err := c.doRequestAndVerifySignature(ctx, "POST", "/v3/ecommerce/applyments/", nil, body)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return

}

type QueryApplymentByIDRequest struct {
	// 微信支付申请单号
	ApplymentID uint
}

type QueryApplymentByOutRequestNoRequest struct {
	// 业务申请编号
	OutRequestNo string
}

type ApplymentQueryResponse struct {
	// 申请状态
	ApplymentState string `json:"applyment_state"`
	// 申请状态描述
	ApplymentStateDesc string `json:"applyment_state_desc"`
	// 签约链接
	SignUrl string `json:"sign_url"`
	// 电商平台二级商户号
	SubMchID string `json:"sub_mchid"`
	// 汇款账户验证信息
	AccountValidation struct {
		// 付款户名
		AccountName string `json:"account_name"`
		// 付款卡号
		AccountNo string `json:"account_no"`
		// 汇款金额 （以分为单位）
		PayAmount string `json:"pay_amount"`
		// 收款卡号
		DestinationAccountNumber string `json:"destination_account_number"`
		// 收款户名
		DestinationAccountName string `json:"destination_account_name"`
		// 开户银行
		DestinationAccountBank string `json:"destination_account_bank"`
		// 省市信息
		City string `json:"city"`
		// 备注信息
		Remark string `json:"remark"`
		// 汇款截止时间
		Deadline string `json:"deadline"`
	} `json:"account_validation"`
	// 驳回原因详情
	AuditDetail []struct {
		// 参数名称
		ParamName string `json:"param_name"`
		// 驳回原因
		RejectReason string `json:"reject_reason"`
	} `json:"audit_detail"`
	// 法人验证链接
	LegalValidationUrl string `json:"legal_validation_url"`
	// 业务申请编号
	OutRequestNo string `json:"out_request_no"`
	// 微信支付申请单号
	ApplymentID uint `json:"applyment_id"`
}

const AccountNeedVerifyState = `ACCOUNT_NEED_VERIFY`

// 申请单查询结果脱敏
func (r *ApplymentQueryResponse) desensitize(priKey *rsa.PrivateKey) (err error) {
	// -汇款账户验证信息 当申请状态为ACCOUNT_NEED_VERIFY 时有返回，可根据指引汇款，完成账户验证。
	// 付款户名和付款卡号需要脱敏
	if r.ApplymentState == AccountNeedVerifyState {
		r.AccountValidation.AccountName, err = decryptCiphertext(r.AccountValidation.AccountName, priKey)
		if err != nil {
			return
		}
		r.AccountValidation.AccountNo, err = decryptCiphertext(r.AccountValidation.AccountNo, priKey)
		if err != nil {
			return
		}
	}
	return
}

// 通过申请单ID查询申请状态
func (c MerchantApiClient) ApplymentQueryByID(ctx context.Context, req QueryApplymentByIDRequest) (resp *ApplymentQueryResponse, err error) {
	url := fmt.Sprintf("/v3/ecommerce/applyments/%d", req.ApplymentID)
	res, err := c.doRequestAndVerifySignature(ctx, "GET", url, nil, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return
	}
	err = resp.desensitize(c.apiPriKey)
	if err != nil {
		return
	}
	return
}

// 通过业务申请编号查询申请状态
func (c MerchantApiClient) ApplymentQueryByOutRequestNo(ctx context.Context, req QueryApplymentByOutRequestNoRequest) (resp *ApplymentQueryResponse, err error) {
	url := fmt.Sprintf("/v3/ecommerce/applyments/out-request-no/%s", req.OutRequestNo)
	res, err := c.doRequestAndVerifySignature(ctx, "GET", url, nil, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return
	}
	err = resp.desensitize(c.apiPriKey)
	if err != nil {
		return
	}
	return
}

type ModifySettlementRequest struct {
	// 特约商户号
	SubMchID string
	// 账户类型
	AccountType string `json:"account_type"`
	// 开户银行
	AccountBank string `json:"account_bank"`
	// 开户银行省市编码
	BankAddressCode string `json:"bank_address_code"`
	// 开户银行全称（含支行）
	BankName string `json:"bank_name"`
	// 开户银行联行号
	BankBranchID string `json:"bank_branch_id"`
	// 银行账号
	AccountNumber string `json:"account_number"`
}

// 修改结算帐号API
func (c MerchantApiClient) SettlementModify(ctx context.Context, req ModifySettlementRequest) (err error) {
	url := fmt.Sprintf("/v3/apply4sub/sub_merchants/%s/modify-settlement", req.SubMchID)
	body, _ := json.Marshal(&req)
	_, err = c.doRequestAndVerifySignature(ctx, "POST", url, nil, body)
	if err != nil {
		return
	}
	return
}

type QuerySettlementRequest struct {
	// 特约商户号
	SubMchID string
}

type QuerySettlementResponse struct {
	// 账户类型
	AccountType string `json:"account_type"`
	// 开户银行
	AccountBank string `json:"account_bank"`
	// 开户银行全称（含支行）
	BankName string `json:"bank_name"`
	// 开户银行联行号
	BankBranchID string `json:"bank_branch_id"`
	// 银行账号
	AccountNumber string `json:"account_number"`
	// 汇款验证结果
	VerifyResult string `json:"verify_result"`
}

// 查询结算账户API
func (c MerchantApiClient) SettlementQuery(ctx context.Context, req QuerySettlementRequest) (resp *QuerySettlementResponse, err error) {
	url := fmt.Sprintf("/v3/apply4sub/sub_merchants/%s/settlement", req.SubMchID)
	res, err := c.doRequestAndVerifySignature(ctx, "GET", url, nil, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	return
}

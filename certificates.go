package wxmch_api

import (
	"context"
	"encoding/json"
)

type GetCertificatesResp struct {
	Data []struct{
		SerialNo string `json:"serial_no"`
		EffectiveTime string `json:"effective_time"`
		ExpireTime string `json:"expire_time"`
		EncryptCertificate struct{
			Algorithm string `json:"algorithm"`
			Nonce string `json:"nonce"`
			AssociatedData string `json:"associated_data"`
			Ciphertext string `json:"ciphertext"`
		} `json:"encrypt_certificate"`
	} `json:"data"`

}

// 获取平台证书列表
func (c MerchantApiClient) GetCertificates() (resp *GetCertificatesResp, err error){
	res, err := c.doRequest(context.Background(), "GET", "/v3/certificates", "", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(res), &resp)
	return
}


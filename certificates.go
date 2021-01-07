package wxmch_api

import (
	"context"
	"encoding/json"
)

type GetCertificatesResp struct {
	Data []struct {
		SerialNo           string `json:"serial_no"`
		EffectiveTime      string `json:"effective_time"`
		ExpireTime         string `json:"expire_time"`
		EncryptCertificate struct {
			Algorithm      string `json:"algorithm"`
			Nonce          string `json:"nonce"`
			AssociatedData string `json:"associated_data"`
			Ciphertext     string `json:"ciphertext"`
		} `json:"encrypt_certificate"`
		CertContent string
	} `json:"data"`
}

// 获取平台证书列表
func (c MerchantApiClient) GetCertificates() (resp *GetCertificatesResp, err error) {
	res, err := c.doRequestWithoutVerifySignature(context.Background(), "GET", "/v3/certificates", "", nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return
	}
	for i := range resp.Data {
		cert := resp.Data[i]
		encrypted := cert.EncryptCertificate
		certContent := decryptCiphertextWithGCM(encrypted.AssociatedData, encrypted.Nonce, encrypted.Ciphertext, c.apiSecret)
		resp.Data[i].CertContent = string(certContent)
	}
	return
}

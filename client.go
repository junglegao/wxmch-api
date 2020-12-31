package wxmch_api

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type MerchantApiClient struct {
	// 商户号
	mchId string
	// 商户api证书序列号
	certSerialNo string
	// apiCert 商户api证书私钥
	apiCert string
	// baseUrl
	baseUrl string
	// timeout 调用微信支付接口超时时间
	timeout time.Duration
}

const maxTimeout = 30*time.Second
const minTimeout = 1*time.Second

func NewMerchantApiClient(mchId string, certSerialNo string, apiCert string, baseUrl string, timeout time.Duration) (client MerchantApiClient) {
	if timeout > maxTimeout {
		timeout = maxTimeout
	}
	if timeout < minTimeout {
		timeout = minTimeout
	}
	client = MerchantApiClient{
		mchId:        mchId,
		certSerialNo: certSerialNo,
		apiCert: apiCert,
		baseUrl: baseUrl,
		timeout: timeout,
	}
	return
}


const AUTH_TYPE = "WECHATPAY2-SHA256-RSA2048"
const BOUNDARY = "boundary"
type ContentType string

const ContentTypePNG ContentType = "image/png"
const ContentTypeJPG ContentType = "image/jpg"
const ContentTypeBMP ContentType = "image/bmp"

func (c MerchantApiClient) formatAuthorizationHeader(nonce string, ts int, signature string) (auth string) {
	auth = fmt.Sprintf("%s mchid=\"%s\",nonce_str=\"%s\",serial_no=\"%s\",timestamp=\"%d\",signature=\"%s\"", AUTH_TYPE, c.mchId, nonce, c.certSerialNo, ts, signature)
	return
}

func (c MerchantApiClient) doRequest(ctx context.Context, method string, url string, query string, body []byte) (resp string, err error){
	nonce := RandStringBytesMaskImprSrc(10)
	ts := int(time.Now().Unix())
	signature, _ := CreateSignature(method, url, ts, nonce, body, c.apiCert)

	h := &http.Client{Timeout: c.timeout}
	requestUrl := c.baseUrl+url
	if query != "" {
		requestUrl += fmt.Sprintf("?%s", query)
	}
	req, _ := http.NewRequestWithContext(ctx, method, c.baseUrl+url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", c.formatAuthorizationHeader(nonce, ts, signature))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0")
	rawResp, err := h.Do(req)
	if err != nil {
		return
	}
	respBody, err := ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return
	}
	resp = string(respBody)
	return
}

// 表单提交上传图片专用
func (c MerchantApiClient) doFormUpload(ctx context.Context, url string, fBytes []byte, fName string, fileType ContentType) (resp string, err error) {
	nonce := RandStringBytesMaskImprSrc(10)
	ts := int(time.Now().Unix())

	hash := sha256.Sum256(fBytes)
	meta := struct {
		Filename string `json:"filename"`
		Sha256 string `json:"sha256"`
	}{
		Filename: fName,
		Sha256:   hex.EncodeToString(hash[:]),
	}
	metaStr, _ := json.Marshal(meta)
	signature, _ := CreateSignature("POST", url, ts, nonce, metaStr, c.apiCert)
	reqBody := fmt.Sprintf("--%s\r\nContent-Disposition: form-data; name=\"meta\";\r\nContent-Type: application/json\r\n\r\n%s\r\n--%s\r\nContent-Disposition: form-data; name=\"file\"; filename=\"%s\";\r\nContent-Type: %s\r\n\r\n%s\r\n--%s--", BOUNDARY, metaStr, BOUNDARY, fName, fileType, fBytes, BOUNDARY)
	h := &http.Client{Timeout: c.timeout}
	requestUrl := c.baseUrl+url
	req, _ := http.NewRequestWithContext(ctx, "POST", requestUrl, strings.NewReader(reqBody))
	req.Header.Set("Content-Type", fmt.Sprintf("multipart/form-data;boundary=%s", BOUNDARY))
	req.Header.Set("Authorization", c.formatAuthorizationHeader(nonce, ts, signature))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0")
	rawResp, err := h.Do(req)
	if err != nil {
		return
	}
	respBody, err := ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return
	}
	resp = string(respBody)
	return
}
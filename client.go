package wxmch_api

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
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
	// apiCert 商户api证书私钥明文
	apiCert string
	// baseUrl
	baseUrl string
	// timeout 调用微信支付接口超时时间
	timeout time.Duration
	// platformCertMap 平台证书map
	platformCertMap PlatformCertificatesMap
	// 平台证书编号（最新的）
	platformSerialNo string
	//api rsa private key 商户api证书私钥
	apiPriKey *rsa.PrivateKey
	// api secret
	apiSecret string
}

const maxTimeout = 30 * time.Second
const minTimeout = 1 * time.Second

func NewMerchantApiClient(mchId string, certSerialNo string, apiCert string, baseUrl string, timeout time.Duration, certMap PlatformCertificatesMap, platformNo string, apiSecret string) (client MerchantApiClient) {
	if timeout > maxTimeout {
		timeout = maxTimeout
	}
	if timeout < minTimeout {
		timeout = minTimeout
	}
	apiPriKey, err := buildRSAPrivateKey(apiCert)
	if err != nil {
		panic("错误的商户证书")
	}
	client = MerchantApiClient{
		mchId:            mchId,
		certSerialNo:     certSerialNo,
		apiCert:          apiCert,
		baseUrl:          baseUrl,
		timeout:          timeout,
		platformCertMap:  certMap,
		platformSerialNo: platformNo,
		apiPriKey:        apiPriKey,
		apiSecret:        apiSecret,
	}
	return
}

const AUTHTYPE = "WECHATPAY2-SHA256-RSA2048"
const BOUNDARY = "boundary"

type ContentType string

const ContentTypePNG ContentType = "image/png"
const ContentTypeJPG ContentType = "image/jpg"
const ContentTypeBMP ContentType = "image/bmp"

func (c MerchantApiClient) formatAuthorizationHeader(nonce string, ts int, signature string) (auth string) {
	auth = fmt.Sprintf("%s mchid=\"%s\",nonce_str=\"%s\",serial_no=\"%s\",timestamp=\"%d\",signature=\"%s\"", AUTHTYPE, c.mchId, nonce, c.certSerialNo, ts, signature)
	return
}

func (c MerchantApiClient) getPlatformPublicKey() (pubKey *rsa.PublicKey) {
	return c.platformCertMap.GetPublicKey(c.platformSerialNo)
}

// 普通http api请求
func (c MerchantApiClient) doRequest(ctx context.Context, method string, url string, query string, body []byte) (resp *http.Response, err error) {
	nonce := RandStringBytesMaskImprSrc(10)
	ts := int(time.Now().Unix())
	signature, _ := CreateSignature(method, url, ts, nonce, body, c.apiPriKey)

	h := &http.Client{Timeout: c.timeout}
	var requestUrl string
	switch query {
	case "":
		requestUrl = c.baseUrl + url
	default:
		requestUrl = c.baseUrl + url + fmt.Sprintf("?%s", query)
	}
	req, _ := http.NewRequestWithContext(ctx, method, requestUrl, bytes.NewBuffer(body))
	req.Header.Set("Authorization", c.formatAuthorizationHeader(nonce, ts, signature))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Wechatpay-Serial", c.platformSerialNo)
	resp, err = h.Do(req)

	return
}

// 没有验签功能的api请求
func (c MerchantApiClient) doRequestWithoutVerifySignature(ctx context.Context, method string, url string, query string, body []byte) (resp []byte, err error) {
	rawResp, err := c.doRequest(ctx, method, url, query, body)
	if err != nil {
		return
	}
	resp, err = ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return
	}
	err = buildErrorIfExist(rawResp.StatusCode, resp)
	if err != nil {
		resp = nil
		return
	}
	return
}

// 带验签功能的api请求
func (c MerchantApiClient) doRequestAndVerifySignature(ctx context.Context, method string, url string, query string, body []byte) (resp []byte, err error) {
	rawResp, err := c.doRequest(ctx, method, url, query, body)
	if err != nil {
		return
	}
	resp, err = ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return
	}
	err = buildErrorIfExist(rawResp.StatusCode, resp)
	if err != nil {
		resp = nil
		return
	}
	// 验证resp签名
	wechatSignature := rawResp.Header.Get("Wechatpay-Signature")
	wechatNonce := rawResp.Header.Get("Wechatpay-Nonce")
	timestamp := rawResp.Header.Get("Wechatpay-Timestamp")
	wechatSerial := rawResp.Header.Get("Wechatpay-Serial")
	if !VerifyWechatSignature(timestamp, wechatNonce, resp, wechatSignature, c.platformCertMap.GetPublicKey(wechatSerial)) {
		err = errors.New("resp签名错误")
		return
	}
	return
}

func buildErrorIfExist(statusCode int, resp []byte) (err error) {
	if statusCode != 200 && statusCode != 202 && statusCode != 204 {
		// 微信支付错误
		wechatErr := ErrBody{}
		err = json.Unmarshal(resp, &wechatErr)
		if err != nil {
			return
		}
		err = &wechatErr
		return
	}
	return
}

// 表单提交上传图片专用
func (c MerchantApiClient) doFormUpload(ctx context.Context, url string, fBytes []byte, fName string, fileType ContentType) (resp []byte, err error) {
	nonce := RandStringBytesMaskImprSrc(10)
	ts := int(time.Now().Unix())

	hash := sha256.Sum256(fBytes)
	meta := struct {
		Filename string `json:"filename"`
		Sha256   string `json:"sha256"`
	}{
		Filename: fName,
		Sha256:   hex.EncodeToString(hash[:]),
	}
	metaStr, _ := json.Marshal(meta)
	signature, _ := CreateSignature("POST", url, ts, nonce, metaStr, c.apiPriKey)
	reqBody := fmt.Sprintf("--%s\r\nContent-Disposition: form-data; name=\"meta\";\r\nContent-Type: application/json\r\n\r\n%s\r\n--%s\r\nContent-Disposition: form-data; name=\"file\"; filename=\"%s\";\r\nContent-Type: %s\r\n\r\n%s\r\n--%s--", BOUNDARY, metaStr, BOUNDARY, fName, fileType, fBytes, BOUNDARY)
	h := &http.Client{Timeout: c.timeout}
	requestUrl := c.baseUrl + url
	req, _ := http.NewRequestWithContext(ctx, "POST", requestUrl, strings.NewReader(reqBody))
	req.Header.Set("Content-Type", fmt.Sprintf("multipart/form-data;boundary=%s", BOUNDARY))
	req.Header.Set("Authorization", c.formatAuthorizationHeader(nonce, ts, signature))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0")
	rawResp, err := h.Do(req)
	if err != nil {
		return
	}
	resp, err = ioutil.ReadAll(rawResp.Body)
	if err != nil {
		return
	}
	err = buildErrorIfExist(rawResp.StatusCode, resp)
	if err != nil {
		resp = nil
		return
	}
	// 验证resp签名
	wechatSignature := rawResp.Header.Get("Wechatpay-Signature")
	wechatNonce := rawResp.Header.Get("Wechatpay-Nonce")
	timestamp := rawResp.Header.Get("Wechatpay-Timestamp")
	wechatSerial := rawResp.Header.Get("Wechatpay-Serial")
	if !VerifyWechatSignature(timestamp, wechatNonce, resp, wechatSignature, c.platformCertMap.GetPublicKey(wechatSerial)) {
		err = errors.New("resp签名错误")
		return
	}
	return
}

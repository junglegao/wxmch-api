package wxmch_api

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"strings"
)

// 生成rsa私钥
func buildRSAPrivateKey(keyContent string) (priKey *rsa.PrivateKey, err error){
	block, _ := pem.Decode([]byte(keyContent))
	if block == nil {
		err = errors.New("private key error")
		return
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	priKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		err = errors.New("private key 不是rsa格式")
		return
	}
	return
}

// 生成API调用时需要的签名
func CreateSignature(method string,  url string, ts int, nounce string, body []byte, priKey *rsa.PrivateKey) (signature string, err error) {
	// 签名前的字符串
	sBeforeSign := strings.Join([]string{method, url, fmt.Sprintf("%d", ts), nounce}, "\n") + "\n"
	if method == "GET" {
		sBeforeSign += "\n"
	} else {
		sBeforeSign += string(body) + "\n"
	}

	h := sha256.New()
	h.Write([]byte(sBeforeSign))
	hashed := h.Sum(nil)

	sign, err := rsa.SignPKCS1v15(rand.Reader, priKey, crypto.SHA256, hashed)
	if err != nil {
		return
	}
	signature = base64.StdEncoding.EncodeToString(sign)
	return
}

// 验证API返回和回调header中的微信签名
func VerifyWechatSignature(ts string, nonce string, body []byte, b64Sig string, pub *rsa.PublicKey) (pass bool) {
	// 签名前的字符串
	sBeforeSign := strings.Join([]string{ts, nonce, string(body)}, "\n") + "\n"
	signature, err := base64.StdEncoding.DecodeString(b64Sig)
	if err != nil {
		pass = false
		return
	}
	// 对签名字符串进行hash
	h := sha256.New()
	h.Write([]byte(sBeforeSign))
	hashed := h.Sum(nil)
	verifyErr := rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed, signature)
	if verifyErr != nil {
		pass = false
		return
	}
	pass = true
	return
}

// 用于平台证书解密和回调报文的解密
func decryptCiphertextWithGCM(associatedData string, nonce string, ciphertext string, apiSecret string) (plaintext []byte){
	ct, _ := base64.StdEncoding.DecodeString(ciphertext)
	nc := []byte(nonce)
	block, err := aes.NewCipher([]byte(apiSecret))
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err = aesgcm.Open(nil, nc, ct, []byte(associatedData))
	if err != nil {
		panic(err.Error())
	}
	return
}

// 从平台证书中获取公钥
func GetKeyFromCertificate(certContent string) (rsaPublicKey *rsa.PublicKey) {
	block, _ := pem.Decode([]byte(certContent))
	var cert* x509.Certificate
	cert, _ = x509.ParseCertificate(block.Bytes)
	rsaPublicKey = cert.PublicKey.(*rsa.PublicKey)
	return
}

// 从P12证书文件内容获取商户证书的公钥和私钥
func ParseP12Cert(content []byte, password string) (rsaPublicKey *rsa.PublicKey, rsaPrivateKey *rsa.PrivateKey, err error) {
	blocks, err := pkcs12.ToPEM(content, password)
	if err != nil {
		return
	}

	for _, b := range blocks {
		switch b.Type {
		case "CERTIFICATE":
			cert, _ := x509.ParseCertificate(b.Bytes)
			rsaPublicKey = cert.PublicKey.(*rsa.PublicKey)

		case "PRIVATE KEY":
			priKey, e := x509.ParsePKCS1PrivateKey(b.Bytes)
			if e != nil {
				err = e
				return
			}
			rsaPrivateKey = priKey
		}
	}
	return
}

// 敏感信息的加密
func encryptCiphertext(text string, rsaPublicKey *rsa.PublicKey) (ciphertext string) {
	secretMessage := []byte(text)
	rng := rand.Reader

	cipherdata, _ := rsa.EncryptOAEP(sha1.New(), rng, rsaPublicKey, secretMessage, nil)
	ciphertext = base64.StdEncoding.EncodeToString(cipherdata)
	return
}

// 敏感信息的解密
func decryptCiphertext(ciphertext string, rsaPrivateKey *rsa.PrivateKey) (text string, err error) {
	cipherdata, _ := base64.StdEncoding.DecodeString(ciphertext)
	rng := rand.Reader

	plaintext, err := rsa.DecryptOAEP(sha1.New(), rng, rsaPrivateKey, cipherdata, nil)
	if err != nil {
		return
	}
	text = string(plaintext)
	return
}

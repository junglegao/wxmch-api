package wxmch_api

import "crypto/rsa"

/*
	平台证书
 */

type PlatformCertificatesMap interface {
	GetPublicKey(serialNo string) (pubKey *rsa.PublicKey)
}


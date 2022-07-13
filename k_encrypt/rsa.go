package k_encrypt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

/**
 * @Author:jiangyc
 * @Description:RAS非对称加密
 * @Date:Created in 2021/6/8 9:38
**/
var rsaPrivateKey = []byte(``)

var rsaPublicKey = []byte(``)

type Rsa struct {
	PrivateKey []byte
	PublicKey  []byte
}

//CreateRsa 创建加密
func CreateRsa() *Rsa {
	return &Rsa{
		PrivateKey: rsaPrivateKey,
		PublicKey:  rsaPublicKey,
	}
}

//SetKeys 更改默认的私钥、共钥
func (r *Rsa) SetKeys(privateKey []byte, publicKey []byte) {
	r.PrivateKey = privateKey
	r.PublicKey = publicKey
}

//RsaEncrypt 加密
func (r *Rsa) RsaEncrypt(origData []byte) (string, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(r.PublicKey)
	if block == nil {
		return "", errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	//加密
	rsaEnData, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubInterface, origData, nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(rsaEnData), nil
}

//RsaDecrypt 解密
func (r *Rsa) RsaDecrypt(ciphertext string) ([]byte, error) {
	//解密
	block, _ := pem.Decode(r.PrivateKey)

	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	//先base64 decode
	rsaDeData, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, rsaDeData, nil)
}

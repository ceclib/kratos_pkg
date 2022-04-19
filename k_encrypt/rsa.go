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
var rsaPrivateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA1/lmYuW5h2EP37HOxpIkfs6ZZ2SKZ2CS3Q8Y2Gzba88PQuh+
/N+RDpEcdQR9F8fP76jsIYT2zxmpa4jnkvkWvkehuhQqOZrTKYBE81WDZ6JR3dGq
Y/aYZ1SdTLzB0Tu6Co9pRBIcbCtAaKQmRDWtTWX4zQePSFlJX1TcVAPJrTNCMK8N
4Ec3LxKwJC+wPot9tuE24BE9ZuVGcQI44VVOv9O6Dlu25wO+nN1RavPYHfSyeNl+
mB4N0phsHZR7JOw8otVcjOOJB6ojUJ7Wy9k0rXsbKpjzwYg5lBFUhJjAHNsuDztP
Y7VYgwdZaQjuseepz+3bYwY/BK8s3qouPykZ0wIDAQABAoIBAEehR/v1BRuJbaJb
bqEdR01naCmKhvzSYuM+ZpchQ4T2H7GpOrTT0WBrcGy/GF4SvzxVYjubxd/aOv6X
YXa1dn2VdGYq2fZRC0fwau3pprNweoZ/S2vsFY/v2FvjHsUucf2eouMYvqSBk9Mg
3jAxPVE2SNZ3c2YjdtpPYuPafcEbaqG+NnEIbJ46CCJftoYv76uUJC6kdtmayGfS
eiI9woeopxLZxEiirXXMnMAPfDzxEG1VkVqFK5wDXdO3+CjOxgETOx8MfzX8ygVh
6rWkS29S3JMAbtvbCi00cvPyDtf9a6guW/Ah2No9pSwn7ulo5BgWpB4oeMRy5ndp
iOl33QECgYEA/eOx5+5FZe0Ce75NhUyi7vZNk2/gWzwqRSW4XOrNKI/iSII3Lql/
wAcSAT0nhUi+rezkV/YeaUuGZ7EV7jMzRfkgkCa7NVIFHFYE5wGNUPye9Bdu4ArM
yx2WvBpKGbILM25wsstLFS2qeRG8PurDzMNBmLxZgMD4w/MFQFTPbxECgYEA2cUE
Zq83pUXA+9AqzVA+CrTZLUBpuoW6st7+dX0U3CjWyzSBuj5gNnPxub8JtJUsD68I
GtqkpYzUeAtguxqIdBLXzKelYSdqzFj3BskVUowKFjv/X2rchsTjLfCzz6mDhbo1
ILY3FtAOlcpg0FzcSNPTN/KQNOz8SbS2i4NDQqMCgYAYvzz0EnGJQdTgIuiDebjX
+gINwPXpbq/gFZEm3Wwp6/xufOLMFZFyMj47CK71euW7JkALot3L/aSYLtaoZS02
QNY/tsbAp8H/xRKtQaV4o2cK+82+4M6dbcDnjNh0MLqOxNEPXGlxIwZezq++ojt7
lfyq2pEOz8BvuRxl7N5bQQKBgQCWQl77C0N2Al9aYRRnfkqQ0KBZnOxXCxb7fOpO
3vOZxYQPyjcI7ykr6WyNIcfsllFEsR0CioK4yCeCfTdNYs1saeQFDlgKZUxHm8s3
H2BbCM2c8eEPUYjAZPHsLP2k9+4MQ6OM4q24S+8EZBPRMYG9ros7O19NZYfew5aB
RR0PyQKBgQDF2h2aWydZUQhckgIim9uHtFquLHtFouHSufUP9WONjJb0u/UUUd+0
8L01rKLAVe2bwyLNwoFFI7c6TMcMvyhEojAyEmEPmZ3eNpzd/eFSyJlnF83omnEL
tcLp5iGWPSWtL5XdT0ZUKMZybHsniEMI+MfbHdjDwtUWQrVHYQOzVg==
-----END RSA PRIVATE KEY-----`)

var rsaPublicKey = []byte(`-----BEGIN RSA PUBLIC KEY-----
MIIBCgKCAQEA1/lmYuW5h2EP37HOxpIkfs6ZZ2SKZ2CS3Q8Y2Gzba88PQuh+/N+R
DpEcdQR9F8fP76jsIYT2zxmpa4jnkvkWvkehuhQqOZrTKYBE81WDZ6JR3dGqY/aY
Z1SdTLzB0Tu6Co9pRBIcbCtAaKQmRDWtTWX4zQePSFlJX1TcVAPJrTNCMK8N4Ec3
LxKwJC+wPot9tuE24BE9ZuVGcQI44VVOv9O6Dlu25wO+nN1RavPYHfSyeNl+mB4N
0phsHZR7JOw8otVcjOOJB6ojUJ7Wy9k0rXsbKpjzwYg5lBFUhJjAHNsuDztPY7VY
gwdZaQjuseepz+3bYwY/BK8s3qouPykZ0wIDAQAB
-----END RSA PUBLIC KEY-----`)

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

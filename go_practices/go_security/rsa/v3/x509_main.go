package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"os"
	"time"
)

func main() {
	// 生成公私钥对
	caPrivkey, err := rsa.GenerateKey(rand.Reader, 2048)
	//caPrivkey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	// 生成证书模板文件
	template := newTemplate()

	// 生成自签证书(template=parent)
	rootCertDer, err := x509.CreateCertificate(rand.Reader, template, template, &caPrivkey.PublicKey, caPrivkey) //DER 格式
	if err != nil {
		panic(err)
	}

	// 将私钥编码为pkcs8格式
	privateKey, err := x509.MarshalPKCS8PrivateKey(caPrivkey)
	if err != nil {
		panic(err)
	}

	// 将私钥转为pem格式
	rootKeyFile, err := os.Create("./root.key")
	if err != nil {
		panic(err)
	}
	if err = pem.Encode(rootKeyFile, &pem.Block{Type: "EC PRIVATE KEY", Bytes: privateKey}); err != nil {
		panic(err)
	}
	rootKeyFile.Close()

	// 将证书转为pem格式
	rootCertFile, err := os.Create("./root.crt")
	if err != nil {
		panic(err)
	}
	if err = pem.Encode(rootCertFile, &pem.Block{Type: "CERTIFICATE", Bytes: rootCertDer}); err != nil {
		panic(err)
	}
	rootCertFile.Close()
}

func newTemplate() *x509.Certificate {
	max := new(big.Int).Lsh(big.NewInt(1), 128)   //把 1 左移 128 位，返回给 big.Int
	serialNumber, _ := rand.Int(rand.Reader, max) //返回在 [0, max) 区间均匀随机分布的一个随机值

	template := &x509.Certificate{
		SerialNumber: serialNumber, // SerialNumber 是 CA 颁布的唯一序列号，在此使用一个大随机数来代表它
		Subject: pkix.Name{ // 证书的主题信息
			Country:            []string{"cn"},       // 证书所属的国家
			Organization:       []string{"sunmi"},    // 证书存放的公司名称
			OrganizationalUnit: []string{"IoT"},      // 证书所属的部门名称
			Province:           []string{"ShangHai"}, // 证书签发机构所在省
			CommonName:         "*.sunmi.com",        // 证书域名
			Locality:           []string{"ShangHai"}, // 证书签发机构所在市
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth}, // 典型用法是指定叶子证书中的公钥的使用目的。它包括一系列的OID，每一个都指定一种用途。例如{id pkix 31}表示用于服务器端的TLS/SSL连接；{id pkix 34}表示密钥可以用于保护电子邮件。
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,                      // 指定了这份证书包含的公钥可以执行的密码操作，例如只能用于签名，但不能用来加密
		IsCA:                  true,                                                                       // 指示证书是不是ca证书
		BasicConstraintsValid: true,                                                                       // 指示证书是不是ca证书
	}

	return template
}

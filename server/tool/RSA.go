package tool

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

var (
	PrivateKeyStr string
	PublicKeyStr  string
)

//
// init
//  @Description: 生成RSA私钥公钥
//
func init() {
	//  使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}

	//  通过x509标准将得到的RSA私钥序列化为ASN.1的DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//  使用pem格式对x509输出的内容进行编码
	privateBlock := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: X509PrivateKey}

	PrivateKeyStr = string(pem.EncodeToMemory(privateBlock))

	//  获取公钥的数据
	publicKey := &privateKey.PublicKey

	//  X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	//  pem格式编码
	publicBlock := &pem.Block{Type: "RSA PUBLIC KEY", Bytes: X509PublicKey}

	PublicKeyStr = string(pem.EncodeToMemory(publicBlock))
}

//
// RSAEncrypt
//  @Description: 公钥加密
//  @param plainText 明文
//  @return []byte
//  @return error
//
func RSAEncrypt(plainText []byte, publicKeyStr string) ([]byte, error) {
	//pem解码
	block, _ := pem.Decode([]byte(publicKeyStr))
	if block == nil {
		return nil, errors.New("public key error")
	}

	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)

	//对明文进行加密
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
}

//
// RSADecrypt
//  @Description: 私钥解密
//  @param cipherText 密文
//  @return []byte
//  @return error
//
func RSADecrypt(cipherText []byte, privateKeyStr string) ([]byte, error) {
	//pem解码
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return nil, errors.New("private key error")
	}

	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	//对密文进行解密
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		return nil, err
	}

	//返回明文
	return plainText, nil
}

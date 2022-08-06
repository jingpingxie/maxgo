//
// @File:rsa
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/5 15:23
//
package xrsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
)

//
// @Title:loadPublicKey
// @Description:
// @Author:jingpingxie
// @Date:2022-08-06 08:15:10
// @Param:publicKey
// @Return:*rsa.PublicKey
// @Return:error
//
func loadPublicKey(publicKey string) (*rsa.PublicKey, error) {
	publicStr := "-----BEGIN PUBLIC KEY-----\n" + publicKey + "\n-----END PUBLIC KEY-----"
	block, _ := pem.Decode([]byte(publicStr))
	if block == nil {
		logs.Error("加载公钥错误")
		return nil, errors.New("加载私钥错误")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		logs.Error("加载公钥错误 %+v\n", err)
		return nil, err
	}
	pb := pubInterface.(*rsa.PublicKey)
	return pb, nil
}

//
// @Title:loadPrivateKey
// @Description:
// @Author:jingpingxie
// @Date:2022-08-06 08:15:07
// @Param:privateKey
// @Return:*rsa.PrivateKey
// @Return:error
//
func loadPrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	privateStr := "-----BEGIN PRIVATE KEY-----\n" + privateKey + "\n-----END PRIVATE KEY-----"
	block, _ := pem.Decode([]byte(privateStr))
	if block == nil {
		logs.Error("加载私钥错误")
		return nil, errors.New("加载私钥错误")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

//
// @Title:encrypt
// @Description: 公钥加密
// @Author:jingpingxie
// @Date:2022-08-05 17:23:54
// @Param:publicKey
// @Param:plainText
// @Return:[]byte
// @Return:error
//
func RsaEncrypt(publicKey string, plainText string) (string, error) {
	pb, err := loadPublicKey(publicKey)
	if err != nil {
		return "", err
	}
	// 对明文进行加密
	cipherBytes, err := rsa.EncryptPKCS1v15(rand.Reader, pb, []byte(plainText))
	if err != nil {
		logs.Error(err)
		return "", err
	}
	return hex.EncodeToString(cipherBytes), nil
}

//
// @Title:decrypt
// @Description: 私钥解密
// @Author:jingpingxie
// @Date:2022-08-05 21:22:48
// @Param:privateKey
// @Param:cipherText
// @Return:string
// @Return:error
//
func RsaDecrypt(privateKey string, cipherText string) (string, error) {
	pr, err := loadPrivateKey(privateKey)
	if err != nil {
		logs.Error("加载私钥错误 %+v\n", err)
		return "", err
	}
	//对密文进行解密
	cipherBytes, err := hex.DecodeString(cipherText)
	if err != nil {
		logs.Error(err)
		return "", err
	}
	plainBytes, err := rsa.DecryptPKCS1v15(rand.Reader, pr, cipherBytes)
	if err != nil {
		return "", err
	}
	return string(plainBytes), nil
}

//
// @Title:RsaSignWithSha256
// @Description: 私钥签名
// @Author:jingpingxie
// @Date:2022-08-05 22:07:07
// @Param:privateKey
// @Param:plainText
// @Return:string
// @Return:error
//
func RsaSignWithSha256(privateKey string, plainText string) (string, error) {
	pr, err := loadPrivateKey(privateKey)
	if err != nil {
		logs.Error("加载私钥错误 %+v\n", err)
		return "", err
	}

	h := sha256.New()
	h.Write([]byte(plainText))
	hashed := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, pr, crypto.SHA256, hashed)
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		panic(err)
	}
	return hex.EncodeToString(signature), nil
}

//
// @Title:RsaVerySignWithSha256
// @Description: 公钥验签
// @Author:jingpingxie
// @Date:2022-08-05 22:06:46
// @Param:publicKey
// @Param:plainText
// @Param:signedText
// @Return:bool
//
func RsaVerySignWithSha256(publicKey string, plainText string, signedText string) bool {
	pb, err := loadPublicKey(publicKey)
	if err != nil {
		return false
	}
	hashed := sha256.Sum256([]byte(plainText))
	signBytes, err := hex.DecodeString(signedText)
	if err != nil {
		logs.Error(err)
		return false
	}
	err = rsa.VerifyPKCS1v15(pb, crypto.SHA256, hashed[:], signBytes)
	if err != nil {
		panic(err)
	}
	return true
}

//
// @Title:convertPrivateKeyToBase64
// @Description:
// @Author:jingpingxie
// @Date:2022-08-06 08:14:54
// @Param:key
// @Return:string
//
func ConvertPrivateKeyToBase64(key *rsa.PrivateKey) string {
	derPkix := x509.MarshalPKCS1PrivateKey(key)
	return base64.StdEncoding.EncodeToString(derPkix)
}

//
// @Title:convertPublicKeyToBase64
// @Description:
// @Author:jingpingxie
// @Date:2022-08-06 08:14:51
// @Param:key
// @Return:string
// @Return:error
//
func ConvertPublicKeyToBase64(key *rsa.PublicKey) (string, error) {
	derPkix, err := x509.MarshalPKIXPublicKey(key)
	if err != nil {
		return "", err
	}
	str := base64.StdEncoding.EncodeToString(derPkix)
	return str, nil
}

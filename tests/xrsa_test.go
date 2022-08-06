//
// @File:xrsa_test
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/6 8:20
//
package test

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/astaxie/beego/logs"
	"maxgo/tools/crypto/xrsa"
	"testing"
)

func TestRsa(t *testing.T) {
	//generate rsa private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return
	}
	strPrivate := xrsa.ConvertPrivateKeyToBase64(privateKey)
	logs.Info(strPrivate)
	// generate rsa public key
	publicKey := &privateKey.PublicKey
	strPublic, err := xrsa.ConvertPublicKeyToBase64(publicKey)
	if err != nil {
		return
	}

	logs.Info(strPublic)

	plainText := "this is test"
	strCipher, err := xrsa.RsaEncrypt(strPublic, plainText)
	if err != nil {
		return
	}

	testStr, err := xrsa.RsaDecrypt(strPrivate, strCipher)
	if err != nil {
		return
	}

	logs.Info(testStr)

	str, err := xrsa.RsaSignWithSha256(strPrivate, plainText)
	if err != nil {
		return
	}
	
	ret := xrsa.RsaVerySignWithSha256(strPublic, plainText, str)
	logs.Info(ret)
}

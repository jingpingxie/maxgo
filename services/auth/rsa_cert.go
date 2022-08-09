//
// @File:auth
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/7 14:44
//
package auth

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding"
	"encoding/json"
	"fmt"
	logs "github.com/sirupsen/logrus"
	"maxgo/tools/crypto/xrsa"
)

//https://www.jianshu.com/p/f8f5a3cbaf91
var _ encoding.BinaryMarshaler = new(RsaCert)
var _ encoding.BinaryUnmarshaler = new(RsaCert)

type RsaCert struct {
	UID        string `json:"key"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
}

func (rc *RsaCert) MarshalBinary() (data []byte, err error) {
	return json.Marshal(rc)
}

func (rc *RsaCert) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, rc)

}

//
// @Title:Generate
// @Description: generate rsa cert data
// @Author:jingpingxie
// @Date:2022-08-09 12:39:52
// @Receiver:rc
// @Return:error
//
func (rc *RsaCert) Generate() error {
	//generate rsa private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		logs.Error("generate rsa private key")
		return err
	}
	rc.PrivateKey = xrsa.ConvertPrivateKeyToBase64(privateKey)
	logs.Info("generate rsa private key:" + rc.PrivateKey)
	// generate rsa public key
	publicKey := &privateKey.PublicKey
	publicBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		logs.Error("generate rsa public bytes")
		return err
	}
	strPublic, err := xrsa.ConvertPublicBytesToBase64(publicBytes)
	if err != nil {
		logs.Error("generate rsa public key")
		return err
	}
	logs.Info("generate rsa public key:" + strPublic)
	rc.PublicKey = strPublic
	rc.UID = fmt.Sprintf("%x", md5.Sum(publicBytes))
	return nil
}

//
// @Title:Decrypt
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 10:42:55
// @Receiver:rc
// @Param:cipherText
// @Return:string
// @Return:error
//
func (rc *RsaCert) Decrypt(cipherText string) (string, error) {
	return xrsa.RsaDecrypt(rc.PrivateKey, cipherText)
}

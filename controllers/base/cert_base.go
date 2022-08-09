//
// @File:CertController
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/9 10:09
//
package base

import (
	"bytes"
	"encoding/json"
	logs "github.com/sirupsen/logrus"
	"io/ioutil"
	"maxgo/services/redis_factory"
	"maxgo/tools/xstring"
)

type encryptJson struct {
	UID     string `json:"uid"`
	Encrypt string `json:"encrypt"`
}
type ICertController interface {
	PreDecrypt() error
}
type CertController struct {
	BaseController
}

//
// @Title:PreDecrypt
// @Description: decrypt the request body before api called
// @Author:jingpingxie
// @Date:2022-08-09 10:11:53
// @Receiver:cc
//
func (cc *CertController) PreDecrypt() error {
	//读取数据 body处理
	payload, err := cc.Ctx.GetRawData()
	if err != nil {
		return err
	}
	///解密body数据 请求的json是{"encryptString":{value}} value含有gcm的12字节nonce,实际长度大于32
	//{"uid":"c2dae63cc88f7c1dbafde6c8365145ff","encrypt":"HXdSB//hadVu26piFfssVWtUjadjrgIrTLCB/4s9r4A6CltnDivTV2HnBN8JuZes89n2OJZckx5MhJCRQnPxbl6LoV/TJ/5osFOBiwbyVq9aEuzmhAZcwFh7ShkSg9VCwlFRiQJfVU8jm7Zh6CHz/nLcUBLdbAK+RYqVRcIZP853qZ/7xRZ5jaCBg153K1Ipz7dHPS7CO+vI5Fm0p4YcLqtUBZ0BnOdkKnM0r1vylT5GUCFnQSYB+DxpF1ZfzosffyofYi/WyRaSEruDNPd5QBRwHScWcLjBd6p5sFpSBu4JA7XfwSJiX7FEY4vO4yrlxD/S9UWhRl0RcXPFqV4ssw=="}
	var requestData encryptJson
	err = json.Unmarshal(payload, &requestData)
	if err != nil {
		logs.Error("unmarshal payload of %s error: %s", cc.Ctx.Request.URL.Path, err)
		return err
	}
	rsaCert, err := redis_factory.GetThrowRsaCert(requestData.UID)
	if err != nil {
		return err
	}
	payloadText, err := rsaCert.Decrypt(requestData.Encrypt)
	if err != nil {
		return err
	}
	if len(payloadText) > 0 {
		payloadText = xstring.StringStrip(payloadText)
		payload = []byte(payloadText)
		cc.Ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(payload))
	}
	return nil
}

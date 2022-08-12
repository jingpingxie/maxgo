//
// @File:cert_base
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/12 9:13
//
package base

import (
	"bytes"
	"encoding/json"
	logs "github.com/sirupsen/logrus"
	"io/ioutil"
	"maxgo/services/rsa_cert"
	"maxgo/tools/xstring"
	"net/http"
)

//
// @Title:ICertBaseController
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 10:32:47
//
type ICertBaseController interface {
	PreDecrypt() error
	GetRsaCert(rsaCertKey string) (rsaCertData *rsa_cert.RsaCert, err error)
}

//
// @Title:CertBaseController
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 10:32:50
//
type CertBaseController struct {
	BaseController
}

//
// @Title:DoPreDecrypt
// @Description:对加密请求进行解密的预处理
// @Author:jingpingxie
// @Date:2022-08-12 10:31:41
// @Receiver:cc
// @Param:icb
// @Return:error
//
func (bc *BaseController) DoPreDecrypt(icb ICertBaseController) error {
	//read body data
	requestData, err := bc.Ctx.GetRawData()
	if err != nil {
		logs.Error("failed to get rawData of %s error: %s", bc.Ctx.Request.URL.Path, err)
		bc.Respond(bc.Ctx, http.StatusBadRequest, -200, "failed to get rawData")
		return err
	}
	requestMap := make(map[string]interface{})
	err = json.Unmarshal(requestData, &requestMap)
	if err != nil {
		logs.Error("failed to unmarshal rawData of %s error: %s", bc.Ctx.Request.URL.Path, err)
		bc.Respond(bc.Ctx, http.StatusUnauthorized, -300, "failed to unmarshal rawData")
		return err
	}
	if requestMap["cert_key"] == nil {
		logs.Error("no cert_key provided of %s error: %s", bc.Ctx.Request.URL.Path, err)
		bc.Respond(bc.Ctx, http.StatusUnauthorized, -400, "no cert_key provided")
		return err
	}
	if requestMap["encrypt"] == nil {
		logs.Error("no encrypt data provided of %s error: %s", bc.Ctx.Request.URL.Path, err)
		bc.Respond(bc.Ctx, http.StatusUnauthorized, -500, "no encrypt data provided")
		return err
	}
	rsaCert, err := icb.GetRsaCert(requestMap["cert_key"].(string))
	if err != nil {
		logs.Error("failed to get rsa cert, %s error: %s", bc.Ctx.Request.URL.Path, err)
		bc.Respond(bc.Ctx, http.StatusUnauthorized, -600, "failed to get rsa cert")
		return err
	}
	delete(requestMap, "cert_key")
	decryptRequestText, err := rsaCert.Decrypt(requestMap["encrypt"].(string))
	if err != nil {
		logs.Error("failed to decrypt request data, %s error: %s", bc.Ctx.Request.URL.Path, err)
		bc.Respond(bc.Ctx, http.StatusUnauthorized, -700, "failed to decrypt request data")
		return err
	}
	delete(requestMap, "encrypt")
	if len(decryptRequestText) > 0 {
		decryptRequestText = xstring.StringStrip(decryptRequestText)
		decryptRequestData := []byte(decryptRequestText)
		decryptRequestMap := make(map[string]interface{})
		err = json.Unmarshal(decryptRequestData, &decryptRequestMap)
		if err != nil {
			logs.Error("failed to unmarshal decryptRequestData of %s error: %s", bc.Ctx.Request.URL.Path, err)
			bc.Respond(bc.Ctx, http.StatusUnauthorized, -800, "failed to unmarshal decryptRequestData")
			return err
		}
		//merge decryptRequestMap and requestMap
		for k, v := range decryptRequestMap {
			requestMap[k] = v
		}
		requestData, err = json.Marshal(requestMap)
		if err != nil {
			logs.Error("failed to marshal requestMap of %s error: %s", bc.Ctx.Request.URL.Path, err)
			bc.Respond(bc.Ctx, http.StatusUnauthorized, -900, "failed to marshal requestMap")
			return err
		}
		bc.Ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestData))
	}
	return nil
}

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
	"errors"
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
	PreDecrypt() (requestMap map[string]interface{}, err error)
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
func (bc *BaseController) DoPreDecrypt(icb ICertBaseController) (requestMap map[string]interface{}, err error) {
	//read body data
	requestData, err := bc.Ctx.GetRawData()
	if err != nil {
		logs.Error("failed to get rawData of %s error: %s", bc.Ctx.Request.URL.Path, err)
		bc.Respond(http.StatusBadRequest, -200, "failed to get rawData")
		return nil, err
	}
	requestMap = make(map[string]interface{})
	err = json.Unmarshal(requestData, &requestMap)
	if err != nil {
		logs.Error("failed to unmarshal rawData of %s error: %s", bc.Ctx.Request.URL.Path, err)
		bc.Respond(http.StatusUnauthorized, -300, "failed to unmarshal rawData")
		return nil, err
	}
	certKey := bc.Ctx.GetHeader("cert_key")
	if certKey == "" {
		logs.Error("no cert_key provided of %s error: %s", bc.Ctx.Request.URL.Path, err)
		bc.Respond(http.StatusUnauthorized, -410, "no cert_key provided")
		return nil, errors.New("no cert_key provided")
	}
	if requestMap["encrypt"] == nil {
		logs.Error("no encrypt data provided of %s error: %s", bc.Ctx.Request.URL.Path, err)
		bc.Respond(http.StatusUnauthorized, -420, "no encrypt data provided")
		return nil, errors.New("no encrypt data provided")
	}
	rsaCertKey := certKey
	rsaCert, err := icb.GetRsaCert(rsaCertKey)
	if err != nil {
		logs.Error("failed to get rsa cert, %s error: %s", bc.Ctx.Request.URL.Path, err)
		bc.Respond(http.StatusUnauthorized, -600, "failed to get rsa cert")
		return nil, err
	}
	//delete(requestMap, "cert_key")
	//token := requestMap["token"]
	token := bc.Ctx.GetHeader("Authorization")
	if token != "" {
		token, err := rsaCert.Decrypt(token)
		if err != nil {
			logs.Error("failed to decrypt token data, %s error: %s", bc.Ctx.Request.URL.Path, err)
			bc.Respond(http.StatusUnauthorized, -700, "failed to decrypt token data")
			return nil, err
		}
		requestMap["token"] = token
	}
	decryptRequestText, err := rsaCert.Decrypt(requestMap["encrypt"].(string))
	if err != nil {
		logs.Error("failed to decrypt request data, %s error: %s", bc.Ctx.Request.URL.Path, err)
		bc.Respond(http.StatusUnauthorized, -710, "failed to decrypt request data")
		return nil, err
	}
	delete(requestMap, "encrypt")
	if len(decryptRequestText) > 0 {
		decryptRequestText = xstring.StringStrip(decryptRequestText)
		decryptRequestData := []byte(decryptRequestText)
		decryptRequestMap := make(map[string]interface{})
		err = json.Unmarshal(decryptRequestData, &decryptRequestMap)
		if err != nil {
			logs.Error("failed to unmarshal decryptRequestData of %s error: %s", bc.Ctx.Request.URL.Path, err)
			bc.Respond(http.StatusUnauthorized, -800, "failed to unmarshal decryptRequestData")
			return nil, err
		}
		//merge decryptRequestMap and requestMap
		for k, v := range decryptRequestMap {
			requestMap[k] = v
		}
		requestData, err = json.Marshal(requestMap)
		if err != nil {
			logs.Error("failed to marshal requestMap of %s error: %s", bc.Ctx.Request.URL.Path, err)
			bc.Respond(http.StatusUnauthorized, -900, "failed to marshal requestMap")
			return nil, err
		}
		bc.Ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestData))
		////extend valid time of rsa cert key
		//err = redis_factory.ExtendIntervalRsaCertExpireTime(rsaCertKey, rsaCert)
		//if err != nil {
		//	logs.Error("failed to extend valid time of rsa cert key of %s error: %s", bc.Ctx.Request.URL.Path, err)
		//	bc.Respond(http.StatusUnauthorized, -900, "failed to marshal requestMap")
		//	return nil, err
		//}
	}
	return requestMap, nil
}

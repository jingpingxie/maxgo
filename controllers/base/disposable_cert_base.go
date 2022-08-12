//
// @File:DisposableCertBaseController
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/9 10:09
//
package base

import (
	"maxgo/services/redis_factory"
	"maxgo/services/rsa_cert"
)

//
// @Title:DisposableCertBaseController
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 10:33:01
//
type DisposableCertBaseController struct {
	CertBaseController
}

//
// @Title:GetRsaCert
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 10:33:03
// @Receiver:cc
// @Param:rsaCertKey
// @Return:rsaCertData
// @Return:err
//
func (dcb *DisposableCertBaseController) GetRsaCert(rsaCertKey string) (rsaCertData *rsa_cert.RsaCert, err error) {
	return redis_factory.GetDisposableRsaCert(rsaCertKey)
}

//
// @Title:PreDecrypt
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 10:33:06
// @Receiver:cc
// @Return:error
//
func (dcb *DisposableCertBaseController) PreDecrypt() error {
	return dcb.CertBaseController.DoPreDecrypt(dcb)
}

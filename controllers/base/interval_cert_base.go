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
// @Title:IntervalCertBaseController
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 10:33:18
//
type IntervalCertBaseController struct {
	CertBaseController
}

//
// @Title:GetRsaCert
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 10:33:28
// @Receiver:cc
// @Param:rsaCertKey
// @Return:rsaCertData
// @Return:err
//
func (icb *IntervalCertBaseController) GetRsaCert(rsaCertKey string) (rsaCertData *rsa_cert.RsaCert, err error) {
	return redis_factory.GetIntervalRsaCert(rsaCertKey)
}

//
// @Title:PreDecrypt
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 10:33:31
// @Receiver:cc
// @Return:error
//
func (icb *IntervalCertBaseController) PreDecrypt() error {
	return icb.CertBaseController.DoPreDecrypt(icb)
}

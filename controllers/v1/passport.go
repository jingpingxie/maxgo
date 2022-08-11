//
// @File:passport
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/7 14:15
//
package v1

import (
	"maxgo/controllers/base"
	"maxgo/routers"
	"maxgo/services/redis_factory"
	"net/http"
)

//
// @Title:init
// @Description:
// @Author:jingpingxie
// @Date:2022-08-07 14:16:22
//
func init() {
	routers.Register(&PassportController{})
}

//
// @Title:PassportController
// @Description:
// @Author:jingpingxie
// @Date:2022-08-07 14:16:24
//
type PassportController struct {
	base.BaseController
}

//
// @Title:Get_IntervalCertData
// @Description:
// @Author:jingpingxie
// @Date:2022-08-07 14:16:27
// @Receiver:uc
//
func (uc *PassportController) Get_IntervalCertData() {
	uid, rsaCertData := redis_factory.GenerateIntervalRsaCert()
	uc.Respond(uc.Ctx, http.StatusOK, 0, "", map[string]string{"uid": uid, "cert": rsaCertData.PublicKey})
}

//
// @Title:Get_DisposableCert
// @Description: get disposable cert data
// @Author:jingpingxie
// @Date:2022-08-09 08:50:08
// @Receiver:uc
//
func (uc *PassportController) Get_DisposableCert() {
	uid, rsaCertData := redis_factory.GenerateDisposableRsaCert()
	uc.Respond(uc.Ctx, http.StatusOK, 0, "", map[string]string{"uid": uid, "cert": rsaCertData.PublicKey})
}

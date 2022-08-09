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
// @Title:Get_Login
// @Description:
// @Author:jingpingxie
// @Date:2022-08-07 14:16:27
// @Receiver:uc
//
func (uc *PassportController) Get_HourCertData() {
	uid, rsaCertData := redis_factory.GetHourRsaCert()
	uc.Respond(uc.Ctx, http.StatusOK, 0, "", map[string]string{"uid": uid, "cert": rsaCertData.PublicKey})
}

//
// @Title:Get_ThrowCert
// @Description: get throwaway cert data
// @Author:jingpingxie
// @Date:2022-08-09 08:50:08
// @Receiver:uc
//
func (uc *PassportController) Get_ThrowCert() {
	uid, rsaCertData := redis_factory.GenerateThrowRsaCert()
	uc.Respond(uc.Ctx, http.StatusOK, 0, "", map[string]string{"uid": uid, "cert": rsaCertData.PublicKey})
}

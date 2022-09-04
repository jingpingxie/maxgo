//
// @File:user
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/6 13:58
//
package disposable_cert

import (
	logs "github.com/sirupsen/logrus"
	"maxgo/controllers/base"
	"maxgo/routers"
	"maxgo/services/auth"
	"net/http"
)

func init() {
	routers.Register(&UserController{})
}

//
// @Title:UserController
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:30:56
//
type UserController struct {
	base.DisposableCertBaseController
}

//
// @Title:Login
// @Description: user login
// @Author:jingpingxie
// @Date:2022-08-04 14:57:20
// @Receiver:uc
//
func (uc *UserController) Post_Login() {
	lr := new(auth.LoginRequest)
	if err := uc.Ctx.ShouldBind(lr); err != nil {
		logs.Error("unmarshal payload of %s error: %s", uc.Ctx.Request.URL.Path, err)
		uc.Respond(http.StatusBadRequest, -100, err.Error())
		return
	}

	logs.Info("account:%s password:%s is login", lr.Account, lr.Password)
	statusCode, lrt, err := auth.DoLogin(lr, uc.Ctx.ClientIP())
	if err != nil {
		uc.Respond(statusCode, -200, err.Error())
		return
	}
	uc.Ctx.Header("Access-Control-Expose-Headers", "Authorization,cert_key,public_key")
	uc.Ctx.Header("Authorization", lrt.Token)
	uc.Ctx.Header("cert_key", lrt.RsaCertKey)
	uc.Ctx.Header("public_key", lrt.RsaPublicKey)
	uc.Respond(http.StatusOK, 0, "succeed to login", lrt.UserResponse)
}

//
// @Title:RegisterUser
// @Description: register user
// @Author:jingpingxie
// @Date:2022-08-04 14:54:33
// @Receiver:uc
//
func (uc *UserController) Post_Register() {
	ur := new(auth.RegisterRequest)
	if err := uc.Ctx.ShouldBind(ur); err != nil {
		logs.Error("unmarshal payload of %s error: %s", uc.Ctx.Request.URL.Path, err)
		uc.Respond(http.StatusBadRequest, -100, err.Error())
		return
	}
	statusCode, lrt, err := auth.DoRegister(ur, uc.Ctx.ClientIP())
	if err != nil {
		uc.Respond(statusCode, -200, err.Error())
		return
	}
	uc.Ctx.Header("Access-Control-Expose-Headers", "Authorization,cert_key,public_key")
	uc.Ctx.Header("Authorization", lrt.Token)
	uc.Ctx.Header("cert_key", lrt.RsaCertKey)
	uc.Ctx.Header("public_key", lrt.RsaPublicKey)
	uc.Respond(http.StatusOK, 0, "succeed to register", lrt.UserResponse)
}

//
// @File:user
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/6 13:58
//
package cert

import (
	logs "github.com/sirupsen/logrus"
	user2 "maxgo/common/user"
	"maxgo/controllers/base"
	"maxgo/routers"

	//"maxgo/routers"
	"maxgo/services/auth"
	"net/http"
)

func init() {
	routers.Register(&UserController{})
}

type UserController struct {
	base.CertBaseController
}

//
// @Title:Login
// @Description: user login
// @Author:jingpingxie
// @Date:2022-08-04 14:57:20
// @Receiver:uc
//
func (uc *UserController) Post_Login() {
	lr := new(user2.LoginRequest)
	if err := uc.Ctx.ShouldBind(lr); err != nil {
		logs.Error("unmarshal payload of %s error: %s", uc.Ctx.Request.URL.Path, err)
		uc.Respond(uc.Ctx, http.StatusBadRequest, -100, err.Error())
		return
	}

	logs.Info("account:%s password:%s is login", lr.Account, lr.Password)
	statusCode, lrs, err := auth.DoLogin(lr)
	if err != nil {
		uc.Respond(uc.Ctx, statusCode, -200, err.Error())
		return
	}
	uc.Ctx.Header("Authorization", lrs.Token) // set token into header
	uc.Respond(uc.Ctx, http.StatusOK, 0, "", lrs)
}

//
// @Title:RegisterUser
// @Description: register user
// @Author:jingpingxie
// @Date:2022-08-04 14:54:33
// @Receiver:uc
//
func (uc *UserController) Post_Register() {
	ur := new(user2.UserRequest)
	if err := uc.Ctx.ShouldBind(ur); err != nil {
		logs.Error("unmarshal payload of %s error: %s", uc.Ctx.Request.URL.Path, err)
		uc.Respond(uc.Ctx, http.StatusBadRequest, -100, err.Error(), nil)
		return
	}
	statusCode, registerUser, err := auth.DoRegisterUser(ur)
	if err != nil {
		uc.Respond(uc.Ctx, http.StatusBadRequest, statusCode, "", err.Error())
		return
	}
	uc.Respond(uc.Ctx, http.StatusOK, 0, "", registerUser)
}

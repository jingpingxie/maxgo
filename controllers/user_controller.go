package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	user2 "maxgo/common/user"
	"maxgo/services/auth"
	"maxgo/tools/jsonUtils"
	"net/http"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) respond(code int, message string, data ...interface{}) {
	u.Ctx.Output.SetStatus(code)
	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}
	u.Data["json"] = struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}{
		Code:    code,
		Message: message,
		Data:    d,
	}
	u.ServeJSON()
}

//
// @Title:Login
// @Description: user login
// @Author:jingpingxie
// @Date:2022-08-04 14:57:20
// @Receiver:uc
//
func (uc *UserController) Login() {
	lr := new(user2.LoginRequest)
	if err := jsonUtils.Unmarshal(uc.Ctx.Input.RequestBody, lr); err != nil {
		logs.Error("unmarshal payload of %s error: %s", uc.Ctx.Request.URL.Path, err)
		uc.respond(http.StatusBadRequest, err.Error())
		return
	}
	logs.Info("account:%s password:%s is login", lr.Account, lr.Password)
	statusCode, lrs, err := auth.DoLogin(lr)
	if err != nil {
		uc.respond(statusCode, err.Error())
		return
	}
	uc.Ctx.Output.Header("Authorization", lrs.Token) // set token into header
	uc.respond(http.StatusOK, "", lrs)
}

//
// @Title:RegisterUser
// @Description: register user
// @Author:jingpingxie
// @Date:2022-08-04 14:54:33
// @Receiver:uc
//
func (uc *UserController) Register() {
	ur := new(user2.UserRequest)
	if err := jsonUtils.Unmarshal(uc.Ctx.Input.RequestBody, ur); err != nil {
		logs.Error("unmarshal payload of %s error: %s", uc.Ctx.Request.URL.Path, err)
		uc.respond(http.StatusBadRequest, err.Error())
		return
	}
	statusCode, registerUser, err := auth.DoRegisterUser(ur)
	if err != nil {
		uc.respond(statusCode, err.Error())
		return
	}
	uc.respond(http.StatusOK, "", registerUser)
}
func (uc *UserController) Logout() {

}

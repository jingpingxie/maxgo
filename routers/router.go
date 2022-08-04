package routers

import (
	"github.com/astaxie/beego"
	"maxgo/controllers"
)

//
// @Title:init
// @Description:
// @Author:jingpingxie
// @Date:2022-08-02 11:57:14
//
func init() {
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/user",
			beego.NSRouter("/login", &controllers.UserController{}, "post:Login"),
			beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),
			beego.NSRouter("/logout", &controllers.UserController{}, "post:Logout"),
		),
	)
	beego.AddNamespace(ns)
}

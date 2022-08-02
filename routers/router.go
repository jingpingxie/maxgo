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
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/login", &controllers.UserController{}, "*:Login")
	beego.Router("/user/register", &controllers.UserController{}, "*:Register")
	beego.Router("/user/logout", &controllers.UserController{}, "*:Logout")
}

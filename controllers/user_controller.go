package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Login() {
	account := c.Input().Get("account")
	password := c.Input().Get("password")
	logs.Info("你好" + account + password)
}
func (c *UserController) Register() {
	account := c.Input().Get("account")
	password := c.Input().Get("password")
	logs.Info("你好" + account + password)
}
func (c *UserController) Logout() {
	account := c.Input().Get("account")
	password := c.Input().Get("password")
	logs.Info("你好" + account + password)
}

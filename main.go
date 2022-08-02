package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"maxgo/db"
	_ "maxgo/models"
	_ "maxgo/routers"
	"maxgo/tools/snowflake"
)

func init() {
	//初始化数据库
	db.InitDb()
}
func main() {
	s, _ := snowflake.GenerateSnowflakeId()
	logs.Info(s)
	beego.Run()
}

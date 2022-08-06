package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "maxgo/models"
	_ "maxgo/routers"
	"maxgo/services"
	"maxgo/tools/snowflake"
)

func init() {
	//初始化数据库
	services.InitDb()
}
func main() {
	//xrsa.GetPublicKey()
	s, _ := snowflake.GenerateSnowflakeId()
	logs.Info(s)
	beego.Run()
	sqlDb, _ := services.Db.DB()
	defer sqlDb.Close()
}

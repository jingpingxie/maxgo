package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"maxgo/dao"
	_ "maxgo/models"
	_ "maxgo/routers"
	"maxgo/tools/snowflake"
)

func init() {
	//初始化数据库
	dao.InitDb()
}
func main() {
	xrsa.GetPublicKey()
	s, _ := snowflake.GenerateSnowflakeId()
	logs.Info(s)
	beego.Run()
	sqlDb, _ := dao.Db.DB()
	defer sqlDb.Close()
}

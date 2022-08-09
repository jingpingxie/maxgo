//
// @File:main
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/6 13:53
//
package main

import (
	logs "github.com/sirupsen/logrus"
	_ "maxgo/controllers"
	_ "maxgo/controllers/cert"
	"maxgo/routers"
	"maxgo/services"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logs.SetFormatter(&logs.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logs.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logs.SetLevel(logs.DebugLevel)
	//初始化数据库
	services.InitDb()
}

func main() {
	//加载路由
	r := routers.InitRouter()
	_ = r.Run("localhost:9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	sqlDb, _ := services.Db.DB()
	defer sqlDb.Close()
}

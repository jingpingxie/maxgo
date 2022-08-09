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

	//result := &auth.RsaCert{}
	//err := services.ClientRedis.Get("test11").Scan(result)
	//if err != nil {
	//	panic(err)
	//}

	//now := time.Now()
	//hourTime := now.Format("2006010215") //小时取整
	//logs.Info(hourTime)

	logs.WithFields(logs.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logs.WithFields(logs.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	//logs.WithFields(logs.Fields{
	//	"omg":    true,
	//	"number": 100,
	//}).Fatal("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := logs.WithFields(logs.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")

	logs.Trace("Something very low level.")
	logs.Debug("Useful debugging information.")
	logs.Info("Something noteworthy happened!")
	logs.Warn("You should probably take a look at this.")
	//logs.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	//logs.Fatal("Bye.")
	// Calls panic() after logging
	//logs.Panic("I'm bailing.")
	//加载路由
	r := routers.InitRouter()
	_ = r.Run("localhost:9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	sqlDb, _ := services.Db.DB()
	defer sqlDb.Close()
}

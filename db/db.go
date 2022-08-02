package db

import (
	"fmt"
	"github.com/astaxie/beego"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"maxgo/models"
	"os"
	"time"

	"gorm.io/driver/mysql"
)

func InitDb() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 彩色打印
		},
	)

	log.Println("db init start")

	//注册默认数据库
	host := beego.AppConfig.String("db::host")
	port, _ := beego.AppConfig.Int("db::port")
	dbname := beego.AppConfig.String("db::databaseName")
	user := beego.AppConfig.String("db::userName")
	pwd := beego.AppConfig.String("db::password")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", user, pwd, host, port, dbname)
	fmt.Print(dataSource)

	//记录sql日志

	db, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {

		fmt.Println("conn mysql db so give up. err:", err)

		os.Exit(1)
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxOpenConns(30) //最大支持的连接数
	sqlDb.SetMaxIdleConns(10) // 最大空闲的连接数
	sqlDb.SetConnMaxIdleTime(time.Minute)

	// 自动建表
	// 根据模型创建数据库(执行数据库迁移文件)
	// 第二个参数：最容易出错的地方，如果值为ture时，表已经存在并且表中有值的情况下，它会先删除原来的表，然后重新创建，这样原表中的数据就全部丢失了。
	// 第三个参数：是否输出建表的sql日志 true:输出 false：不输出
	if err := db.AutoMigrate(
		&models.Admin{},
		&models.Company{},
		&models.Role{},
		&models.User{},
		&models.UserCommonWechat{},
		&models.UserCompanyWechat{},
		&models.UserDetail{},
		&models.UserIdCard{},
		&models.UserVisit{},
		&models.WechatCompany{},
		&models.WechatMenu{},
	); err != nil {
		log.Fatal(err)
	}
	log.Println("db init start successful")
}

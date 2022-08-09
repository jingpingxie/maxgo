//
// @File:conf
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/6 15:03
//
package conf

import (
	"fmt"
	logs "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	AppConf   *AppConfig
	SqlConn   string
	RedisConf *RedisConfig
	DBConf    *DBConfig
}

type AppConfig struct {
	Name string
	Port string
	Mode string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	Db       int
}

type DBConfig struct {
	Host         string
	Port         int32
	UserName     string
	Password     string
	DatabaseName string
}

//
// @Title:init
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 12:46:10
//
func init() {
	viper.SetConfigType("yaml") //设置配置文件格式
	viper.AddConfigPath("conf") //设置配置文件的路径
	viper.SetConfigName("app")  //设置配置文件名
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logs.Error("can't find config file")
		} else {
			logs.Error("error for config setting")
		}
	}
	//打印获取到的配置文件的key
	fmt.Println(viper.AllKeys())
}

//
// @Title:InitConf
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 12:46:06
// @Return:*Config
//
func InitConf() *Config {
	//返回配置文件的数据
	return &Config{
		AppConf: &AppConfig{
			Name: viper.GetString("app.name"),
			Port: viper.GetString("app.port"),
			Mode: viper.GetString("app.mode"),
		},
		SqlConn: viper.GetString("mysql.conn"),
		RedisConf: &RedisConfig{
			Host:     viper.GetString("redis.host"),
			Port:     viper.GetString("redis.port"),
			Password: viper.GetString("redis.password"),
			Db:       viper.GetInt("redis.db"),
		},
	}
}

//
// @Title:GetDBConf
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 12:46:01
// @Return:*DBConfig
//
func GetDBConf() *DBConfig {
	return &DBConfig{
		Host:         viper.GetString("db.host"),
		Port:         viper.GetInt32("db.port"),
		UserName:     viper.GetString("db.userName"),
		Password:     viper.GetString("db.password"),
		DatabaseName: viper.GetString("db.databaseName"),
	}
}

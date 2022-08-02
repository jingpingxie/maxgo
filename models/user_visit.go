//
// @Title:models
// @File:${name}
// @Description:
// @Author:jingpingxie
// @Date:2022-08-01 21:50:00
//
package models

import (
	_ "github.com/go-sql-driver/mysql" //导入数据库驱动
	"time"
)

//
// @Title:UserVisit
// @Description:
// @Author:jingpingxie
// @Date:2022-08-02 11:56:53
//
type UserVisit struct {
	UserId        uint64    `gorm:"Column:user_id;PrimaryKey:true;AutoIncrement:false;NotNull:true;Comment:user_id;" json:"user_id"`
	LastLoginTime time.Time `gorm:"Column:last_login_time;Type:datetime;null;Comment:最近登录时间;" json:"last_login_time"`
	LastIp        time.Time `gorm:"Column:last_ip;Type:datetime;null;Comment:最近登录ip地址;" json:"last_ip"`
	VisitCount    uint16    `gorm:"Column:visit_count;Comment:登录次数;" json:"visit_count"`
	ModelTime     ModelTime `gorm:"Embedded;"`
}

func init() {
}

//
// @Title:TableName
// @Description:自定义表名 (默认模型名小写)
// @Author:jingpingxie
// @Date:2022-08-01 21:41:31
// @Receiver:uv
// @Return:string
//
func (uv *UserVisit) TableName() string {
	return "user_visit"
}

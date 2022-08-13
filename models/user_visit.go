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
	UserVisitID uint64    `gorm:"Column:user_visit_id;PrimaryKey:true;AutoIncrement:false;NotNull:true;Comment:user_visit_id;" json:"user_visit_id"`
	UserID      uint64    `gorm:"Column:user_id;Index:user_id_idx;AutoIncrement:false;NotNull:true;Comment:user_id;" json:"user_id"`
	LoginTime   time.Time `gorm:"Column:login_time;Type:datetime;null;Comment:登录时间;" json:"login_time"`
	LogoutTime  time.Time `gorm:"Column:logout_time;Type:datetime;null;Comment:登出时间;" json:"logout_time"`
	VisitIp     string    `gorm:"Column:visit_ip;Type:varchar(40);null;Comment:登录ip地址;" json:"visit_ip"`
	ModelTime   ModelTime `gorm:"Embedded;"`
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

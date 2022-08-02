package models

import (
	_ "github.com/go-sql-driver/mysql" //导入数据库驱动
)

//
// @Title:UserIdCard
// @Description:
// @Author:jingpingxie
// @Date:2022-08-02 11:46:03
//
type UserIdCard struct {
	UserId    uint64    `gorm:"Column:user_id;PrimaryKey:true;AutoIncrement:false;NotNull:true;Comment:user_id;" json:"user_id"`
	IdNumber  string    `gorm:"Column:id_number;Type:varchar(50);null;Comment:身份证号码;" json:"id_number"`
	FrontUrl  string    `gorm:"column:front_url;Comment:身份证正面图片URL;" json:"front_url"`
	BackUrl   string    `gorm:"column:back_url;Comment:身份证反面图片URL;" json:"back_url"`
	ModelTime ModelTime `gorm:"Embedded;"`
}

//
// @Title:init
// @Description:
// @Author:jingpingxie
// @Date:2022-08-02 11:51:30
//
func init() {
}

//
// @Title:TableName
// @Description: 自定义表名 (默认模型名小写)
// @Author:jingpingxie
// @Date:2022-08-02 11:52:12
// @Receiver:uic
// @Return:string
//
func (uic *UserIdCard) TableName() string {
	return "user_id_card"
}

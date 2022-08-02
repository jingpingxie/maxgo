package models

import (
	_ "github.com/go-sql-driver/mysql" //导入数据库驱动
)

//
// @Title:WechatCompany
// @Description:
// @Author:jingpingxie
// @Date:2022-08-02 11:57:02
//
type WechatCompany struct {
	WechatCompanyId uint64    `gorm:"Column:wechat_company_id;PrimaryKey:true;AutoIncrement:false;NotNull:true;Comment:wechat_company_id;" json:"wechat_company_id"`
	CompanyId       uint64    `gorm:"Column:company_id;AutoIncrement:false;NotNull:true;Comment:company_id;" json:"company_id"`
	Name            string    `gorm:"Column:name;Type:varchar(50);NotNull:true;Comment:公众号名称;" json:"name"`
	OrgId           string    `gorm:"Column:orgid;Type:varchar(20);NotNull:true;Comment:公众号原始id;" json:"orgid"`
	Token           string    `gorm:"Column:token;Type:varchar(20);NotNull:true;" json:"token"`
	AppId           string    `gorm:"Column:appid;Type:varchar(30);Unique:true;NotNull:true;" json:"appid"`
	AppSecret       string    `gorm:"Column:appsecret;Type:varchar(50);NotNull:true;" json:"appsecret"`
	WechatType      string    `gorm:"Column:wechat_type;Type:varchar(20);NotNull:true;Comment:公众号类型 0 微信公众号 1 微信开放平台 2 微信小程序;" json:"wechat_type"`
	IsDefault       uint8     `gorm:"Column:is_default;Comment:1家企业只允许使用1个公众号;" json:"is_default"`
	ModelTime       ModelTime `gorm:"Embedded;"`
}

//
// @Title:init
// @Description:
// @Author:jingpingxie
// @Date:2022-08-01 21:37:00
//
func init() {
}

//
// @Title:TableName
// @Description: 自定义表名 (默认模型名小写)
// @Author:jingpingxie
// @Date:2022-08-02 11:53:43
// @Receiver:wc
// @Return:string
//
func (wc *WechatCompany) TableName() string {
	return "wechat_company"
}

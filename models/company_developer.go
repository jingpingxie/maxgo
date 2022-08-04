//
// @File:company_developer
// @Version:1.0.0
// @Description:企业开发者信息
// @Author:jingpingxie
// @Date:2022/8/2 18:17
//
package models

//https://blog.csdn.net/CSDN_WYL2016/article/details/124696845
type CompanyDeveloper struct {
	CompanyID uint64    `gorm:"Column:company_id;PrimaryKey:true;AutoIncrement:false;NotNull:true;Comment:企业ID;" json:"company_id"`
	AppID     string    `gorm:"Column:app_id;Type:varchar(20);Unique:true;NotNull:true;Comment:app_id;" json:"app_id"`
	AppSecret string    `gorm:"Column:app_secret;Type:varchar(50);Unique:true;NotNull:true;Comment:app_secret;" json:"app_secret"`
	ModelTime ModelTime `gorm:"Embedded;"`
}

func init() {
}

//
// @Title:TableName
// @Description: 自定义表名 (默认模型名小写)
// @Author:jingpingxie
// @Date:2022-08-02 14:15:01
// @Receiver:a
// @Return:string
//
func (c *CompanyDeveloper) TableName() string {
	return "company_developer"
}

package models

import (
	_ "github.com/go-sql-driver/mysql" //导入数据库驱动
)

//
// @Title:Admin
// @Description:
// @Author:jingpingxie
// @Date:2022-08-02 11:55:01
//
type Admin struct {
	UserId      uint64    `gorm:"Column:user_id;PrimaryKey:true;AutoIncrement:false;NotNull:true;Comment:user_id;" json:"user_id"`
	CompanyId   uint64    `gorm:"Column:company_id;AutoIncrement:false;NotNull:true;Comment:company_id;" json:"company_id"`
	RoleId      uint8     `gorm:"Column:role_id;NotNull:true;Comment:角色id" json:"role_id"`
	PurviewList string    `gorm:"Column:purview_list;Type:text;NotNull:true;Comment:补充的管理员权限;" json:"purview_list"`
	ModelTime   ModelTime `gorm:"Embedded;"`
}

func init() {
}

// TableName 自定义表名 (默认模型名小写)
func (a *Admin) TableName() string {
	return "admin"
}

// TableUnique 联合唯一键
func (a *Admin) TableUnique() [][]string {
	return [][]string{
		{"user_id", "company_id"},
	}
}

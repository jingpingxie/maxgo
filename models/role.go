package models

import (
	_ "github.com/go-sql-driver/mysql" //导入数据库驱动
)

//
// @Title:Role
// @Description:
// @Author:jingpingxie
// @Date:2022-08-02 11:55:14
//
type Role struct {
	RoleID      uint64    `gorm:"Column:role_id;PrimaryKey:true;AutoIncrement:false;NotNull:true;Comment:角色ID;" json:"role_id"`
	CompanyID   uint64    `gorm:"Column:company_id;AutoIncrement:false;NotNull:true;UniqueIndex:cn;Comment:company_id;" json:"company_id"`
	Name        string    `gorm:"Column:name;Type:varchar(20);NotNull:true;UniqueIndex:cn;Comment:角色名称;" json:"name"`
	Describe    string    `gorm:"Column:describe;Type:varchar(50);Comment:角色描述;" json:"describe"`
	PurviewList string    `gorm:"Column:purview_list;Type:text;NotNull:true;Comment:管理员权限;" json:"purview_list"`
	Editor      uint64    `gorm:"Column:editor;NotNull:true;Comment:操作员用户id;" json:"editor"`
	ModelTime   ModelTime `gorm:"Embedded;"`
}

//
func init() {
}

//
// @Title:TableName
// @Description: 自定义表名 (默认模型名小写)
// @Author:jingpingxie
// @Date:2022-08-02 11:55:20
// @Receiver:r
// @Return:string
//
func (r *Role) TableName() string {
	return "role"
}

package models

import (
	_ "github.com/go-sql-driver/mysql" //导入数据库驱动
)

//
// @Title:WechatMenu
// @Description:
// @Author:jingpingxie
// @Date:2022-08-02 11:57:07
//
type WechatMenu struct {
	WechatMenuID    uint64    `gorm:"Column:wechat_menu_id;PrimaryKey:true;AutoIncrement:false;NotNull:true;Comment:wechat_menu_id;" json:"wechat_menu_id"`
	WechatCompanyID uint64    `gorm:"Column:wechat_company_id;NotNull:true;Comment:wechat_company_id;" json:"wechat_company_id"`
	ParentID        uint64    `gorm:"Column:parent_id;Comment:父级菜单id;" json:"parent_id"`
	Name            string    `gorm:"Column:name;Type:varchar(30);NotNull:true;Comment:菜单标题;" json:"name"`
	Type            string    `gorm:"Column:type;Type:varchar(20);NotNull:true;Comment:菜单类型 view click;" json:"type"`
	Url             string    `gorm:"Column:url;Type:varchar(100);NotNull:true;" json:"url"`
	Key             string    `gorm:"Column:key;Type:varchar(30);NotNull:true;" json:"key"`
	Sort            int8      `gorm:"Column:sort;" json:"sort"`
	ModelTime       ModelTime `gorm:"Embedded;"`
}

//
// @Title:init
// @Description:
// @Author:jingpingxie
// @Date:2022-08-02 11:53:06
//
func init() {
}

//
// @Title:TableName
// @Description:自定义表名 (默认模型名小写)
// @Author:jingpingxie
// @Date:2022-08-02 11:53:12
// @Receiver:wm
// @Return:string
//
func (wm *WechatMenu) TableName() string {
	return "wechat_menu"
}

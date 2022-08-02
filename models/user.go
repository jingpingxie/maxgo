package models

import (
	_ "github.com/go-sql-driver/mysql" //导入数据库驱动
	"time"
)

//
// @Title:User
// @Description:
// @Author:jingpingxie
// @Date:2022-08-02 11:56:22
//
type User struct {
	UserId        uint64    `gorm:"Column:user_id;PrimaryKey:true;AutoIncrement:false;NotNull:true;Comment:user_id;" json:"user_id"`
	UserName      string    `gorm:"Column:user_name;Type:varchar(100);NotNull:true;Comment:用户名;" json:"user_name"` //用户名,可重复
	NickName      string    `gorm:"Column:nick_name;Type:varchar(100);Comment:用户昵称;" json:"nick_name"`
	RealName      string    `gorm:"Column:real_name;Type:varchar(100);Comment:真实姓名;" json:"real_name"`
	CountryCode   uint8     `gorm:"Column:country_code;null;Comment:手机的国家区号;" json:"country_code"`
	Mobile        string    `gorm:"Column:mobile;Type:varchar(20);Unique:true;NotNull:true;Comment:用户手机号;" json:"mobile"`
	ContactMobile string    `gorm:"Column:contact_mobile;Type:varchar(20);null;Comment:联系人手机号;" json:"contact_mobile"`
	Password      string    `gorm:"Column:password;Type:varchar(30);NotNull:true;Comment:加密的密码;" json:"password"`
	Salt          string    `gorm:"Column:salt;Type:varchar(10);null;Comment:密码加盐;" json:"salt"`
	Email         string    `gorm:"Column:email;Type:varchar(100);null;Comment:邮箱;" json:"email"`
	Birthday      time.Time `gorm:"Column:birthday;Type:datetime;null;Comment:生日;" json:"birthday"`
	WechatUnionId string    `gorm:"Column:wechat_union_id;Type:varchar(30);null;Comment:微信union id;" json:"wechat_union_id"`
	ParentUserId  uint64    `gorm:"Column:parent_user_id;Comment:介绍人;" json:"parent_user_id"`
	Gender        byte      `gorm:"Column:gender;default(1);Comment:性别：1 男 2 女;" json:"gender"`
	Avatar        string    `gorm:"Column:avatar;Type:varchar(100);null;Comment:头像url;" json:"avatar"`
	QQ            string    `gorm:"Column:qq;Type:varchar(20);null;Comment:qq号码;" json:"qq"`
	Money         float32   `gorm:"Column:money;Type:decimal(8,2);Comment:金额;" json:"money"`
	ModelTime     ModelTime `gorm:"Embedded;"`
}

func init() {
}

//
// @Title:TableName
// @Description: 自定义表名 (默认模型名小写)
// @Author:jingpingxie
// @Date:2022-08-02 11:55:55
// @Receiver:u
// @Return:string
//
func (u *User) TableName() string {
	return "user"
}

//
// @Title:TableEngine
// @Description: 设置引擎为 INNODB
// @Author:jingpingxie
// @Date:2022-08-02 11:55:45
// @Receiver:u
// @Return:string
//
func (u *User) TableEngine() string {
	return "INNODB"
}

//// TableUnique 联合唯一键
//func (u *User) TableUnique() [][]string {
//	return [][]string{
//		{"mobile", "user_name"},
//	}
//}
//
//// TableIndex 普通联合 索引
//func (u *User) TableIndex() [][]string {
//	return [][]string{
//		{"gender", "user_name"},
//	}
//}

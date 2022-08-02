package models

//
// @Title:UserDetail
// @Description:
// @Author:jingpingxie
// @Date:2022-08-02 11:56:45
//
type UserDetail struct {
	UserId    uint64    `gorm:"Column:user_id;PrimaryKey:true;AutoIncrement:false;NotNull:true;Comment:用户id;" json:"user_id"`
	Resume    string    `gorm:"Column:resume;Type:longtext;Comment:个人简历;" json:"resume"`
	ModelTime ModelTime `gorm:"Embedded;"`
}

func init() {
}

//
// @Title:TableName
// @Description:自定义表名 (默认模型名小写)
// @Author:jingpingxie
// @Date:2022-08-02 11:54:08
// @Receiver:ud
// @Return:string
//
func (ud *UserDetail) TableName() string {
	return "user_detail"
}

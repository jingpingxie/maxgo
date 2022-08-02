package models

//
// @Title:UserCommonWechat
// @Description: 用户关注公众号后拉取用户的信息
// @Author:jingpingxie
// @Date:2022-08-02 11:56:28
//
type UserCommonWechat struct {
	UserWechatId uint64    `gorm:"Column:user_wechat_id;PrimaryKey:true;AutoIncrement:false;Comment:微信用户id;" json:"user_wechat_id"`
	UserId       uint64    `gorm:"Column:user_id;Unique:true;NotNull:true;Comment:用户id;" json:"user_id"`
	UnionId      string    `gorm:"column:unionid;Unique:true;" json:"unionid"`
	Nickname     string    `gorm:"column:nickname" json:"nickname"`
	Sex          uint8     `gorm:"column:sex" json:"sex"`
	Language     string    `gorm:"column:language;Comment:用户的语言" json:"language"`
	City         string    `gorm:"column:city" json:"city"`
	Province     string    `gorm:"column:province" json:"province"`
	Country      string    `gorm:"column:country" json:"country"`
	HeadImgUrl   string    `gorm:"column:headimgurl" json:"headimgurl"`
	Privilege    string    `gorm:"column:privilege;Comment:用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）" json:"privilege"`
	ModelTime    ModelTime `gorm:"Embedded;"`
}

func init() {
}

// TableName 自定义表名 (默认模型名小写)
func (ud *UserCommonWechat) TableName() string {
	return "user_common_wechat"
}

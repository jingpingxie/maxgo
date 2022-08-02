package models

import "time"

//
// @Title:UserCompanyWechat
// @Description: 用户关注企业公众号后拉取的用户信息
// @Author:jingpingxie
// @Date:2022-08-02 11:54:49
//
type UserCompanyWechat struct {
	UserCompanyWechatId uint64    `gorm:"Column:user_wechat_id;PrimaryKey:true;AutoIncrement:false;Comment:微信用户id;" json:"user_wechat_id"`
	CompanyId           uint64    `gorm:"Column:company_id;NotNull:true;Comment:企业id;" json:"company_id"`
	UserId              uint64    `gorm:"Column:user_id;NotNull:true;Comment:用户id;" json:"user_id"`
	Subscribe           uint8     `gorm:"column:subscribe;Comment:用户是否订阅该公众号标识 0 未关注 1 关注" json:"subscribe"`
	Openid              string    `gorm:"column:openid" json:"openid"`
	SubscribeTime       time.Time `gorm:"column:subscribe_time;Comment:用户关注时间" json:"subscribe_time"`
	Remark              string    `gorm:"column:remark;Comment:公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注" json:"remark"`
	GroupId             int64     `gorm:"column:groupid;Comment:用户组id" json:"groupid"`
	ParentUserId        int64     `gorm:"column:parent_user_id" json:"parent_user_id"`
	ParentWechatId      int64     `gorm:"column:parent_wechat_id" json:"parent_wechat_id"`
	ModelTime           ModelTime `gorm:"Embedded;"`
}

func init() {
}

//
// @Title:TableName
// @Description: 自定义表名 (默认模型名小写)
// @Author:jingpingxie
// @Date:2022-08-02 11:54:34
// @Receiver:ucw
// @Return:string
//
func (ucw *UserCompanyWechat) TableName() string {
	return "user_company_wechat"
}

//
// @Title:TableUnique
// @Description: 联合唯一键
// @Author:jingpingxie
// @Date:2022-08-02 11:54:20
// @Receiver:ucw
// @Return:[][]string
//
func (ucw *UserCompanyWechat) TableUnique() [][]string {
	return [][]string{
		{"company_id", "user_id"},
	}
}

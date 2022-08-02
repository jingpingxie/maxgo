package models

import "time"

//
// @Title:Company
// @Description:
// @Author:jingpingxie
// @Date:2022-08-02 11:55:06
//
type Company struct {
	CompanyId             uint64    `gorm:"Column:company_id;PrimaryKey:true;AutoIncrement:false;NotNull:true;Comment:企业ID;" json:"company_id"`
	Name                  string    `gorm:"column:name;Comment:企业全称;" json:"name"`
	ShortName             string    `gorm:"column:short_name;Comment:企业简称;" json:"short_name"`
	LogoUrl               string    `gorm:"column:logo_url;Comment:企业logo" json:"logo_url"`
	MainImageUrl          string    `gorm:"column:main_image_url;Comment:企业主图或者背景图" json:"main_image_url"`
	Zipcode               string    `gorm:"column:zipcode" json:"zipcode"`
	Tel                   string    `gorm:"column:tel" json:"tel"`
	Fax                   string    `gorm:"column:fax" json:"fax"`
	Description           string    `gorm:"column:description;Comment:企业简介;" json:"description"`
	BusinessCountry       int16     `gorm:"column:business_country" json:"business_country"`
	BusinessProvince      int16     `gorm:"column:business_province" json:"business_province"`
	BusinessCity          int16     `gorm:"column:business_city" json:"business_city"`
	BusinessDistrict      int16     `gorm:"column:business_district" json:"business_district"`
	BusinessAddress       string    `gorm:"column:business_address;Comment:经营地址" json:"business_address"`
	RegisterCountry       int64     `gorm:"column:register_country" json:"register_country"`
	RegisterProvince      int64     `gorm:"column:register_province" json:"register_province"`
	RegisterCity          int64     `gorm:"column:register_city" json:"register_city"`
	RegisterDistrict      int64     `gorm:"column:register_district" json:"register_district"`
	RegisterAddress       string    `gorm:"column:register_address;Comment:注册地址" json:"register_address"`
	EmployeesNumber       uint8     `gorm:"column:employees_number" json:"employees_number"`
	Nature                uint8     `gorm:"column:nature;Comment:公司性质 0 未知 1 政府机关/事业单位 2 国营 3 私营 4 中外合资 5 外资 6 其它" json:"nature"`
	CompanyType           int8      `gorm:"column:company_type;Comment:店铺类型 0 未知 1 个人 2 个体经营 3 企业" json:"company_type"`
	LegalPersonUserId     string    `gorm:"column:legal_person_user_id;Comment:法人用户ID" json:"legal_person_user_id"`
	RegisterCapital       string    `gorm:"column:register_capital;Comment:注册资金" json:"register_capital"`
	FoundedTime           time.Time `gorm:"column:founded_time;Comment:成立时间" json:"founded_time"`
	BusinessLicenseId     string    `gorm:"column:business_license_id;Comment:统一社会信用代码" json:"business_license_id"`
	BusinessLicenseImgUrl string    `gorm:"column:business_license_img_url;Comment:营业执照图片" json:"business_license_img_url"`
	BusinessExpiryDate    string    `gorm:"column:business_expiry_date;Comment:营业执照有效期限" json:"business_expiry_date"`
	BankName              string    `gorm:"column:bank_name;Comment:开户行" json:"bank_name"`
	AccountNumber         string    `gorm:"column:account_number;Comment:对公账号" json:"account_number"`
	RegFlag               uint8     `gorm:"column:reg_flag;Comment:申请步骤 0 提交申请 1 发回重改 2 审核通过 3 拒绝注册" json:"reg_flag"`
	RegTime               time.Time `gorm:"column:reg_time;Comment:注册时间" json:"reg_time"`
	ExpireTime            time.Time `gorm:"column:expire_time;Comment:失效时间" json:"expire_time"`
	Remark                string    `gorm:"column:remark;Comment:备注" json:"remark"`
	ParentCompanyId       int64     `gorm:"column:parent_company_id;Comment:母公司ID" json:"parent_company_id"`
	IsClosed              int8      `gorm:"column:is_closed;Comment:店铺是否已经打烊 0 正常营业 1 店铺已经打烊;" json:"is_closed"`
	ModelTime             ModelTime `gorm:"Embedded;"`
}

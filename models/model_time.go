package models

import (
	"gorm.io/gorm"
	"time"
)

//
// @Title:ModelTime
// @Description:
// @Author:jingpingxie
// @Date:2022-08-02 11:55:10
//
type ModelTime struct {
	CreatedAt time.Time      `gorm:"Column:created_at;Type:datetime;Comment:创建时间;Default:current_timestamp;<-:create" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"Column:updated_at;Type:datetime;Comment:更新时间;Default:current_timestamp on update current_timestamp" json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"Column:deleted_at;Type:datetime;NotNull:false;Comment:删除时间;" json:"deleted_at,omitempty"`
}

//
// @File:json
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/6 14:39
//
package xjson

import (
	"encoding/json"
	logs "github.com/sirupsen/logrus"
)

//
// @Title:unmarshalPayload
// @Description: json字符串转结构体
// @Author:jingpingxie
// @Date:2022-08-03 16:31:05
// @Param:jsonData json二进制数据
// @Param:v
// @Return:error
//
func Unmarshal(jsonData []byte, v interface{}) error {
	if err := json.Unmarshal(jsonData, v); err != nil {
		logs.Error("unmarshal json data of %s error: %s", jsonData, err)
		return err
	}
	logs.Info("unmarshal json data:%#v\n", v)
	return nil
}

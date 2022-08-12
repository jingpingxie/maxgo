//
// @File:user_redis
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/11 11:42
//
package user

import (
	"encoding"
	"encoding/json"
)

var _ encoding.BinaryMarshaler = new(UserRedis)
var _ encoding.BinaryUnmarshaler = new(UserRedis)

//
// @Title:UserRedis
// @Description:the user info which save to redis
// @Author:jingpingxie
// @Date:2022-08-11 09:49:06
//
type UserRedis struct {
	//
	//  CID
	//  @Description: client user id，在客户端创建，这个ID从客户端发送给服务端的时候会被加密，因此是安全的，用于识别是否该用户
	//
	CID string `json:"cid"`
	//
	//  TimeDiff
	//  @Description: 用于校准登录客户端和服务器的时差,以判断客户端的真实性,在客户登录时设置这个校准值
	//
	TimeDiff    float64 `json:"time_diff"`
	UserID      uint64  `json:"user_id"`
	Mobile      string  `json:"mobile"`
	UserVisitID uint64  `json:"user_visit_id"`
}

//
// @Title:MarshalBinary
// @Description:
// @Author:jingpingxie
// @Date:2022-08-11 11:45:07
// @Receiver:ur
// @Return:data
// @Return:err
//
func (ur *UserRedis) MarshalBinary() (data []byte, err error) {
	return json.Marshal(ur)
}

//
// @Title:UnmarshalBinary
// @Description:
// @Author:jingpingxie
// @Date:2022-08-11 11:45:10
// @Receiver:ur
// @Param:data
// @Return:error
//
func (ur *UserRedis) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, ur)

}

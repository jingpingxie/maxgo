//
// @File:logout
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/11 14:58
//
package user

type LogoutRequest struct {
	CID string `json:"cid"`
	//
	//  CTIME
	//  @Description: 客户端登录时候上传的客户端的时间
	//
	CTIME int64 `json:"ctime"`
}

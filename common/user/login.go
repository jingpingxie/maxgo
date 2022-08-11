//
// @File:login
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/6 14:46
//
package user

//
// @Title:LoginRequest
// @Description:
// @Author:jingpingxie
// @Date:2022-08-10 18:33:29
//
type LoginRequest struct {
	CID string `json:"cid"`
	//
	//  CTIME
	//  @Description: 客户端登录时候上传的客户端的时间
	//
	CTIME    int64  `json:"ctime"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

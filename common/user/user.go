//
// @File:user
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/11 11:45
//
package user

//
// @Title:UserRequest
// @Description:
// @Author:jingpingxie
// @Date:2022-08-10 18:33:24
//
type UserRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

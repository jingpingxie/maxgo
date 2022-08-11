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
	Account  string `json:"account"`
	Password string `json:"password"`
}

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

//
// @Title:UserResponse
// @Description:
// @Author:jingpingxie
// @Date:2022-08-10 18:33:33
//
type UserResponse struct {
	SID      string `json:"sid"`
	UserName string `json:"user_name"`
	//Token    string `json:"token"`
}

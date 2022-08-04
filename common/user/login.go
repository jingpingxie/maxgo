//
// @File:login
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/3 11:41

//
package user

// LoginRequest defines login request format
type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
type UserRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

// UserResponse defines login response
type UserResponse struct {
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

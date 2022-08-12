//
// @File:user
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/9 13:01
//
package interval_cert

import (
	"maxgo/controllers/base"
	"maxgo/routers"
	"maxgo/services/auth"
	"net/http"
)

func init() {
	routers.Register(&UserController{})
}

//
// @Title:UserController
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:31:06
//
type UserController struct {
	base.LoginBaseController
}

//
// @Title:Post_Logout
// @Description:
// @Author:jingpingxie
// @Date:2022-08-11 14:43:14
// @Receiver:uc
//
func (uc *UserController) Post_Logout() {
	statusCode, err := auth.DoLogout(uc.LoginUserInfo)
	if err != nil {
		uc.Respond(statusCode, -200, err.Error())
		return
	}
	uc.Respond(http.StatusOK, 0, "success to logout")
}

//get请求其实本身HTTP协议并没有限制它的URL大小，但是不同的浏览器对其有不同的大小长度限制
func (uc *UserController) Get_Test() {

}

//本质上来讲， PUT和POST极为相似，都是向服务器发送数据，但它们之间有一个重要区别，PUT通常指定了资源的存放位置，而POST则没有，POST的数据存放位置由服务器自己决定。且put的侧重点在于对于数据的修改操作，但是post侧重于对于数据的增加
func (uc *UserController) Put_Test() {

}

//delete请求用来删除服务器的资源
func (uc *UserController) Delete_Test() {

}

//用于创建、更新资源，于PUT类似，区别在于PATCH代表部分更新
func (uc *UserController) Patch_Test() {

}

//HEAD和GET本质是一样的，区别在于HEAD不含有呈现数据，而仅仅是HTTP头信息。欲判断某个资源是否存在，我们通常使用GET，但这里用HEAD则意义更加明确。
func (uc *UserController) Head_Test() {

}

//options请求属于浏览器的预检请求，查看服务器是否接受请求，预检通过后，浏览器才会去发get，post，put，delete等请求。至于什么情况下浏览器会发预检请求，浏览器会会将请求分为两类，简单请求与非简单请求，非简单请求会产生预检options请求：它用于获取当前URL所支持的方法。若请求成功，则它会在HTTP头中包含一个名为“Allow”的头，值是所支持的方法，如“GET, POST”
func (uc *UserController) Options_Test() {

}

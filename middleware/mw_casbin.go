//
// @File:mw_casbin
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/2 21:17

//
package middleware

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//)
//
////https://zhuanlan.zhihu.com/p/397391798
////https://github.com/it234/goapp
//// CasbinMiddleware casbin中间件
//func CasbinMiddleware(skipper ...SkipperFunc) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		if len(skipper) > 0 && skipper[0](c) {
//			c.Next()
//			return
//		}
//		// 用户ID
//		uid, isExit := c.Get(common.USER_ID_Key)
//		if !isExit {
//			common.ResFailCode(c, "token 无效3", 50008)
//			return
//		}
//		if convert.ToUint64(uid) == common.SUPER_ADMIN_ID {
//			c.Next()
//			return
//		}
//		p := c.Request.URL.Path
//		m := c.Request.Method
//		if b, err := common.CsbinCheckPermission(convert.ToString(uid), p, m); err != nil {
//			common.ResFail(c, "err303"+err.Error())
//			fmt.Println("err303**", err)
//			return
//		} else if !b {
//			common.ResFail(c, "没有访问权限")
//			return
//		}
//		c.Next()
//	}
//}

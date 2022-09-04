//
// @File:router
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/6 13:55
//
package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	logs "github.com/sirupsen/logrus"
	"maxgo/controllers/base"
	"net/http"
	"reflect"
	"regexp"
	"strings"
)

//
// @Title:Route
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 12:42:48
//
type Route struct {
	controller base.IBaseController
	apiPath    string         //url路径
	httpMethod string         //http方法 get post
	methodPath string         //url路径
	Method     reflect.Value  //方法路由
	Args       []reflect.Type //参数类型
}

//路由集合
var Routes = []Route{}

//
// @Title:InitRouter
// @Description: init router
// @Author:jingpingxie
// @Date:2022-08-09 12:42:28
// @Return:*gin.Engine
//
func InitRouter() *gin.Engine {
	//初始化路由
	engine := gin.Default()
	engine.Use(Cors())
	//绑定基本路由，访问路径：/User/List
	Bind(engine)
	// 启动路由 设定端口
	Cors()
	return engine
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma,cert_key")
			//				允许跨域设置																										可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //	跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //	处理请求
	}
}

//
// @Title:Register
// @Description: 注册控制器
// @Author:jingpingxie
// @Date:2022-08-09 12:42:19
// @Param:controller
// @Return:bool
//
func Register(controller base.IBaseController) bool {
	ctrlName := reflect.TypeOf(controller).String()
	fmt.Println("ctrlName=", ctrlName)
	module := ctrlName
	if strings.Contains(ctrlName, ".") {
		module = ctrlName[strings.Index(ctrlName, ".")+1:]
	}
	fmt.Println("module=", module)
	regex, _ := regexp.Compile("Controller$")
	apiModule := regex.ReplaceAllString(module, "")
	apiModule = strings.ToLower(apiModule)

	v := reflect.ValueOf(controller)
	var apiAction string
	//遍历方法
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		action := v.Type().Method(i).Name
		httpMethod := "POST"
		len := len(action)
		if len > 4 && action[:4] == "Get_" {
			httpMethod = "GET"
			apiAction = action[4:len]
		} else if len > 5 && action[:5] == "Post_" {
			httpMethod = "POST"
			apiAction = action[5:len]
		} else if len > 4 && action[:4] == "Put_" {
			httpMethod = "PUT"
			apiAction = action[4:len]
		} else if len > 7 && action[:7] == "Delete_" {
			httpMethod = "DELETE"
			apiAction = action[7:len]
		} else if len > 5 && action[:5] == "Head_" {
			httpMethod = "HEAD"
			apiAction = action[5:len]
		} else if len > 6 && action[:6] == "Patch_" {
			httpMethod = "PATCH"
			apiAction = action[6:len]
		} else if len > 8 && action[:8] == "Options_" {
			httpMethod = "OPTIONS"
			apiAction = action[8:len]
		} else {
			continue
		}
		methodPath := "/" + module + "/" + action
		apiPath := "/" + apiModule + "/" + strings.ToLower(apiAction)
		//遍历参数
		params := make([]reflect.Type, 0, v.NumMethod())
		for j := 0; j < method.Type().NumIn(); j++ {
			params = append(params, method.Type().In(j))
			fmt.Println("params-name=", method.Type().In(j))
		}
		fmt.Println("params=", params)
		fmt.Println("action=", action)
		route := Route{controller: controller, apiPath: apiPath, methodPath: methodPath, Method: method, Args: params, httpMethod: httpMethod}
		Routes = append(Routes, route)
	}
	fmt.Println("Routes=", Routes)
	return true
}

//
// @Title:Bind
// @Description: 绑定基本路由
// @Author:jingpingxie
// @Date:2022-08-09 12:42:04
// @Param:e
//
func Bind(e *gin.Engine) {
	//pathInit()
	//apiv1 := e.Group("/api/v1")
	api := e.Group("/api/v1")
	for _, route := range Routes {
		//e.POST(path, match(path))
		if route.httpMethod == "GET" {
			api.GET(route.apiPath, match(route.methodPath, route))
		} else if route.httpMethod == "POST" {
			api.POST(route.apiPath, match(route.methodPath, route))
		} else if route.httpMethod == "PUT" {
			api.PUT(route.apiPath, match(route.methodPath, route))
		} else if route.httpMethod == "DELETE" {
			api.DELETE(route.apiPath, match(route.methodPath, route))
		} else if route.httpMethod == "HEAD" {
			api.HEAD(route.apiPath, match(route.methodPath, route))
		} else if route.httpMethod == "PATCH" {
			api.PATCH(route.apiPath, match(route.methodPath, route))
		} else if route.httpMethod == "OPTIONS" {
			api.OPTIONS(route.apiPath, match(route.methodPath, route))
		}
	}
}

//
// @Title:match
// @Description: 根据path匹配对应的方法
// @Author:jingpingxie
// @Date:2022-08-09 12:41:50
// @Param:path
// @Param:route
// @Return:gin.HandlerFunc
//
func match(path string, route Route) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fields := strings.Split(path, "/")
		fmt.Println("fields,len(fields)=", fields, len(fields))
		if len(fields) < 3 {
			return
		}

		if len(Routes) > 0 {
			//arguments := make([]reflect.Value, 1)
			//arguments[0] = reflect.ValueOf(ctx) // *gin.Context
			//route.Method.Call(arguments)
			route.controller.SetContext(ctx)
			certController, ok := route.controller.(base.ICertBaseController)
			if ok {
				//decrypt the request data first to encrypt api
				requestMap, err := certController.PreDecrypt()
				if err != nil {
					return
				}
				loginController, ok := route.controller.(base.ILoginBaseController)
				if ok {
					//must login
					userRedis, err := loginController.CheckUser(requestMap)
					if err != nil {
						return
					}
					certKey := ctx.GetHeader("cert_key")
					if certKey == "" {
						logs.Error("without provide cert key")
						respondData := gin.H{
							"code": -200,
							"msg":  "without provide cert key",
						}
						ctx.JSON(http.StatusUnauthorized, respondData)
						return
					}
					err = loginController.RenewExpiredToken(certKey, userRedis.UserID, userRedis.Mobile)
					if err != nil {
						return
					}
				}
			}
			route.Method.Call(nil)
		}
	}
}

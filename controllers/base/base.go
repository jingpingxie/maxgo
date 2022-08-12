//
// @File:base
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/6 14:02
//
package base

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	logs "github.com/sirupsen/logrus"
)

//
// @Title:IBaseController
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 12:33:51
//
type IBaseController interface {
	SetContext(ctx *gin.Context)
	GetContext() (ctx *gin.Context)
}

//
// @Title:BaseController
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 12:33:46
//
type BaseController struct {
	Ctx *gin.Context
}

//
// @Title:SetContext
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 12:33:08
// @Receiver:bc
// @Param:ctx
//
func (bc *BaseController) SetContext(ctx *gin.Context) {
	bc.Ctx = ctx
}

//
// @Title:GetContext
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 12:33:11
// @Receiver:bc
// @Return:ctx
//
func (bc *BaseController) GetContext() (ctx *gin.Context) {
	return bc.Ctx
}

//
// @Title:Respond
// @Description:
// @Author:jingpingxie respond to front end
// @Date:2022-08-09 12:33:19
// @Receiver:bc
// @Param:ctx
// @Param:httpStatus
// @Param:code
// @Param:message
// @Param:data
//
func (bc *BaseController) Respond(httpStatus int, code int, message string, data ...interface{}) {
	respondData := gin.H{
		"code": code,
	}
	if len(data) > 0 {
		respondData["data"] = data[0]
	}
	if len(message) > 0 {
		respondData["msg"] = message
	}

	if logs.GetLevel() >= logs.DebugLevel {
		jsonData, _ := json.Marshal(respondData)
		logs.Debug(string(jsonData))
	}
	bc.Ctx.JSON(httpStatus, respondData)
}

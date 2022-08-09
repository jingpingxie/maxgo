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

type IBaseController interface {
	SetContext(ctx *gin.Context)
	GetContext() (ctx *gin.Context)
}

type BaseController struct {
	Ctx *gin.Context
}

func (bc *BaseController) SetContext(ctx *gin.Context) {
	bc.Ctx = ctx
}
func (bc *BaseController) GetContext() (ctx *gin.Context) {
	return bc.Ctx
}
func (bc *BaseController) Respond(ctx *gin.Context, httpStatus int, code int, message string, data ...interface{}) {
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
	ctx.JSON(httpStatus, respondData)
}

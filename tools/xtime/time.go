//
// @File:time
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/12 12:40
//
package xtime

//
// @Title:GetTimeDiffBetweenSeverAndClient
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:35:19
// @Param:serverTime
// @Param:clientTime
// @Return:float64
//
func GetTimeDiffBetweenSeverAndClient(serverTime float64, clientTime float64) float64 {
	return serverTime - clientTime
}

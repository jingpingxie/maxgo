//
// @File:time
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/12 12:40
//
package xtime

func GetTimeDiffBetweenSeverAndClient(serverTime float64, clientTime float64) float64 {
	return serverTime - clientTime
}

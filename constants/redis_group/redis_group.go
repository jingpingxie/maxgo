//
// @File:redis_group
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/7 17:50
//
package redis_group

type RedisGroup string

const (
	IntervalCert RedisGroup = "interval_cert"
	InstantCert  RedisGroup = "instant_cert"
)

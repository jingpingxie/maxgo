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
	// 定时产生的证书组
	IntervalCert RedisGroup = "interval_cert"
	// 一次性证书组
	DisposableCert RedisGroup = "disposable_cert"
	// 用户组
	User RedisGroup = "user"
)

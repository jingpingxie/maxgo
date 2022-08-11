//
// @File:user
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/9 13:37
//
package user

type User int

const (
	DEFAULT_ACCOUNT_EXPIRE_SECONDS int64   = 3600 //3600s 1hour
	DISPOSABLE_CERT_EXPIRE_SECONDS float64 = 3600 //0.5s
)

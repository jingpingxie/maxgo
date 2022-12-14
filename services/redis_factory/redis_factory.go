//
// @File:redis_factory
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/7 16:24
//
package redis_factory

import (
	"crypto/md5"
	"fmt"
	"github.com/go-redis/redis"
	logs "github.com/sirupsen/logrus"
	user2 "maxgo/common/user"
	"maxgo/constants/redis_group"
	"maxgo/constants/user"
	"maxgo/services/rsa_cert"
	"maxgo/tools/number"
	"strconv"
	"time"
)

var ClientRedis *redis.Client

func init() {
	options := redis.Options{}
	// 新建一个client
	ClientRedis = redis.NewClient(&options)

}

type RedisFactory struct {
}

func String() {
	// 添加string
	ClientRedis.Set("golang_redis", "golang", 0)
	ClientRedis.Set("golang_string", "golang", 0)
	// 获取string
	stringCmd := ClientRedis.Get("golang_redis")
	fmt.Println(stringCmd.String(), stringCmd.Args(), stringCmd.Val())
	// 删除string
	ClientRedis.Del("golang_redis")
}
func Hash() {
	// hash - 添加field
	ClientRedis.HSet("golang_hash", "key_1", "val_1")
	ClientRedis.HSet("golang_hash", "key_2", []string{"key_3", "val_3", "key_4", "val_4"})
	// hash - 获取一个field
	hCmd := ClientRedis.HGet("golang_hash", "user")
	fmt.Println(hCmd.String(), hCmd.Err(), hCmd.Val())
	// hash - 获取长度
	cmd := ClientRedis.HLen("golang_hash")
	fmt.Println(cmd.String(), cmd.Args(), cmd.Val())
	// hash - 获取全部
	cmdAll := ClientRedis.HGetAll("golang_hash")
	fmt.Println(cmdAll.String(), cmdAll.Args(), cmdAll.Val())
	// hash - 获取多个key值
	hmCmd := ClientRedis.HMGet("golang_hash", "key_1", "key_2")
	fmt.Println(hmCmd.String(), hmCmd.Args(), hmCmd.Val())
	// hash - 添加field，没有发现和HSet有什么区别

	//设置元素（多个）
	_field := map[string]interface{}{
		"key1": "value1",
		"key2": "value2",
	}
	ClientRedis.HMSet("golang_hash", _field)
	// 取出元素（多个）
	valHashMany := ClientRedis.HMGet("test-hash", "key1", "key2").Val()
	for k, v := range valHashMany {
		fmt.Println(k)
		fmt.Println(v)
	}

	// hash - 删除field
	ClientRedis.HDel("golang_hash", "key_1", "key_2", "key_3")
}
func List() {
	/**
	 * =============================================
	 * List
	 * =============================================*/
	// 入队列
	ClientRedis.LPush("test-list", "a")
	// 出队列
	valList, errList := ClientRedis.LPop("test-list").Result()
	if errList != nil {
		fmt.Println(errList.Error())
	} else {
		fmt.Println(valList)
	}
}
func Set() {
	// 无序集合 ("马超", "关羽", "赵云") 后面的赵云会覆盖前面的赵云
	ClientRedis.SAdd("golang_set", "马超", "赵云", "关羽", "张飞", "曹植", "司马懿")
	// 从右删除，并返回
	sCmd := ClientRedis.SPop("golang_set")
	fmt.Println(sCmd.String(), sCmd.Args(), sCmd.Val())
	// 指定删除
	ClientRedis.SRem("golang_set", "赵云")
	// 获取集合成语
	sMembers := ClientRedis.SMembers("golang_set")
	fmt.Println(sMembers.String(), sMembers.Args(), sMembers.Val())
	// 返回集合成员数
	cCmd := ClientRedis.SCard("golang_set")
	fmt.Println(cCmd.String(), cCmd.Args(), cCmd.Val())
}
func Zset() {
	/**
	 * =============================================
	 * zSet
	 * =============================================*/
	// 添加元素
	m := redis.Z{
		Score:  5,
		Member: "c",
	}
	m1 := redis.Z{
		Score:  6,
		Member: "d",
	}
	ClientRedis.ZAdd("teat-zSet", m, m1)
	// 取出元素
	valZSet := ClientRedis.ZRange("teat-zSet", 0, 10).String()
	fmt.Println(valZSet)
}

//
// @Title:GetCurrenIntervalRsaCertKey
// @Description:
// @Author:jingpingxie
// @Date:2022-08-13 20:17:13
// @Return:rsaCertKey
//
func GetCurrenIntervalRsaCertKey() (rsaCertKey string) {
	now := time.Now().Unix()
	now -= now % user.DEFAULT_ACCOUNT_EXPIRE_SECONDS
	nowBytes := number.Int64ToBytes(now)
	rsaCertKey = fmt.Sprintf("%x", md5.Sum(nowBytes))
	return rsaCertKey
}

//
// @Title:GenerateIntervalRsaCert
// @Description: get rsa cert data,the cert data should be changed interval to prevent compromised
// @Author:jingpingxie
// @Date:2022-08-07 21:31:19
// @Receiver:rf
//
func GenerateIntervalRsaCert(rsaCertKey string) (rsaCertData *rsa_cert.RsaCert) {
	rsaCertData = &rsa_cert.RsaCert{}
	key := string(redis_group.IntervalCert) + ":" + rsaCertKey
	err := ClientRedis.Get(key).Scan(rsaCertData)
	if err != nil {
		//if without the cert data, generate new cert data
		rsaCertData = &rsa_cert.RsaCert{}
		err = rsaCertData.Generate()
		if err != nil {
			return nil
		}
		err = ClientRedis.SetNX(key, rsaCertData, time.Duration(user.DEFAULT_ACCOUNT_EXPIRE_SECONDS*1e9*2)).Err()
		if err != nil {
			logs.Error("set rsa cert data to redis")
			return nil
		}
	}
	return rsaCertData
}

//
////
//// @Title:ExtendIntervalRsaCertExpireTime
//// @Description:
//// @Author:jingpingxie
//// @Date:2022-08-13 16:34:42
//// @Param:rsaCertKey
//// @Param:rsaCertData
//// @Return:err
////
//func ExtendIntervalRsaCertExpireTime(rsaCertKey string, rsaCertData *rsa_cert.RsaCert) (err error) {
//	key := string(redis_group.IntervalCert) + ":" + rsaCertKey
//	err = ClientRedis.Set(key, rsaCertData, time.Duration(user.DEFAULT_ACCOUNT_EXPIRE_SECONDS*1e9)).Err()
//	if err != nil {
//		logs.Error("extend rsa cert expire time to redis")
//		return err
//	}
//	return nil
//}

//
// @Title:GetIntervalRsaCert
// @Description:
// @Author:jingpingxie
// @Date:2022-08-10 18:33:00
// @Param:uid
// @Return:*auth.RsaCert
// @Return:error
//
func GetIntervalRsaCert(rsaCertKey string) (rsaCertData *rsa_cert.RsaCert, err error) {
	rsaCertData = &rsa_cert.RsaCert{}
	err = ClientRedis.Get(string(redis_group.IntervalCert) + ":" + rsaCertKey).Scan(rsaCertData)
	if err != nil {
		return nil, err
	}
	return rsaCertData, nil
}

//
// @Title:GenerateDisposableRsaCert
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 09:36:20
// @Return:certKey
// @Return:certData
//
func GenerateDisposableRsaCert() (rsaCertKey string, certData *rsa_cert.RsaCert) {
	//generate new cert data
	rsaCertData := &rsa_cert.RsaCert{}
	err := rsaCertData.Generate()
	if err != nil {
		return "", nil
	}
	//the cert data will out of date after DISPOSABLE_CERT_EXPIRE_SECONDS ms
	rsaCertKey = rsaCertData.UID
	key := string(redis_group.DisposableCert) + ":" + rsaCertKey
	err = ClientRedis.SetNX(key, rsaCertData, time.Duration(user.DISPOSABLE_CERT_EXPIRE_SECONDS*1e9)).Err()
	if err != nil {
		logs.Error("set rsa cert data to redis")
		return "", nil
	}
	return rsaCertData.UID, rsaCertData
}

//
// @Title:GetDisposableRsaCert
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 09:36:16
// @Param:uid
// @Return:*auth.RsaCert
// @Return:error
//
func GetDisposableRsaCert(rsaCertKey string) (rsaCertData *rsa_cert.RsaCert, err error) {
	rsaCertData = &rsa_cert.RsaCert{}
	key := string(redis_group.DisposableCert) + ":" + rsaCertKey
	err = ClientRedis.Get(key).Scan(rsaCertData)
	if err != nil {
		logs.Error("maybe it is out of date for getting rsa cert data,")
		return nil, err
	}
	return rsaCertData, err
}

//
// @Title:SaveUser
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:34:37
// @Param:userID
// @Param:userRedis
// @Return:error
//
func SaveUser(userID uint64, userRedis *user2.UserRedis) error {
	err := ClientRedis.Set(string(redis_group.User)+":"+strconv.FormatUint(userID, 10), userRedis, time.Duration(user.DEFAULT_ACCOUNT_EXPIRE_SECONDS*1e9)).Err()
	if err != nil {
		logs.Error("set user info to redis")
		return err
	}
	return nil
}

//
// @Title:GetUser
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:34:35
// @Param:userID
// @Return:*user2.UserRedis
// @Return:error
//
func GetUser(userID uint64) (*user2.UserRedis, error) {
	userRedis := &user2.UserRedis{}
	err := ClientRedis.Get(string(redis_group.User) + ":" + strconv.FormatUint(userID, 10)).Scan(userRedis)
	if err != nil {
		logs.Error("maybe it is out of date for getting user data,")
		return nil, err
	}
	return userRedis, err
}

//
// @Title:ExtendUserExpireTime
// @Description:
// @Author:jingpingxie
// @Date:2022-08-13 16:13:45
// @Param:userID
// @Param:userRedis
// @Return:err
//
func ExtendUserExpireTime(userID uint64, userRedis *user2.UserRedis) (err error) {
	err = ClientRedis.Set(string(redis_group.User)+":"+strconv.FormatUint(userID, 10), userRedis, time.Duration(user.DEFAULT_ACCOUNT_EXPIRE_SECONDS*1e9)).Err()
	if err != nil {
		logs.Error("extend user expire time to redis")
		return err
	}
	return nil
}

//
// @Title:DeleteUser
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:34:31
// @Param:userID
// @Return:error
//
func DeleteUser(userID uint64) error {
	err := ClientRedis.Del(string(redis_group.User) + ":" + strconv.FormatUint(userID, 10)).Err()
	if err != nil {
		logs.Error("failed to delete user info from redis")
		return err
	}
	return nil
}

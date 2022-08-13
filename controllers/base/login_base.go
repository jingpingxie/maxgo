//
// @File:logined_base
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/12 11:03
//
package base

import (
	"errors"
	logs "github.com/sirupsen/logrus"
	"maxgo/common/user"
	"maxgo/services/auth"
	"maxgo/services/redis_factory"
	"maxgo/tools/auth/jwt"
	"maxgo/tools/xtime"
	"net/http"
	"time"
)

//
// @Title:ILoginBaseController
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:30:36
//
type ILoginBaseController interface {
	CheckUser(requestMap map[string]interface{}) (userRedis *user.UserRedis, err error)
	RenewExpiredToken(oldCertKey string, userID uint64, mobile string) error
}

//
// @Title:LoginBaseController
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:30:39
//
type LoginBaseController struct {
	IntervalCertBaseController
	LoginUserInfo *user.UserRedis
}

//
// @Title:CheckUser
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:30:42
// @Receiver:lbc
// @Param:requestMap
// @Return:error
//
func (lbc *LoginBaseController) CheckUser(requestMap map[string]interface{}) (userRedis *user.UserRedis, err error) {
	token := requestMap["token"].(string)
	if len(token) == 0 {
		lbc.Respond(http.StatusUnauthorized, -200, "failed to get token")
		return nil, errors.New("failed to get token")
	}
	jwtPayload, err := jwt.ValidateToken(token)
	if err != nil {
		lbc.Respond(http.StatusUnauthorized, -300, "failed to decrypt token")
		return nil, errors.New("failed to decrypt token")
	}
	userID := jwtPayload.UserID
	logs.Info("user id %s has login", userID)
	userRedis, err = redis_factory.GetUser(userID)
	if err != nil {
		lbc.Respond(http.StatusUnauthorized, -400, "user is not login")
		return nil, errors.New("user is not login")
	}
	if userRedis.CID != requestMap["cid"].(string) {
		lbc.Respond(http.StatusUnauthorized, -500, "different client request,maybe you are hacker")
		return nil, errors.New("different client request,maybe you are hacker")
	}
	//get time diff between server and client
	timeDiff := xtime.GetTimeDiffBetweenSeverAndClient(float64(time.Now().Unix()), requestMap["ctime"].(float64))
	if float64(timeDiff-userRedis.TimeDiff) > 5 {
		lbc.Respond(http.StatusUnauthorized, -600, "client request time invalid,maybe you are hacker")
		return nil, errors.New("client request time invalid,maybe you are hacker")
	}
	//set login user info
	lbc.LoginUserInfo = userRedis
	redis_factory.ExtendUserExpireTime(userID, userRedis)
	return userRedis, nil
}

//
// @Title:UpdateExpiredToken
// @Description:
// @Author:jingpingxie
// @Date:2022-08-13 21:15:13
// @Receiver:lbc
// @Param:oldCertKey
// @Param:userID
// @Param:mobile
// @Return:error
//
func (lbc *LoginBaseController) RenewExpiredToken(oldCertKey string, userID uint64, mobile string) error {

	currentRsaCertKey := redis_factory.GetCurrenIntervalRsaCertKey()
	if currentRsaCertKey == oldCertKey {
		return nil
	}
	//create new rsa cert key
	rsaCertKey, rsaPublicKey, encryptToken, err := auth.GenerateUserToken(userID, mobile)
	if err != nil {
		return err
	}
	lbc.Ctx.Header("Authorization", encryptToken)
	lbc.Ctx.Header("CertKey", rsaCertKey)
	lbc.Ctx.Header("PublicKey", rsaPublicKey)
	return nil
}

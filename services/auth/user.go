//
// @File:user
// @Version:1.0.0
// @Description:
// @Author:jingpingxie
// @Date:2022/8/6 14:49
//
package auth

import (
	"errors"
	"maxgo/common/user"
	user2 "maxgo/constants/user"
	"maxgo/models"
	"maxgo/services"
	"maxgo/services/redis_factory"
	"maxgo/tools/auth/jwt"
	"maxgo/tools/snowflake"
	"maxgo/tools/xtime"
	"net/http"
	"strconv"
	"time"
)

//
// @Title:LoginRequest
// @Description:
// @Author:jingpingxie
// @Date:2022-08-10 18:33:29
//
type LoginRequest struct {
	CID string `json:"cid"`
	//
	//  CTIME
	//  @Description: 客户端登录时候上传的客户端的时间
	//
	CTIME    float64 `json:"ctime"`
	Account  string  `json:"account"`
	Password string  `json:"password"`
}

//
// @Title:LoginResponse
// @Description:
// @Author:jingpingxie
// @Date:2022-08-10 18:33:33
//
type LoginResponse struct {
	UserName string `json:"user_name"`
}

//
// @Title:LoginResult
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:42:23
//
type LoginResult struct {
	UserResponse LoginResponse
	RsaCertKey   string `json:"rsa_key"`
	RsaPublicKey string `json:"rsa_public"`
	Token        string `json:"token"`
}

//
// @Title:RegisterRequest
// @Description:
// @Author:jingpingxie
// @Date:2022-08-10 18:33:24
//
type RegisterRequest struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

//
// @Title:DoLogin
// @Description: do user login
// @Author:jingpingxie
// @Date:2022-08-04 12:47:12
// @Param:lr
// @Return:int
// @Return:*login.LoginResponse
// @Return:error
//
func DoLogin(lr *LoginRequest, clientIP string) (httpStatus int, lrt *LoginResult, err error) {
	// get username and password
	account := lr.Account
	password := lr.Password

	//validate username and password if is empty
	if len(account) == 0 {
		return http.StatusBadRequest, nil, errors.New("error: username is empty")
	}
	if len(password) == 0 {
		return http.StatusBadRequest, nil, errors.New("error: password is empty")
	}

	// check the username if existing
	var dbUser models.User
	services.Db.Raw("select user_id,user_name,salt,password from user where mobile = ? or email=?", account, account).Scan(&dbUser)
	if dbUser.UserID == 0 {
		return http.StatusUnauthorized, nil, errors.New("error: user is not existing")
	}

	// check the password
	if err = checkUserPassword(password, dbUser.Password, dbUser.Salt); err != nil {
		return http.StatusUnauthorized, nil, err
	}

	//generate encrypt jwt token
	//rsaCertKey, tokenString, err := generateToken(dbUser.UserID, dbUser.Mobile)
	rsaCertKey, rsaPublicKey, encryptToken, err := generateToken(dbUser.UserID, dbUser.Mobile)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	//用户登录信息保存到redis
	userRedis := &user.UserRedis{
		CID:      lr.CID,
		TimeDiff: xtime.GetTimeDiffBetweenSeverAndClient(float64(time.Now().Unix()), lr.CTIME),
		UserID:   dbUser.UserID,
		Mobile:   dbUser.Mobile,
	}

	go saveAndUpdateLoginUserInfo(dbUser.UserID, password, clientIP, userRedis)

	lrt = &LoginResult{
		UserResponse: LoginResponse{UserName: dbUser.UserName},
		RsaCertKey:   rsaCertKey,
		RsaPublicKey: rsaPublicKey,
		Token:        encryptToken,
	}
	return http.StatusOK, lrt, nil
}

//
// @Title:DoLogout
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:33:04
// @Param:loginUser
// @Return:httpStatus
// @Return:err
//
func DoLogout(loginUser *user.UserRedis) (httpStatus int, err error) {
	err = redis_factory.DeleteUser(loginUser.UserID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	err = updateLogoutInfo(loginUser.UserVisitID)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

//
// @Title:saveAndUpdateLoginUserInfo
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:32:44
// @Param:userID
// @Param:password
// @Param:clientIP
// @Param:userRedis
// @Return:err
//
func saveAndUpdateLoginUserInfo(userID uint64, password string, clientIP string, userRedis *user.UserRedis) (err error) {
	//generate new salt and password hash
	err = updateUserPassword(userID, password)
	if err != nil {
		return err
	}
	userVisitID, err := createVisitInfo(userID, clientIP)
	if err != nil {
		return err
	}
	//用户登录信息保存到redis
	userRedis.UserVisitID = userVisitID
	err = redis_factory.SaveUser(userID, userRedis)
	if err != nil {
		return err
	}

	return nil
}

//
// @Title:createVisitInfo
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:32:51
// @Param:userID
// @Param:clientIP
// @Return:userVisitID
// @Return:err
//
func createVisitInfo(userID uint64, clientIP string) (userVisitID uint64, err error) {
	//insert user info into db
	userVisitID, err = snowflake.GenerateSnowflakeId()
	if err != nil {
		return 0, err
	}
	duration, _ := time.ParseDuration(strconv.FormatInt(user2.DEFAULT_ACCOUNT_EXPIRE_SECONDS, 10) + "s")
	userVisitInfo := models.UserVisit{
		UserVisitID: userVisitID,
		UserID:      userID,
		LoginTime:   time.Now(),
		LogoutTime:  time.Now().Add(duration),
		VisitIp:     clientIP}
	if err = services.Db.Create(userVisitInfo).Error; err != nil {
		return 0, errors.New("register user," + err.Error())
	}
	return userVisitID, nil
}

//
// @Title:updateLogoutInfo
// @Description:
// @Author:jingpingxie
// @Date:2022-08-12 17:32:39
// @Param:userVisitID
// @Return:err
//
func updateLogoutInfo(userVisitID uint64) (err error) {
	if err = services.Db.Model(&models.UserVisit{}).Where("user_visit_id = ?", userVisitID).Updates(models.UserVisit{LogoutTime: time.Now()}).Error; err != nil {
		return errors.New("update user visit logout time," + err.Error())
	}
	return nil
}

//
// @Title:generateToken
// @Description:
// @Author:jingpingxie
// @Date:2022-08-09 12:39:27
// @Param:userID
// @Param:mobile
// @Return:string
// @Return:error
//
func generateToken(userID uint64, mobile string) (rsaCertKey string, rsaPublicKey string, encryptToken string, err error) {
	tokenString, err := jwt.GenerateToken(userID, mobile, 0)
	if err != nil {
		return "", "", "", errors.New("failed to generate jwt token")
	}
	//generate rsa private key
	rsaCertKey, rsaCert := redis_factory.GenerateIntervalRsaCert()
	if rsaCert == nil {
		return "", "", "", errors.New("failed to get interval rsa cert")
	}
	encryptToken, err = rsaCert.Encrypt(tokenString)
	return rsaCertKey, rsaCert.PublicKey, encryptToken, err
}

//
// @Title:checkUserPassword
// @Description: check whether user password is correct
// @Author:jingpingxie
// @Date:2022-08-04 14:47:02
// @Param:inputPassword
// @Param:dbPassword
// @Param:salt
// @Return:error
//
func checkUserPassword(inputPassword string, dbPassword string, salt string) error {
	var hash string
	var err error
	if hash, err = jwt.GeneratePassHash(inputPassword, salt); err != nil {
		return err
	}
	if hash != dbPassword {
		return errors.New("error: password is error")
	}
	return nil
}

//
// @Title:updateUserPassword
// @Description: update new salt and password for user
// @Author:jingpingxie
// @Date:2022-08-04 14:11:13
// @Param:userID
// @Param:password
// @Return:int
// @Return:*user.RegisterResponse
// @Return:error
//
func updateUserPassword(userID uint64, password string) error {
	//generate new salt and password hash
	var salt string
	var hash string
	var err error
	if salt, hash, err = generateUserPasswordHash(password); err != nil {
		return err
	}
	//update db salt and password hash
	if err = services.Db.Model(&models.User{}).Where("user_id = ?", userID).Updates(models.User{Salt: salt, Password: hash}).Error; err != nil {
		return errors.New("update user salt and password," + err.Error())
	}
	return nil
}

//
// @Title:generateUserPasswordHash
// @Description: generate user password hash
// @Author:jingpingxie
// @Date:2022-08-04 14:49:12
// @Param:password
// @Return:saltRet
// @Return:hashRet
// @Return:errRet
//
func generateUserPasswordHash(password string) (saltRet string, hashRet string, errRet error) {
	//generate new salt and password hash
	var salt string
	var hash string
	var err error
	if salt, err = jwt.GenerateSalt(); err != nil {
		return "", "", err
	}
	if hash, err = jwt.GeneratePassHash(password, salt); err != nil {
		return salt, "", err
	}
	return salt, hash, nil
}

//
// @Title:DoRegisterUser
// @Description: register user
// @Author:jingpingxie
// @Date:2022-08-04 14:49:41
// @Param:rr
// @Return:int
// @Return:*user.LoginResponse
// @Return:error
//
func DoRegister(rr *RegisterRequest) (httpStatus int, lrt *LoginResult, err error) {
	// get mobile and password
	mobile := rr.Mobile
	password := rr.Password

	//validate username and password if is empty
	if len(mobile) == 0 {
		return http.StatusBadRequest, nil, errors.New("error: mobile is empty")
	}
	if len(password) == 0 {
		return http.StatusBadRequest, nil, errors.New("error: password is empty")
	}
	// check the username if existing
	var dbUser models.User
	services.Db.Raw("select user_id,user_name,salt,password from user where mobile = ?", mobile).Scan(&dbUser)
	if dbUser.UserID != 0 {
		// the user is existed
		// check the password
		if err = checkUserPassword(password, dbUser.Password, dbUser.Salt); err != nil {
			return http.StatusUnauthorized, nil, err
		}
		//generate new salt and password hash
		if err = updateUserPassword(dbUser.UserID, password); err != nil {
			return http.StatusInternalServerError, nil, err
		}
		// generate token
		rsaCertKey, rsaPublicKey, encryptToken, err := generateToken(dbUser.UserID, rr.Mobile)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
		lrt = &LoginResult{
			UserResponse: LoginResponse{UserName: dbUser.UserName},
			RsaCertKey:   rsaCertKey,
			RsaPublicKey: rsaPublicKey,
			Token:        encryptToken,
		}
		return http.StatusOK, lrt, nil
	}

	//generate salt and password hash
	var salt string
	var hash string
	if salt, hash, err = generateUserPasswordHash(password); err != nil {
		return http.StatusInternalServerError, nil, err
	}
	//insert user info into db
	userID, err := snowflake.GenerateSnowflakeId()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newUser := models.User{
		UserID:   userID,
		UserName: rr.Mobile,
		Mobile:   rr.Mobile,
		Salt:     salt,
		Password: hash}
	if err = services.Db.Create(newUser).Error; err != nil {
		return http.StatusInternalServerError, nil, errors.New("register user," + err.Error())
	}

	// generate token
	rsaCertKey, rsaPublicKey, encryptToken, err := generateToken(userID, rr.Mobile)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	lrt = &LoginResult{
		UserResponse: LoginResponse{UserName: rr.Mobile},
		RsaCertKey:   rsaCertKey,
		RsaPublicKey: rsaPublicKey,
		Token:        encryptToken,
	}
	return http.StatusOK, lrt, nil
}

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
	"maxgo/models"
	"maxgo/services"
	"maxgo/services/redis_factory"
	"maxgo/tools/auth/jwt"
	"maxgo/tools/snowflake"
	"net/http"
	"time"
)

type User struct {
}

//
// @Title:DoLogin
// @Description: do user login
// @Author:jingpingxie
// @Date:2022-08-04 12:47:12
// @Param:lr
// @Return:int
// @Return:*login.UserResponse
// @Return:error
//
func DoLogin(lr *user.LoginRequest) (int, *user.UserResponse, error) {
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

	var err error

	// check the password
	if err = checkUserPassword(password, dbUser.Password, dbUser.Salt); err != nil {
		return http.StatusUnauthorized, nil, err
	}
	//generate new salt and password hash
	if err = updateUserPassword(dbUser.UserID, password); err != nil {
		return http.StatusInternalServerError, nil, err
	}

	//generate encrypt jwt token
	//uid, tokenString, err := generateToken(dbUser.UserID, dbUser.Mobile)
	uid, _, err := generateToken(dbUser.UserID, dbUser.Mobile)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	//用户登录信息保存到redis
	redis_factory.SaveUser(dbUser.UserID, &user.UserRedis{
		SID:      uid,
		CID:      lr.CID,
		TimeDiff: time.Now().Unix() - lr.CTIME,
		UserID:   dbUser.UserID,
		Mobile:   dbUser.Mobile,
	})
	return http.StatusOK, &user.UserResponse{
		UserName: dbUser.UserName,
		SID:      uid,
	}, nil
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
func generateToken(userID uint64, mobile string) (uid string, encryptToken string, err error) {
	tokenString, err := jwt.GenerateToken(userID, mobile, 0)
	//generate rsa private key
	uid, rsaCert := redis_factory.GenerateIntervalRsaCert()
	encryptToken, err = rsaCert.Encrypt(tokenString)
	return uid, encryptToken, err
}

//
// @Title:VerifyToken
// @Description: verify logined user
// @Author:jingpingxie
// @Date:2022-08-10 17:28:32
// @Param:uid
// @Param:encryptToken
// @Return:uint64
// @Return:error
//
func VerifyToken(uid string, encryptToken string) (uint64, error) {
	rsaCert, err := redis_factory.GetIntervalRsaCert(uid)
	if err != nil {
		return 0, errors.New("there is no suitable cert,maybe expired")
	}
	tokenString, err := rsaCert.Decrypt(encryptToken)
	if err != nil {
		return 0, errors.New("token is invalid")
	}
	jwtPayload, err := jwt.ValidateToken(tokenString)
	if err != nil {
		return 0, err
	}
	return jwtPayload.UserID, nil
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
// @Return:*user.UserResponse
// @Return:error
//
func DoRegister(rr *user.UserRequest) (int, *user.UserResponse, error) {
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
	var err error

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
		uid, _, err := generateToken(dbUser.UserID, rr.Mobile)
		if err != nil {
			return http.StatusInternalServerError, nil, err
		}
		return http.StatusOK, &user.UserResponse{
			UserName: dbUser.UserName,
			SID:      uid,
		}, nil
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
	uid, _, err := generateToken(userID, rr.Mobile)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, &user.UserResponse{
		SID:      uid,
		UserName: rr.Mobile,
	}, nil
}

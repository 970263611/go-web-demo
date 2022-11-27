package token

import (
	"encoding/json"
	"errors"
	_ "fmt"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"project/consts"
	"project/file"
	"project/jwt"
	"project/module"
	"sync"
	"time"
)

var tokenMap sync.Map

func init() {
	tokenMapTempStr := file.Read(consts.TokenPersistence)
	tokenMapTemp := make(map[string]module.TokenUser)
	json.Unmarshal([]byte(tokenMapTempStr), &tokenMapTemp)
	for k, v := range tokenMapTemp {
		tokenMap.Store(k, &v)
	}
	check()
}

func Create(user *module.TokenUser) (string, error) {
	out, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	token := out.String()
	tokenMap.Store(token, user)
	jwtToken, err := jwt.GenerateJwtToken(jwt.Claims{UUID: token})
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func Refresh(tokenJwt string) error {
	claims, err := jwt.ParseJwtToken(tokenJwt)
	if err != nil {
		return errors.New("token expired")
	}
	token := claims.UUID
	user, ok := tokenMap.Load(token)
	if ok {
		timeUnix := time.Now().Unix()
		expirationTime := user.(*module.TokenUser).ExpirationTime
		if timeUnix-expirationTime > file.GetEnvParam().TokenExpired {
			tokenMap.Delete(token)
			return errors.New("token expired")
		}
		user.(*module.TokenUser).ExpirationTime = timeUnix + file.GetEnvParam().TokenExpired
	} else {
		return errors.New("token expired")
	}
	return nil
}

func GetUser(tokenJwt string) (*module.TokenUser, error) {
	claims, err := jwt.ParseJwtToken(tokenJwt)
	if err != nil {
		return nil, errors.New("token expired")
	}
	token := claims.UUID
	user, ok := tokenMap.Load(token)
	if ok {
		timeUnix := time.Now().Unix()
		expirationTime := user.(*module.TokenUser).ExpirationTime
		if timeUnix-expirationTime < file.GetEnvParam().TokenExpired {
			user.(*module.TokenUser).ExpirationTime = timeUnix + file.GetEnvParam().TokenExpired
			return user.(*module.TokenUser), nil
		}
		tokenMap.Delete(token)
	}
	return nil, errors.New("token expired")
}

func RemoveAdmin(token string) error {
	user, ok := tokenMap.Load(token)
	if ok {
		user.(*module.TokenUser).IsAdmin = false
	} else {
		return errors.New("token expired")
	}
	return nil
}

func check() {
	//创建定时任务，精确到秒
	crontab := cron.New(cron.WithSeconds())
	//定义定时器调用的任务函数
	task := func() {
		tokenMapTemp := make(map[string]module.TokenUser)
		tokenMap.Range(func(token, user interface{}) bool {
			tokenUser := user.(*module.TokenUser)
			expirationTime := tokenUser.ExpirationTime
			timeUnix := time.Now().Unix()
			if timeUnix-expirationTime > file.GetEnvParam().TokenExpired {
				tokenMap.Delete(token)
			} else {
				tokenMapTemp[token.(string)] = *tokenUser
			}
			return true
		})
		tokenMapStr, err := json.Marshal(tokenMapTemp)
		if err == nil {
			file.Write(consts.TokenPersistence, string(tokenMapStr))
		}
	}
	//定时任务
	spec := file.GetEnvParam().TokenRefreshCron
	// 添加定时任务,
	crontab.AddFunc(spec, task)
	// 启动定时器
	crontab.Start()
}

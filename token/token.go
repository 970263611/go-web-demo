package token

import (
	"errors"
	_ "fmt"
	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
	"log"
	"project/module"
	"sync"
	"time"
)

var tokenMap sync.Map
var expiration int64 = 30 * 60 * 1000

func init() {
	check()
}

func Create(user *module.TokenUser) string {
	out, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	token := out.String()
	tokenMap.Store(token, user)
	return token
}

func Refresh(token string) error {
	user, ok := tokenMap.Load(token)
	if ok {
		timeUnix := time.Now().Unix()
		expirationTime := user.(*module.TokenUser).ExpirationTime
		if timeUnix-expirationTime > expiration {
			tokenMap.Delete(token)
			return errors.New("token expired")
		}
		user.(*module.TokenUser).ExpirationTime = timeUnix + expiration
	} else {
		return errors.New("token expired")
	}
	return nil
}

func GetUser(token string) (*module.TokenUser, error) {
	user, ok := tokenMap.Load(token)
	if ok {
		timeUnix := time.Now().Unix()
		expirationTime := user.(*module.TokenUser).ExpirationTime
		if timeUnix-expirationTime < expiration {
			user.(*module.TokenUser).ExpirationTime = timeUnix + expiration
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
		tokenMap.Range(func(token, user interface{}) bool {
			expirationTime := user.(*module.TokenUser).ExpirationTime
			timeUnix := time.Now().Unix()
			if timeUnix-expirationTime > expiration {
				tokenMap.Delete(token)
			}
			return true
		})
	}
	//定时任务
	spec := "*/30 * * * * ?" //cron表达式，每30秒一次
	// 添加定时任务,
	crontab.AddFunc(spec, task)
	// 启动定时器
	crontab.Start()
}

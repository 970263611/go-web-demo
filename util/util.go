package util

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project/module"
	"project/token"
)

func GetBodyMap(c *gin.Context) (map[string]any, error) {
	json := make(map[string]any)
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, errors.New("http body cannot find value")
	}
	return json, nil
}

func GetLoginUser(c *gin.Context) (*module.TokenUser, error) {
	loginUser, err := token.GetUser(c.GetHeader("token"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return nil, errors.New("login user cannot find")
	}
	return loginUser, nil
}

func CheckAdmin(c *gin.Context) bool {
	loginUser, err := GetLoginUser(c)
	if err == nil {
		if loginUser.IsAdmin {
			return true
		}
	}
	return false
}

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

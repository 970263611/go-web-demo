package util

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"project/module"
	"project/result"
	"project/token"
)

func GetEnvDefault(key, defVal string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defVal
}

func GetBodyMap(c *gin.Context) map[string]any {
	json := make(map[string]any)
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic(nil)
	}
	return json
}

func GetLoginUser(c *gin.Context) *module.TokenUser {
	loginUser, err := token.GetUser(c.GetHeader("token"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		panic(nil)
	}
	return loginUser
}

func CheckAdmin(c *gin.Context) {
	loginUser := GetLoginUser(c)
	if !loginUser.IsAdmin {
		c.JSON(http.StatusForbidden, result.Join(false, "you do not have permission"))
		panic(nil)
	}
}

func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

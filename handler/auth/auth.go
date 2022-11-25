package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"project/datasource"
	"project/ent/user"
	"project/module"
	"project/result"
	"project/token"
	"project/util"
	"strings"
	"time"
)

func Login(c *gin.Context) {
	var flag = false
	json := util.GetBodyMap(c)
	username := json["username"].(string)
	password := json["password"].(string)
	resultUser, err := datasource.Client().User.Query().Where(user.UserCode(username), user.Password(util.MD5(password))).
		//func(selector *sql.Selector) {
		//	selector.Where(sql.EQ(user.FieldDefaultPassword, user.FieldPassword)).Or().Where(sql.EQ(user.FieldPassword, password))
		//},
		Only(context.Background())
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"status": "authentication failed"})
	} else {
		flag = true
		res := make(map[string]any, 1)
		var tokenUser = module.TokenUser{
			UserCode:       resultUser.UserCode,
			Username:       resultUser.Username,
			IsAdmin:        strings.Compare(resultUser.IsAdmin.String(), "0") == 0,
			ExpirationTime: time.Now().Unix(),
		}
		res["token"] = token.Create(&tokenUser)
		res["needResetPassword"] = strings.Compare(resultUser.DefaultPassword, resultUser.Password) == 0
		c.JSON(http.StatusOK, result.Join(flag, res))
	}
}

func Reset(c *gin.Context) {
	loginUser := util.GetLoginUser(c)
	password := util.GetBodyMap(c)["password"].(string)
	datasource.Client().User.Update().Where(user.UserCode(loginUser.UserCode)).SetPassword(util.MD5(password)).Save(context.Background())
	c.JSON(http.StatusOK, result.Join(true, nil))
}

func Info(c *gin.Context) {
	var flag = false
	username := c.Query("username")
	password := c.Query("password")
	if strings.Compare(username, "admin") == 0 && strings.Compare(password, "123456") == 0 {
		flag = true
	}
	c.JSON(http.StatusOK, result.Join(flag, nil))
}

func Search(c *gin.Context) {
	util.CheckAdmin(c)
	userCode := util.GetBodyMap(c)["id"].(string)
	user, _ := datasource.Client().User.Query().Where(user.UserCode(userCode)).Only(context.Background())
	user.Password = ""
	user.DefaultPassword = ""
	c.JSON(http.StatusOK, result.Join(true, user))
}

func Add(c *gin.Context) {
	util.CheckAdmin(c)
	bodyMap := util.GetBodyMap(c)
	userCode := bodyMap["username"].(string)
	username := bodyMap["name"].(string)
	defaultPassword := bodyMap["password"].(string)
	authList := bodyMap["authList"].(map[string]any)
	datasource.Client().User.Create().
		SetUserCode(userCode).
		SetUsername(username).
		SetPassword(util.MD5(defaultPassword)).
		SetDefaultPassword(util.MD5(defaultPassword)).
		SetAuthList(authList).
		Save(context.Background())
	c.JSON(http.StatusOK, result.Join(true, nil))
}

func Update(c *gin.Context) {
	util.CheckAdmin(c)
	bodyMap := util.GetBodyMap(c)
	userCode := bodyMap["username"].(string)
	authList := bodyMap["authList"].(map[string]any)
	datasource.Client().User.Update().Where(user.UserCode(userCode)).SetAuthList(authList).Save(context.Background())
	c.JSON(http.StatusOK, result.Join(true, nil))
}

func Delete(c *gin.Context) {
	util.CheckAdmin(c)
	userCode := util.GetBodyMap(c)["username"].(string)
	datasource.Client().User.Delete().Where(user.UserCode(userCode)).Exec(context.Background())
	c.JSON(http.StatusOK, result.Join(true, nil))
}

func AdminChange(c *gin.Context) {
	util.CheckAdmin(c)
	userCodeWillAdmin := util.GetBodyMap(c)["username"].(string)
	loginUser := util.GetLoginUser(c)
	//开启事务
	tx, err := datasource.Client().Tx(context.Background())
	if err != nil {
		panic("start tx error")
	}
	_, err1 := tx.User.Update().Where(user.UserCode(loginUser.UserCode)).SetIsAdmin("1").Save(context.Background())
	_, err2 := tx.User.Update().Where(user.UserCode(userCodeWillAdmin)).SetIsAdmin("0").Save(context.Background())
	if err1 != nil || err2 != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	token.RemoveAdmin(c.GetHeader("token"))
	c.JSON(http.StatusOK, result.Join(true, nil))
}

func ResetPassword(c *gin.Context) {
	util.CheckAdmin(c)
	userCode := util.GetBodyMap(c)["username"].(string)
	userTemp, err := datasource.Client().User.Query().Where(user.UserCode(userCode)).Only(context.Background())
	if err != nil {
		datasource.Client().User.Update().Where(user.UserCode(userTemp.UserCode)).SetPassword(userTemp.DefaultPassword).Save(context.Background())
		c.JSON(http.StatusOK, result.Join(true, nil))
	} else {
		c.JSON(http.StatusOK, result.Join(false, "cannot find match user"))
	}
}

package handler

import (
	"github.com/gin-gonic/gin"
	"project/handler/auth"
	"project/handler/filter"
	"project/handler/health"
)

func Do(router *gin.Engine) {
	router.Use(Cors())
	//健康检查
	router.GET(
		"/health",
		health.Check,
	)
	//登录
	router.POST(
		"/login",
		auth.Login,
	)
	//添加token验证
	router.Use(filter.AuthFilter())
	//修改用户密码
	router.POST(
		"/resetPass",
		auth.Reset,
	)
	//获取用户信息
	router.POST(
		"/getUserInfo",
		auth.Info,
	)
	userGroup := router.Group("/user")
	//搜索用户
	userGroup.POST(
		"/search",
		auth.Search,
	)
	//新增用户
	userGroup.POST(
		"/add",
		auth.Add,
	)
	//更新用户
	userGroup.POST(
		"/update",
		auth.Update,
	)
	//删除用户
	userGroup.POST(
		"/delete",
		auth.Delete,
	)
	//移交管理员
	router.POST(
		"/change/admin",
		auth.AdminChange,
	)
	//重置密码
	router.POST(
		"/reset/password",
		auth.ResetPassword,
	)
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

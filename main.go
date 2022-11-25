package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project/handler"
	"time"
)

func main() {
	//为日志加上颜色
	gin.ForceConsoleColor()
	//设置日志文件
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	//设置日志格式
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	//定义URI处理器
	handler.Do(router)
	//启动服务
	router.Run(":8099")
}

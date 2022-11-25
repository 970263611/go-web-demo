package result

import "github.com/gin-gonic/gin"

func Join(success bool, object any) gin.H {
	code := 200
	if success != true {
		code = 500
	}
	return gin.H{
		"code": code,
		"data": object,
	}
}

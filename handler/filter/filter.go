package filter

import (
	"github.com/gin-gonic/gin"
	"net/http"
	tokenUtil "project/token"
)

// AuthFilter token认证
func AuthFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		err := tokenUtil.Refresh(token)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		} else {
			c.Next()
		}
	}
}

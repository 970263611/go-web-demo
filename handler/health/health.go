package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/result"
)

func Check(c *gin.Context) {
	res := make(map[string]any, 1)
	res["health"] = true
	c.JSON(http.StatusOK, result.Join(
		true,
		res,
	))
}

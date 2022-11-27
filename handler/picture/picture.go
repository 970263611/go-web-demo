package picture

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/datasource"
	"project/result"
)

func Score(c *gin.Context) {
	res, err := datasource.NativeSqlQuery("select * from users where username = $1;", "dahua")
	if err == nil {
		c.JSON(http.StatusOK, result.Join(true, res))
	} else {
		c.JSON(http.StatusInternalServerError, result.Join(false, err.Error()))
	}
}

func Risk(c *gin.Context) {}

func Select(c *gin.Context) {}

func UpAndDown(c *gin.Context) {}

func PersonalEchart(c *gin.Context) {}

func Steps(c *gin.Context) {}

func Cloud(c *gin.Context) {}

func History(c *gin.Context) {}

func UserDetails(c *gin.Context) {}

func Depart(c *gin.Context) {}

func RiskReason(c *gin.Context) {}

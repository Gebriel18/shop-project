package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type ApiController struct{}

type Myclaims struct {
	Uid  int
	Name string
	jwt.RegisteredClaims
}

var expireTime = time.Now().Add(24 * time.Hour)

func (con ApiController) Index(c *gin.Context) {
	c.String(200, "我是一个api接口")
}
func (con ApiController) Userlist(c *gin.Context) {
	c.String(200, "我是一个api接口-Userlist")
}
func (con ApiController) Plist(c *gin.Context) {
	//c.String(200, "我是一个api接口-Plist")
	myClaimsObj := Myclaims{
		23,
		"zhangzhiti",
		jwt.RegisteredClaims{

			Issuer:    "zhiti",
			ExpiresAt: jwt.NewNumericDate(expireTime),
		},
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodES256, myClaimsObj)

	tokenStr, err := tokenObj.SigningString()

	if err != nil {
		c.JSON(200, gin.H{
			"message": "生成token失败重试",
			"success": false,
		})
	}

	c.JSON(200, gin.H{
		"message": "获取token成功",
		"token":   tokenStr,
		"success": true,
	})
}
func (con ApiController) AddressList(c *gin.Context) {
	tokenData := c.Request.Header.Get("Authorization")

	fmt.Println(tokenData)

	c.String(200, tokenData)
}

package middlewares

import (
	"ginshop42/models"
	"github.com/gin-gonic/gin"
)

func InitUserAuthMiddleware(c *gin.Context) {

	user := models.User{}
	isLogin := models.Cookie.Get(c, "userinfo", &user)

	if !isLogin || len(user.Phone) != 11 {
		c.Redirect(302, "/pass/login")
		return
	}
}

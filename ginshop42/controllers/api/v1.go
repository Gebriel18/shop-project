package api

import (
	"ginshop42/models"
	"github.com/gin-gonic/gin"
)

type V1Controller struct{}

func (con V1Controller) Index(c *gin.Context) {
	c.String(200, "我是一个api接口")
}
func (con V1Controller) NavList(c *gin.Context) {
	var navList []models.Nav
	models.DB.Find(&navList)
	c.JSON(200, gin.H{
		"navList": navList,
	})
	//c.String(200, "我是一个api接口-Userlist")
}
func (con V1Controller) Plist(c *gin.Context) {
	c.String(200, "我是一个api接口-Plist")
}

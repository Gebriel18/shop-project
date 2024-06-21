package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

// 定义结构体  缓存结构体 私有
type ginCookie struct{}

// 写入数据的方法
func (cookie ginCookie) Set(c *gin.Context, key string, value interface{}) {

	bytes, _ := json.Marshal(value)

	desKey := []byte("itying.c") //必须8位

	encData, _ := DesEncrypt(bytes, desKey)

	c.SetCookie(key, string(encData), 3600, "/", "", false, true)
}

// 获取数据的方法
func (cookie ginCookie) Get(c *gin.Context, key string, obj interface{}) bool {

	valueStr, err1 := c.Cookie(key)

	if err1 == nil && valueStr != "" && valueStr != "[]" {

		desKey := []byte("itying.c") //必须8位
		decData, err := DesDecrypt([]byte(valueStr), desKey)
		if err != nil {
			return false
		} else {
			err2 := json.Unmarshal([]byte(decData), obj)
			return err2 == nil
		}
	}
	return false
}

func (cookie ginCookie) Remove(c *gin.Context, key string) bool {
	c.SetCookie(key, "", -1, "/", "", false, true)
	return true
}

// 实例化结构体
var Cookie = &ginCookie{}

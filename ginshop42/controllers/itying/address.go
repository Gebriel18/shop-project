package itying

import (
	"ginshop42/models"
	"github.com/gin-gonic/gin"
)

type AddressController struct {
	BaseController
}

func (con AddressController) AddAddress(c *gin.Context) {
	user := models.User{}
	models.Cookie.Get(c, "userinfo", &user)
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	address := c.PostForm("address")
	// 2、判断收货地址的数量
	var addressNum int64
	models.DB.Table("address").Where("uid = ?", user.Id).Count(&addressNum)
	if addressNum > 10 {
		c.JSON(200, gin.H{
			"success": false,
			"message": "收货地址的数量超过了限制，请编辑以前的收货地址",
		})
		return
	}

	// 3、更新当前用户的所有收货地址的默认收货地址状态为0
	models.DB.Table("address").Where("uid = ?", user.Id).Updates(map[string]interface{}{"default_address": 0})

	// 4、增加当前收货地址，让默认收货地址状态是1
	addressResult := models.Address{
		Uid:            user.Id,
		Name:           name,
		Phone:          phone,
		Address:        address,
		DefaultAddress: 1,
	}
	models.DB.Create(&addressResult)
	// 5、返回当前用户的所有收货地址返回

	addressList := []models.Address{}
	models.DB.Where("uid = ?", user.Id).Order("id desc").Find(&addressList)

	c.JSON(200, gin.H{
		"success":     true,
		"addressList": addressList,
	})
}

// 获取一个收货地址  返回指定收货地址id的收货地址
func (con AddressController) GetOneAddressList(c *gin.Context) {
	addressId, err := models.Int(c.Query("addressId"))
	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误1",
		})
		return
	}
	user := models.User{}
	models.Cookie.Get(c, "userinfo", &user)

	var addressList []models.Address

	models.DB.Where("id=? AND uid = ?", addressId, user.Id).Find(&addressList)

	if len(addressList) > 0 {
		c.JSON(200, gin.H{
			"success": true,
			"result":  addressList[0],
		})
	} else {
		c.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误2",
		})
		return
	}

	//c.String(200, " 获取一个收货地址")
}

// 编辑收货地址
func (con AddressController) EditAddress(c *gin.Context) {
	/*
	   1、获取表单增加的数据

	   2、更新当前用户的所有收货地址的默认收货地址状态为0

	   3、修改当前收货地址，让默认收货地址状态是1

	    4、查询当前用户的所有收货地址并返回

	*/
	//c.String(200, " 编辑收货地址")
	user := models.User{}
	models.Cookie.Get(c, "userinfo", &user)
	id, err := models.Int(c.PostForm("id"))
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	address := c.PostForm("address")

	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误",
		})
		return
	}

	// 2、更新当前用户的所有收货地址的默认收货地址状态为0
	models.DB.Table("address").Where("uid = ?", user.Id).Updates(map[string]interface{}{"default_address": 0})

	// 3、修改当前收货地址，让默认收货地址状态是1
	editAddress := models.Address{Id: id}
	models.DB.Find(&editAddress)
	editAddress.Name = name
	editAddress.Phone = phone
	editAddress.Address = address
	editAddress.DefaultAddress = 1
	models.DB.Save(&editAddress)

	// 4、返回当前用户的所有收货地址返回

	var addressList []models.Address
	models.DB.Where("uid = ?", user.Id).Order("id desc").Find(&addressList)
	c.JSON(200, gin.H{
		"success": true,
		"result":  addressList,
	})
}

// 修改默认的收货地址
func (con AddressController) ChangeDefaultAddress(c *gin.Context) {
	/*
	   1、获取当前用户收货地址id 以及用户id
	   2、更新当前用户的所有收货地址的默认收货地址状态为0
	   3、更新当前收货地址的默认收货地址状态为1
	*/
	//c.String(200, " 修改默认的收货地址")

	user := models.User{}
	models.Cookie.Get(c, "userinfo", &user)
	addressId, err := models.Int(c.Query("addressId"))
	if err != nil {
		c.JSON(200, gin.H{
			"success": false,
			"message": "传入参数错误",
		})
		return
	}

	models.DB.Table("address").Where("uid = ?", user.Id).Updates(map[string]interface{}{"default_address": 0})

	models.DB.Table("address").Where("uid = ? AND id = ?", user.Id, addressId).Updates(map[string]interface{}{"default_address": 1})

	c.JSON(200, gin.H{
		"success": true,
		"message": "修改数据成功",
	})
}

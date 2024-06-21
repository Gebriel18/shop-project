package itying

import (
	"ginshop51/models"
	"github.com/gin-gonic/gin"
)

type BuyController struct {
	BaseController
}

func (con BuyController) Checkout(c *gin.Context) {
	var userInfo models.User
	if !models.Cookie.Get(c, "userinfo", &userInfo) {
		c.Redirect(302, "/")
	}

	cartList := []models.Cart{}
	models.Cookie.Get(c, "cartList", &cartList)

	var orderList []models.Cart
	var allPrice float64
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
		}
	}

	con.Render(c, "itying/buy/checkout.html", gin.H{
		//"success": true,
		//"msg":     "ok",

		"orderList": orderList,
		"allPrice":  allPrice,
	})

}

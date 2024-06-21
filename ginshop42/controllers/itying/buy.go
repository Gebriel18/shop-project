package itying

import (
	"fmt"
	"ginshop42/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type BuyController struct {
	BaseController
}

func (con BuyController) CheckOut(c *gin.Context) {

	var cartList []models.Cart
	models.Cookie.Get(c, "cartList", &cartList)

	var orderList []models.Cart

	var allPrice float64
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
		}
	}

	var user models.User
	models.Cookie.Get(c, "userinfo", &user)

	var addressList []models.Address
	models.DB.Where("uid=?", user.Id).Find(&addressList)

	orderSign := models.Md5(models.GetRandomNum())

	session := sessions.Default(c)
	session.Set("orderSign", orderSign)
	session.Save()

	if len(orderList) == 0 {
		c.Redirect(302, "/")
		return
	}

	con.Render(c, "itying/buy/checkout.html", gin.H{
		"orderList":   orderList,
		"allPrice":    allPrice,
		"addressList": addressList,
		"orderSign":   orderSign,
	})

}

func (con BuyController) DoCheckOut(c *gin.Context) {

	orderSignClient := c.PostForm("orderSign")
	session := sessions.Default(c)
	orderSignSession := session.Get("orderSign")
	orderSignServer, ok := orderSignSession.(string)
	if !ok {
		c.Redirect(302, "/")
		return
	}
	if orderSignClient != orderSignServer {
		c.Redirect(302, "/")
		return
	}
	session.Delete("orderSign")
	session.Save()
	user := models.User{}
	models.Cookie.Get(c, "userinfo", &user)

	addressResult := []models.Address{}
	models.DB.Where("uid = ? AND default_address=1", user.Id).Find(&addressResult)

	if len(addressResult) == 0 {
		c.Redirect(302, "/buy/checkout")
		return
	}

	// 2、获取购买商品的信息
	cartList := []models.Cart{}
	models.Cookie.Get(c, "cartList", &cartList)
	orderList := []models.Cart{}
	var allPrice float64
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
		}
	}
	// 3、把订单信息放在订单表，把商品信息放在商品表
	order := models.Order{
		OrderId:     models.GetOrderId(),
		Uid:         user.Id,
		AllPrice:    allPrice,
		Phone:       addressResult[0].Phone,
		Name:        addressResult[0].Name,
		Address:     addressResult[0].Address,
		PayStatus:   0,
		PayType:     0,
		OrderStatus: 0,
		AddTime:     int(models.GetUnix()),
	}

	err := models.DB.Create(&order).Error
	//增加数据成功以后可以通过  order.Id
	if err == nil {
		// 把商品信息放在商品对应的订单表
		for i := 0; i < len(orderList); i++ {
			orderItem := models.OrderItem{
				OrderId:      order.Id,
				Uid:          user.Id,
				ProductTitle: orderList[i].Title,
				ProductId:    orderList[i].Id,
				ProductImg:   orderList[i].GoodsImg,
				ProductPrice: orderList[i].Price,
				ProductNum:   orderList[i].Num,
				GoodsVersion: orderList[i].GoodsVersion,
				GoodsColor:   orderList[i].GoodsColor,
			}
			models.DB.Create(&orderItem)
		}
	}

	// 4、删除购物车里面的选中数据
	var noSelectCartList []models.Cart
	for i := 0; i < len(cartList); i++ {
		if !cartList[i].Checked {
			noSelectCartList = append(noSelectCartList, cartList[i])
		}
	}
	models.Cookie.Set(c, "cartList", noSelectCartList)

	c.Redirect(302, "/buy/pay?orderId="+models.String(order.Id))
}

func (con BuyController) Pay(c *gin.Context) {
	orderId, err := models.Int(c.Query("orderId"))
	if err != nil {
		c.Redirect(302, "/")
		return
	}
	user := models.User{}
	models.Cookie.Get(c, "userinfo", &user)

	order := models.Order{}
	models.DB.Where("id= ?", orderId).Find(&order)

	if order.Uid != user.Id {
		c.Redirect(302, "/")
		return
	}

	var orderItem []models.OrderItem
	models.DB.Where("order_id=?", orderId).Find(&orderItem)
	var allPrice float64
	allPrice = 0
	for _, v := range orderItem {
		allPrice += v.ProductPrice * float64(v.ProductNum)
	}
	fmt.Println(orderItem)
	con.Render(c, "itying/buy/pay.html", gin.H{
		"order":      order,
		"orderItems": orderItem,
		"allPrice":   allPrice,
	})
}

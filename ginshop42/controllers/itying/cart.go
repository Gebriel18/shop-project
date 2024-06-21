package itying

import (
	"ginshop42/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CartController struct {
	BaseController
}

func (con CartController) Get(c *gin.Context) {
	var cartList []models.Cart
	models.Cookie.Get(c, "cartList", &cartList)
	//c.JSON(200, gin.H{
	//	"cartList": cartList,
	//})
	var allPrice float64
	var allNum int
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			allNum += cartList[i].Num
		}
	}
	var tpl = "itying/cart/cart.html"

	con.Render(c, tpl, gin.H{
		"cartList": cartList,
		"allPrice": allPrice,
	})
}
func (con CartController) AddCart(c *gin.Context) {
	colorId, _ := models.Int(c.Query("color_id"))
	goodsId, err := models.Int(c.Query("goods_id"))
	if err != nil {
		c.Redirect(302, "/")
	}

	goods := models.Goods{}
	goodsColor := models.GoodsColor{}

	models.DB.Where("id=?", goodsId).Find(&goods)
	models.DB.Where("id=?", colorId).Find(&goodsColor)

	currentData := models.Cart{
		Id:           goodsId,
		Title:        goods.Title,
		Price:        goods.Price,
		GoodsVersion: goods.GoodsVersion,
		Num:          1,
		GoodsColor:   goodsColor.ColorName,
		GoodsImg:     goods.GoodsImg,
		GoodsGift:    goods.GoodsGift,
		GoodsAttr:    "",
		Checked:      true,
	}
	cartList := []models.Cart{}

	models.Cookie.Get(c, "cartList", &cartList)

	if len(cartList) > 0 {
		if models.HasCartData(cartList, currentData) {
			for i := 0; i < len(cartList); i++ {
				if cartList[i].Id == currentData.Id && cartList[i].GoodsColor == currentData.GoodsColor && cartList[i].GoodsAttr == currentData.GoodsAttr {
					cartList[i].Num = cartList[i].Num + 1
				}
			}
		} else {
			cartList = append(cartList, currentData)

		}
		models.Cookie.Set(c, "cartList", cartList)
	} else {
		cartList = append(cartList, currentData)
		models.Cookie.Set(c, "cartList", cartList)
	}

	//var ca []models.Cart
	//models.Cookie.Get(c, "cartList", &ca)
	//c.JSON(200, gin.H{
	//	"colorId": colorId,
	//	"goodsId": goodsId,
	//})
	c.Redirect(302, "/cart/successTip?goods_id="+models.String(goodsId))

}

func (con CartController) AddCartSuccess(c *gin.Context) {
	goodsId, err := models.Int(c.Query("goods_id"))
	if err != nil {
		c.Redirect(302, "/")
	}

	goods := models.Goods{}
	models.DB.Where("id=?", goodsId).Find(&goods)

	var tpl = "itying/cart/addcart_success.html"
	con.Render(c, tpl, gin.H{
		"goods": goods,
	})

}

func (con CartController) DecCart(c *gin.Context) {
	goodsId, err := models.Int(c.Query("goods_id"))
	goodsColor := c.Query("goods_color")
	goodsAttr := ""
	var allPrice float64
	var currentPrice float64
	var num int
	var response gin.H

	if err != nil {

		response = gin.H{
			"success": false,
			"message": "传入参数错误1",
		}

	} else {
		var cartList []models.Cart
		models.Cookie.Get(c, "cartList", &cartList)
		if len(cartList) > 0 {

			for i := 0; i < len(cartList); i++ {
				if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == goodsAttr {
					if cartList[i].Num > 1 {
						cartList[i].Num = cartList[i].Num - 1
					}

					currentPrice = float64(cartList[i].Num) * cartList[i].Price
					num = cartList[i].Num
				}
				if cartList[i].Checked {
					allPrice += cartList[i].Price * float64(cartList[i].Num)
				}
			}
			models.Cookie.Set(c, "cartList", cartList)
			response = gin.H{
				"success":      true,
				"message":      "更新数据成功",
				"allPrice":     allPrice,
				"num":          num,
				"currentPrice": currentPrice,
			}
		} else {
			response = gin.H{
				"success": false,
				"message": "传入参数错误2",
			}
		}

	}

	c.JSON(200, response)
}
func (con CartController) IncCart(c *gin.Context) {
	goodsId, err := models.Int(c.Query("goods_id"))
	goodsColor := c.Query("goods_color")
	goodsAttr := ""
	var allPrice float64
	var currentPrice float64
	var num int
	var response gin.H

	if err != nil {

		response = gin.H{
			"success": false,
			"message": "传入参数错误1",
		}

	} else {
		var cartList []models.Cart
		models.Cookie.Get(c, "cartList", &cartList)
		if len(cartList) > 0 {

			for i := 0; i < len(cartList); i++ {
				if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == goodsAttr {
					cartList[i].Num = cartList[i].Num + 1
					currentPrice = float64(cartList[i].Num) * cartList[i].Price
					num = cartList[i].Num
				}
				if cartList[i].Checked {
					allPrice += cartList[i].Price * float64(cartList[i].Num)
				}
			}
			models.Cookie.Set(c, "cartList", cartList)
			response = gin.H{
				"success":      true,
				"message":      "更新数据成功",
				"allPrice":     allPrice,
				"num":          num,
				"currentPrice": currentPrice,
			}
		} else {
			response = gin.H{
				"success": false,
				"message": "传入参数错误2",
			}
		}

	}

	c.JSON(200, response)
	//goods := models.Goods{}
	//goodsColor := models.GoodsColor{}
	//models.DB.Where("id=?", goodsId).Find(&goods)
	//models.DB.Where("id=?", colorId)

}

func (con CartController) ChangeOneCart(c *gin.Context) {
	goodsId, err := models.Int(c.Query("goods_id"))
	goodsColor := c.Query("goods_color")
	goodsAttr := ""
	var allPrice float64

	var response gin.H

	if err != nil {

		response = gin.H{
			"success": false,
			"message": "传入参数错误1",
		}

	} else {
		var cartList []models.Cart
		models.Cookie.Get(c, "cartList", &cartList)
		if len(cartList) > 0 {

			for i := 0; i < len(cartList); i++ {
				if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == goodsAttr {
					cartList[i].Checked = !cartList[i].Checked

				}
				if cartList[i].Checked {
					allPrice += cartList[i].Price * float64(cartList[i].Num)
				}
			}
			models.Cookie.Set(c, "cartList", cartList)
			response = gin.H{
				"success":  true,
				"message":  "更新数据成功",
				"allPrice": allPrice,
			}
		} else {
			response = gin.H{
				"success": false,
				"message": "传入参数错误2",
			}
		}

	}

	c.JSON(200, response)
}
func (con CartController) ChangeAllCart(c *gin.Context) {
	flag, _ := models.Int(c.Query("flag"))

	//定义返回的数据
	var allPrice float64

	var response gin.H

	cartList := []models.Cart{}
	models.Cookie.Get(c, "cartList", &cartList)
	if len(cartList) > 0 {
		for i := 0; i < len(cartList); i++ {
			if flag == 1 {
				cartList[i].Checked = true
			} else {
				cartList[i].Checked = false
			}
			if cartList[i].Checked {
				allPrice += cartList[i].Price * float64(cartList[i].Num)
			}

		}
		//重新写入数据
		models.Cookie.Set(c, "cartList", cartList)

		response = gin.H{
			"success":  true,
			"message":  "更新数据成功",
			"allPrice": allPrice,
		}
	} else {
		response = gin.H{
			"success": false,
			"message": "传入参数错误",
		}
	}

	c.JSON(http.StatusOK, response)
}
func (con CartController) DelCart(c *gin.Context) {
	goodsId, _ := models.Int(c.Query("goods_id"))
	goodsColor := c.Query("goods_color")
	GoodsAttr := ""

	cartList := []models.Cart{}
	models.Cookie.Get(c, "cartList", &cartList)

	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == GoodsAttr {
			cartList = append(cartList[:i], cartList[(i+1):]...)
		}
	}
	models.Cookie.Set(c, "cartList", cartList)
	c.Redirect(302, "/cart")
}

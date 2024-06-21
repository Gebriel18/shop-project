package routers

import (
	"ginshop42/controllers/itying"
	"ginshop42/middlewares"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.DefaultController{}.Index)
		defaultRouters.GET("/category:id", itying.ProductController{}.Category)
		defaultRouters.GET("/detail", itying.ProductController{}.Detail)
		defaultRouters.GET("/product/getImgList", itying.ProductController{}.GetImgList)

		defaultRouters.GET("/cart", itying.CartController{}.Get)
		defaultRouters.GET("/cart/addCart", itying.CartController{}.AddCart)

		defaultRouters.GET("/cart/successTip", itying.CartController{}.AddCartSuccess)

		defaultRouters.GET("/cart/decCart", itying.CartController{}.DecCart)
		defaultRouters.GET("/cart/incCart", itying.CartController{}.IncCart)

		defaultRouters.GET("/cart/changeOneCart", itying.CartController{}.ChangeOneCart)
		defaultRouters.GET("/cart/changeAllCart", itying.CartController{}.ChangeAllCart)
		defaultRouters.GET("/cart/delCart", itying.CartController{}.DelCart)

		defaultRouters.GET("/pass/login", itying.PassController{}.Login)
		defaultRouters.GET("/pass/loginOut", itying.PassController{}.LoginOut)
		defaultRouters.GET("/pass/captcha", itying.PassController{}.Captcha)
		defaultRouters.POST("/pass/doLogin", itying.PassController{}.DoLogin)

		defaultRouters.GET("/pass/registerStep1", itying.PassController{}.RegisterStep1)
		defaultRouters.GET("/pass/registerStep2", itying.PassController{}.RegisterStep2)
		defaultRouters.GET("/pass/registerStep3", itying.PassController{}.RegisterStep3)
		defaultRouters.GET("/pass/sendCode", itying.PassController{}.SendCode)
		defaultRouters.GET("/pass/validateSmsCode", itying.PassController{}.ValidateSmsCode)
		defaultRouters.POST("/pass/doRegister", itying.PassController{}.DoRegister)

		defaultRouters.GET("/buy/checkout", middlewares.InitUserAuthMiddleware, itying.BuyController{}.CheckOut)
		defaultRouters.POST("/buy/doCheckout", middlewares.InitUserAuthMiddleware, itying.BuyController{}.DoCheckOut)
		defaultRouters.GET("/buy/pay", middlewares.InitUserAuthMiddleware, itying.BuyController{}.Pay)

		defaultRouters.POST("/address/addAddress", middlewares.InitUserAuthMiddleware, itying.AddressController{}.AddAddress)
		defaultRouters.POST("/address/editAddress", middlewares.InitUserAuthMiddleware, itying.AddressController{}.EditAddress)
		defaultRouters.GET("/address/changeDefaultAddress", middlewares.InitUserAuthMiddleware, itying.AddressController{}.ChangeDefaultAddress)
		defaultRouters.GET("/address/getOneAddressList", middlewares.InitUserAuthMiddleware, itying.AddressController{}.GetOneAddressList)
	}
}

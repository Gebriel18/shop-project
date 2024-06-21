package routers

import (
	"ginshop51/controllers/itying"
	"ginshop51/middlewares"

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
		defaultRouters.GET("/pass/captcha", itying.PassController{}.Captcha)

		defaultRouters.GET("/pass/registerStep1", itying.PassController{}.RegisterStep1)
		defaultRouters.GET("/pass/registerStep2", itying.PassController{}.RegisterStep2)
		defaultRouters.GET("/pass/registerStep3", itying.PassController{}.RegisterStep3)
		defaultRouters.GET("/pass/sendCode", itying.PassController{}.SendCode)
		defaultRouters.GET("/pass/validateSmsCode", itying.PassController{}.ValidateSmsCode)
		defaultRouters.POST("/pass/doRegister", itying.PassController{}.DoRegister)
		defaultRouters.POST("/pass/doLogin", itying.PassController{}.DoLogin)
		defaultRouters.GET("/pass/loginOut", itying.PassController{}.LoginOut)
		//判断用户权限
		defaultRouters.GET("/buy/checkout", middlewares.InitUserAuthMiddleware, itying.BuyController{}.Checkout)
	}
}

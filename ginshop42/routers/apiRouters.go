package routers

import (
	"ginshop42/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/addressList", api.ApiController{}.AddressList)
		apiRouters.GET("/plist", api.ApiController{}.Plist)
	}
	apiV1Routers := r.Group("/api/v1")
	{
		apiV1Routers.GET("/", api.V1Controller{}.Index)
		apiV1Routers.GET("/navList", api.V1Controller{}.NavList)
		apiV1Routers.GET("/plist", api.V1Controller{}.Plist)
	}

	apiV2Routers := r.Group("/api/v2")
	{
		apiV2Routers.GET("/", api.V2Controller{}.Index)
		apiV2Routers.GET("/userlist", api.V2Controller{}.Userlist)
		apiV2Routers.GET("/plist", api.V2Controller{}.Plist)
	}
}

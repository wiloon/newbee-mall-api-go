package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type ManageOrderRouter struct {
}

func (r *ManageOrderRouter) InitManageOrderRouter(Router *gin.RouterGroup) {
	mallOrderRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())
	var mallOrderApi = v1.ApiGroupApp.ManageApiGroup.ManageOrderApi
	{
		mallOrderRouter.PUT("orders/checkDone", mallOrderApi.CheckDoneOrder)   // 发货
		mallOrderRouter.PUT("orders/checkOut", mallOrderApi.CheckOutOrder)     // 出库
		mallOrderRouter.PUT("orders/close", mallOrderApi.CloseOrder)           // 出库
		mallOrderRouter.GET("orders/:orderId", mallOrderApi.FindMallOrder)     // 根据ID获取MallOrder
		mallOrderRouter.GET("shoporders/:orderId", mallOrderApi.FindShopOrder) // get shop order by order id
		mallOrderRouter.GET("orders", mallOrderApi.GetMallOrderList)           // 获取MallOrder列表
		mallOrderRouter.GET("shoporders", mallOrderApi.GetMallShopOrderList)   // shop order list
		mallOrderRouter.POST("adminSaveOrder", mallOrderApi.AdminSaveOrder)    // admin save order
		mallOrderRouter.POST("adminDeleteOrder", mallOrderApi.AdminDeleteOrder)
		mallOrderRouter.POST("adminSaveShop", mallOrderApi.AdminSaveShop) // admin save shop
	}
}

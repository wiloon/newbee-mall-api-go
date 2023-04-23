package manage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/common/request"
	"main.go/model/common/response"
	mallReq "main.go/model/mall/request"
	"main.go/model/manage"
	"main.go/utils"
	"time"
)

type ManageOrderApi struct {
}

// CheckDoneOrder 发货
func (m *ManageOrderApi) CheckDoneOrder(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := mallOrderService.CheckDone(IDS); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// CheckOutOrder 出库
func (m *ManageOrderApi) CheckOutOrder(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := mallOrderService.CheckOut(IDS); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// CloseOrder 出库
func (m *ManageOrderApi) CloseOrder(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := mallOrderService.CloseOrder(IDS); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMallOrder 用id查询MallOrder
func (m *ManageOrderApi) FindMallOrder(c *gin.Context) {
	id := c.Param("orderId")
	if err, newBeeMallOrderDetailVO := mallOrderService.GetMallOrder(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(newBeeMallOrderDetailVO, c)
	}
}
func (m *ManageOrderApi) GetMallShopOrderList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	orderNo := c.Query("orderNo")
	orderStatus := c.Query("orderStatus")
	shopId := c.Query("shopId")
	if err, list, total := mallOrderService.GetMallShopOrderInfoList(pageInfo, orderNo, orderStatus, shopId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetMallOrderList 分页获取MallOrder列表
func (m *ManageOrderApi) GetMallOrderList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	orderNo := c.Query("orderNo")
	orderStatus := c.Query("orderStatus")
	shopId := c.Query("shopId")
	if err, list, total := mallOrderService.GetMallOrderInfoList(pageInfo, orderNo, orderStatus, shopId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}
func (m *ManageOrderApi) AdminSaveShop(c *gin.Context) {
	var saveShopParam mallReq.AdminSaveShopParam
	_ = c.ShouldBindJSON(&saveShopParam)

	var shop manage.MallShop
	shop.Name = saveShopParam.Name
	shop.Owner = saveShopParam.OwnerId
	shop.CreateTime = common.JSONTime{Time: time.Now()}

	if err := global.GVA_DB.Save(&shop).Error; err != nil {
		response.FailWithMessage("生成店铺失败:"+err.Error(), c)
	}

	response.OkWithData("ok", c)
}
func (m *ManageOrderApi) AdminSaveOrder(c *gin.Context) {
	var saveOrderParam mallReq.AdminSaveOrderParam
	_ = c.ShouldBindJSON(&saveOrderParam)
	token := c.GetHeader("token")
	global.GVA_LOG.Info(fmt.Sprintf("token: %v, params: %+v", token, saveOrderParam))

	//生成订单号
	orderNo := utils.GenOrderNo()
	priceTotal := 0
	//保存订单
	var newBeeMallOrder manage.MallOrder
	newBeeMallOrder.OrderNo = orderNo
	newBeeMallOrder.UserId = saveOrderParam.Member
	err, goodsInfo := mallGoodsInfoService.GetMallGoodsInfo(saveOrderParam.Goods)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败"+err.Error(), c)
	}

	priceTotal = saveOrderParam.Number * goodsInfo.SellingPrice
	global.GVA_LOG.Info(fmt.Sprintf("goods info: %+v, price total:=·]%v", goodsInfo, priceTotal))

	newBeeMallOrder.CreateTime = common.JSONTime{Time: time.Now()}
	newBeeMallOrder.UpdateTime = common.JSONTime{Time: time.Now()}
	newBeeMallOrder.TotalPrice = priceTotal
	newBeeMallOrder.ExtraInfo = ""
	newBeeMallOrder.OrderStatus = saveOrderParam.OrderStatus
	newBeeMallOrder.PayType = saveOrderParam.PayType
	//生成订单项并保存订单项纪录
	if err = global.GVA_DB.Save(&newBeeMallOrder).Error; err != nil {
		response.FailWithMessage("生成订单失败:"+err.Error(), c)
	}

	orderItem := manage.MallOrderItem{}
	orderItem.OrderId = newBeeMallOrder.OrderId
	orderItem.CreateTime = common.JSONTime{Time: time.Now()}
	orderItem.GoodsCount = saveOrderParam.Number
	orderItem.GoodsId = saveOrderParam.Goods
	orderItem.GoodsCoverImg = goodsInfo.GoodsCoverImg
	orderItem.GoodsName = goodsInfo.GoodsName
	orderItem.SellingPrice = goodsInfo.SellingPrice
	if err = global.GVA_DB.Save(&orderItem).Error; err != nil {
		response.FailWithMessage("生成订单失败:"+err.Error(), c)
	}

	response.OkWithData("ok", c)
}

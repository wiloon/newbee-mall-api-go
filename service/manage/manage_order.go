package manage

import (
	"errors"
	"github.com/jinzhu/copier"
	"main.go/global"
	"main.go/model/common"
	"main.go/model/common/enum"
	"main.go/model/common/request"
	"main.go/model/manage"
	manageRes "main.go/model/manage/response"
	"strconv"
	"time"
)

type ManageOrderService struct {
}

// CheckDone 修改订单状态为配货成功
func (m *ManageOrderService) CheckDone(ids request.IdsReq) (err error) {
	var orders []manage.MallOrder
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error
	var errorOrders string
	if len(orders) != 0 {
		for _, order := range orders {
			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}
			if order.OrderStatus != enum.ORDER_PAID.Code() {
				errorOrders = order.OrderNo + " "
			}
		}
		if errorOrders == "" {
			if err = global.GVA_DB.Where("order_id in ?", ids.Ids).UpdateColumns(manage.MallOrder{OrderStatus: 2, UpdateTime: common.JSONTime{Time: time.Now()}}).Error; err != nil {
				return err
			}
		} else {
			return errors.New("订单的状态不是支付成功无法执行出库操作")
		}
	}
	return
}

// CheckOut 出库
func (m *ManageOrderService) CheckOut(ids request.IdsReq) (err error) {
	var orders []manage.MallOrder
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error
	var errorOrders string
	if len(orders) != 0 {
		for _, order := range orders {
			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}
			if order.OrderStatus != enum.ORDER_PAID.Code() && order.OrderStatus != enum.ORDER_PACKAGED.Code() {
				errorOrders = order.OrderNo + " "
			}
		}
		if errorOrders == "" {
			if err = global.GVA_DB.Where("order_id in ?", ids.Ids).UpdateColumns(manage.MallOrder{OrderStatus: 3, UpdateTime: common.JSONTime{Time: time.Now()}}).Error; err != nil {
				return err
			}
		} else {
			return errors.New("订单的状态不是支付成功或配货完成无法执行出库操作")
		}
	}
	return
}

// CloseOrder 商家关闭订单
func (m *ManageOrderService) CloseOrder(ids request.IdsReq) (err error) {
	var orders []manage.MallOrder
	err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&orders).Error
	var errorOrders string
	if len(orders) != 0 {
		for _, order := range orders {
			if order.IsDeleted == 1 {
				errorOrders = order.OrderNo + " "
				continue
			}
			if order.OrderStatus == enum.ORDER_SUCCESS.Code() || order.OrderStatus < 0 {
				errorOrders = order.OrderNo + " "
			}
		}
		if errorOrders == "" {
			if err = global.GVA_DB.Where("order_id in ?", ids.Ids).UpdateColumns(manage.MallOrder{OrderStatus: enum.ORDER_CLOSED_BY_JUDGE.Code(), UpdateTime: common.JSONTime{Time: time.Now()}}).Error; err != nil {
				return err
			}
		} else {
			return errors.New("订单不能执行关闭操作")
		}
	}
	return
}
func (m *ManageOrderService) GetShopOrder(orderId string) (err error, list interface{}) {

	var shopOrders []ShopOrderResult
	db := global.GVA_DB.Model(&manage.MallOrder{})
	db.Select("tb_newbee_mall_order.order_id,tb_newbee_mall_order.order_no,tb_newbee_mall_order.user_id,g.shop_id,g.goods_id,i.goods_count,tb_newbee_mall_order.pay_type,tb_newbee_mall_order.order_status,tb_newbee_mall_order.create_time")
	db.Joins("join tb_newbee_mall_order_item i on tb_newbee_mall_order.order_id=i.order_id")
	db.Joins("join tb_newbee_mall_goods_info g on i.goods_id=g.goods_id")
	db.Joins("join shop s on g.shop_id=s.id")

	err = db.Limit(1).Offset(0).Where("tb_newbee_mall_order.order_id=?", orderId).Order("tb_newbee_mall_order.update_time desc").Scan(&shopOrders).Error
	if len(shopOrders) > 0 {
		return nil, shopOrders[0]
	} else {
		return err, shopOrders
	}
}

// GetMallOrder 根据id获取MallOrder记录
func (m *ManageOrderService) GetMallOrder(id string) (err error, newBeeMallOrderDetailVO manageRes.NewBeeMallOrderDetailVO) {
	var newBeeMallOrder manage.MallOrder
	if err = global.GVA_DB.Where("order_id = ?", id).First(&newBeeMallOrder).Error; err != nil {
		return
	}
	var orderItems []manage.MallOrderItem
	if err = global.GVA_DB.Where("order_id = ?", newBeeMallOrder.OrderId).Find(&orderItems).Error; err != nil {
		return
	}
	//获取订单项数据
	if len(orderItems) > 0 {
		var newBeeMallOrderItemVOS []manageRes.NewBeeMallOrderItemVO
		copier.Copy(&newBeeMallOrderItemVOS, &orderItems)
		copier.Copy(&newBeeMallOrderDetailVO, &newBeeMallOrder)

		_, OrderStatusStr := enum.GetNewBeeMallOrderStatusEnumByStatus(newBeeMallOrderDetailVO.OrderStatus)
		_, payTapStr := enum.GetNewBeeMallOrderStatusEnumByStatus(newBeeMallOrderDetailVO.PayType)
		newBeeMallOrderDetailVO.OrderStatusString = OrderStatusStr
		newBeeMallOrderDetailVO.PayTypeString = payTapStr
		newBeeMallOrderDetailVO.NewBeeMallOrderItemVOS = newBeeMallOrderItemVOS
	}
	return
}

// GetMallOrderInfoList 分页获取MallOrder记录
func (m *ManageOrderService) GetMallOrderInfoList(info request.PageInfo, orderNo string, orderStatus string, shopId string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)
	// 创建db
	db := global.GVA_DB.Model(&manage.MallOrder{})
	if orderNo != "" {
		db.Where("order_no", orderNo)
	}
	// 0.待支付 1.已支付 2.配货完成 3:出库成功 4.交易成功 -1.手动关闭 -2.超时关闭 -3.商家关闭
	if orderStatus != "" {
		status, _ := strconv.Atoi(orderStatus)
		db.Where("order_status", status)
	}
	var mallOrders []manage.MallOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("update_time desc").Find(&mallOrders).Error
	return err, mallOrders, total
}

// tb_newbee_mall_order.user_id,g.shop_id,g.goods_id,i.goods_count,tb_newbee_mall_order.pay_type,tb_newbee_mall_order.order_status
type ShopOrderResult struct {
	OrderId      int    `json:"orderId"`
	OrderNo      string `json:"orderNo"`
	UserId       int    `json:"userId,omitempty"`
	ShopId       int    `json:"shopId,omitempty"`
	GoodsId      int    `json:"goodsId,omitempty"`
	PayType      int    `json:"payType"`
	TotalPrice   string `json:"totalPrice"`
	OrderStatus  int    `json:"orderStatus"`
	CreateTime   string `json:"createTime"`
	GoodsName    string `json:"goodsName"`
	GoodsCount   int    `json:"goodsCount"`
	SellingPrice string `json:"sellingPrice"`
}

// GetMallShopOrderInfoList 分页获取MallOrder记录
func (m *ManageOrderService) GetMallShopOrderInfoList(info request.PageInfo, shopId string) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageNumber - 1)

	var shopOrders []ShopOrderResult
	db := global.GVA_DB.Model(&manage.MallOrder{})
	db.Select("tb_newbee_mall_order.order_no,tb_newbee_mall_order.total_price,tb_newbee_mall_order.order_status,tb_newbee_mall_order.create_time,i.goods_name,i.goods_count,i.selling_price")
	db.Joins("join tb_newbee_mall_order_item i on tb_newbee_mall_order.order_id=i.order_id")
	db.Joins("join tb_newbee_mall_goods_info g on i.goods_id=g.goods_id")
	db.Joins("join shop s on g.shop_id=s.id")

	if shopId != "" {
		shopIdInt, _ := strconv.Atoi(shopId)
		db.Where("shop_id", shopIdInt)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Order("tb_newbee_mall_order.update_time desc").Scan(&shopOrders).Error
	return err, shopOrders, total
}

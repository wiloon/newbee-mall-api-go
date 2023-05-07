package response

import "main.go/model/common"

type NewBeeMallOrderDetailVO struct {
	OrderId                int                     `json:"orderId"`
	OrderNo                string                  `json:"orderNo"`
	TotalPrice             float32                 `json:"totalPrice"`
	PayType                int                     `json:"payType"`
	PayTypeString          string                  `json:"payTypeString"`
	OrderStatus            int                     `json:"orderStatus"`
	OrderStatusString      string                  `json:"orderStatusString"`
	CreateTime             common.JSONTime         `json:"createTime"`
	NewBeeMallOrderItemVOS []NewBeeMallOrderItemVO `json:"newBeeMallOrderItemVOS"`
}

type NewBeeMallOrderItemVO struct {
	GoodsId       int     `json:"goodsId"`
	GoodsName     string  `json:"goodsName"`
	GoodsCount    int     `json:"goodsCount"`
	GoodsCoverImg string  `json:"goodsCoverImg"`
	SellingPrice  float32 `json:"sellingPrice"`
}

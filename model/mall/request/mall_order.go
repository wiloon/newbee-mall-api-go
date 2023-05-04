package request

type PaySuccessParams struct {
	OrderNo string `json:"orderNo"`
	PayType int    `json:"payType"`
}

type OrderSearchParams struct {
	Status     string `form:"status"`
	PageNumber int    `form:"pageNumber"`
}

type SaveOrderParam struct {
	CartItemIds []int `json:"cartItemIds"`
	AddressId   int   `json:"addressId"`
}

type AdminSaveOrderParam struct {
	OrderId     int    `json:"orderId"`
	OrderNo     string `json:"orderNo"`
	Member      int    `json:"member"`
	Shop        int    `json:"shop"`
	Goods       int    `json:"goods"`
	PayType     int    `json:"payType"`
	Number      int    `json:"number"`
	OrderStatus int    `json:"orderStatus"`
	CreateTime  string `json:"createTime"`
}
type AdminSaveShopParam struct {
	Name    string `json:"name"`
	OwnerId int    `json:"ownerId"`
}

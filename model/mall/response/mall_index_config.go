package response

type MallIndexConfigGoodsResponse struct {
	GoodsId       int     `json:"goodsId"`
	GoodsName     string  `json:"goodsName"`
	GoodsIntro    string  `json:"goodsIntro"`
	GoodsCoverImg string  `json:"goodsCoverImg"`
	SellingPrice  float32 `json:"sellingPrice"`
	Tag           string  `json:"tag"`
}

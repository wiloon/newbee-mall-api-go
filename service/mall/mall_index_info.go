package mall

import (
	"main.go/global"
	"main.go/model/mall/response"
	"main.go/model/manage"
	"main.go/utils"
)

type MallIndexInfoService struct {
}

type ShopGoodsResult struct {
	SumValue      int     `json:"sumValue"`
	GoodsCoverImg string  `json:"goodsCoverImg"`
	GoodsName     string  `json:"goodsName"`
	SellingPrice  float32 `json:"sellingPrice"`
}

func (m *MallIndexInfoService) GetShopGoods(shop int) (err error, goodsList interface{}) {
	var mallGoodsInfos []manage.MallGoodsInfo
	global.GVA_DB.Where("shop_id=?", shop).Find(&mallGoodsInfos)

	var result []ShopGoodsResult
	for _, v := range mallGoodsInfos {
		goodId := v.GoodsId
		sr := ShopGoodsResult{}
		global.GVA_DB.Select("sum(goods_count) as sum_value").Table("tb_newbee_mall_order_item").Where("goods_id=?", goodId).Find(&sr)
		sr.GoodsCoverImg = v.GoodsCoverImg
		sr.GoodsName = v.GoodsName
		sr.SellingPrice = v.SellingPrice
		result = append(result, sr)
	}
	return err, result
}

// GetConfigGoodsForIndex 首页返回相关IndexConfig
func (m *MallIndexInfoService) GetConfigGoodsForIndex(configType int, num int) (err error, list interface{}) {
	var indexConfigs []manage.MallIndexConfig
	err = global.GVA_DB.Where("config_type = ?", configType).Where("is_deleted = 0").Order("config_rank desc").Limit(num).Find(&indexConfigs).Error
	if err != nil {
		return
	}
	// 获取商品id
	var ids []int
	for _, indexConfig := range indexConfigs {
		ids = append(ids, indexConfig.GoodsId)
	}
	// 获取商品信息
	var goodsList []manage.MallGoodsInfo
	err = global.GVA_DB.Where("goods_id in ?", ids).Find(&goodsList).Error
	var indexGoodsList []response.MallIndexConfigGoodsResponse
	// 超出30个字符显示....
	for _, indexGoods := range goodsList {
		res := response.MallIndexConfigGoodsResponse{
			GoodsId:       indexGoods.GoodsId,
			GoodsName:     utils.SubStrLen(indexGoods.GoodsName, 30),
			GoodsIntro:    utils.SubStrLen(indexGoods.GoodsIntro, 30),
			GoodsCoverImg: indexGoods.GoodsCoverImg,
			SellingPrice:  indexGoods.SellingPrice,
			Tag:           indexGoods.Tag,
		}
		indexGoodsList = append(indexGoodsList, res)
	}
	return err, indexGoodsList
}

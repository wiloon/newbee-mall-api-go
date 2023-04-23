package manage

import "main.go/model/common"

// MallShop 商城店铺信息
type MallShop struct {
	Id         int             `json:"id" form:"id" gorm:"primarykey;AUTO_INCREMENT"`
	Name       string          `json:"name" form:"name" gorm:"column:name;comment:店铺名;type:varchar(50);"`
	Owner      int             `json:"owner" form:"owner" gorm:"column:owner;comment:店主id;type:bigint;"`
	CreateTime common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:注册时间;type:datetime"`
}

// TableName MallShop 表名
func (MallShop) TableName() string {
	return "shop"
}

package manage

import (
	"main.go/model/common"
)

// TestData 商城用户信息
type TestData struct {
	TestDataId    int             `json:"testDataId" form:"testDataId" gorm:"primarykey;AUTO_INCREMENT"`
	NetworkType   string          `json:"networkType" form:"networkType" gorm:"column:network_type;comment:制式;type:varchar(50);"`
	Gps           string          `json:"gps" form:"gps" gorm:"column:gps;comment:gps;type:tinyint;"`
	Snr           string          `json:"snr" form:"snr" gorm:"column:snr;comment:信噪比;type:varchar(32);"`
	IntroduceSign string          `json:"introduceSign" form:"introduceSign" gorm:"column:introduce_sign;comment:个性签名;type:varchar(100);"`
	IsDeleted     int             `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:注销标识字段(0-正常 1-已注销);type:tinyint"`
	LockedFlag    int             `json:"lockedFlag" form:"lockedFlag" gorm:"column:locked_flag;comment:锁定标识字段(0-未锁定 1-已锁定);type:tinyint"`
	CreateTime    common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:注册时间;type:datetime"`
}

// TableName MallUser 表名
func (TestData) TableName() string {
	return "test_data"
}

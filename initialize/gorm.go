package initialize

import (
	"fmt"
	"gorm.io/gorm"
	"main.go/global"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	fmt.Println("gorm init")
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

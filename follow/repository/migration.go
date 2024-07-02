package repository

import (
	"follow/pkg/util"
	"os"
)

func migration() {
	//自动迁移模式
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&Follow{},
		)
	if err != nil {
		util.LogrusObj.Infoln("create follow table fail")
		os.Exit(0)
	}
	util.LogrusObj.Infoln("create follow table success")
}

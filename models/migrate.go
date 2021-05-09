package models

import (
	"blog_project/components/db"
	"log"
)

// AutoMigrate 自动迁移数据库
func AutoMigrate() {
	db := db.Get()

	err := db.Set("gorm:table_options", "CHARSET=utf8").AutoMigrate(
		&User{},
		&ThumbCollect{},
		&Follow{},
		&Blog{},
	).Error

	if err != nil {
		log.Panic(err.Error())
	}
}

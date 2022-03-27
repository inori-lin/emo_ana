package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func migration() {
	DB.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&User{})
}

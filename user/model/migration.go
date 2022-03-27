package model

import (
	_ "github.com/go-sql-driver/mysql"
)

func migration() {
	DB.Set(`gorm:table_options`, "charset=utf8mb4").
		AutoMigrate(&User{})
}

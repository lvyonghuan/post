package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB 初始化数据库
func InitDB() {
	dsn := "root:1145141919@tcp(127.0.0.1:3306)/chat"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

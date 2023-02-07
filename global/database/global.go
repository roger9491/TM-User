package database

import "gorm.io/gorm"

var (
	DBName   string
	Username string
	Password string
	Host     string
	Port     string

	DB *gorm.DB		// 資料庫對象
)
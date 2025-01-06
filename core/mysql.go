package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func InitGorm() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/fim_server_db?charset=utf8mb4&parseTime=True&loc=Local"
	var mysqlLogger logger.Interface
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)              // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 设置连接可复用的最大时间
	return db

}

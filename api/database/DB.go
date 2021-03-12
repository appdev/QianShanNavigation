package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)
	var err error
	DB, err = gorm.Open(sqlite.Open("nav.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}
}

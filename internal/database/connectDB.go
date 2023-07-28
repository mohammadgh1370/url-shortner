package database

import (
	"fmt"
	"github.com/mohammadgh1370/url-shortner/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func ConnectDB() (db *gorm.DB) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)

	gormConfig := gorm.Config{}
	if config.APP_ENV == config.ENV_LOCAL {
		gormConfig = gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	DB, err := gorm.Open(mysql.Open(dsn), &gormConfig)
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("Connected Successfully to the Database")

	return DB
}

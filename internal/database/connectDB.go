package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mohammadgh1370/url-shortner/internal/config"
	mysqlDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func ConnectDB() (db *gorm.DB) {
	dsn := getDsn()

	gormConfig := gorm.Config{}
	if config.APP_ENV == config.ENV_LOCAL {
		gormConfig = gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	}

	DB, err := gorm.Open(mysqlDriver.Open(dsn), &gormConfig)
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("Connected Successfully to the Database")

	return DB
}

func Migrate(name string) {
	dsn := getDsn()

	db, _ := sql.Open("mysql", dsn)
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://internal/database/migrations",
		"mysql",
		driver,
	)

	if err != nil {
		log.Fatal(err)
	}

	if name == "migrate" {
		if err := m.Up(); err != nil {
			fmt.Println(err.Error())
		}
	}

	if name == "migrate:rollback" {
		if err := m.Down(); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func getDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DB_USERNAME, config.DB_PASSWORD, config.DB_HOST, config.DB_PORT, config.DB_NAME)
}

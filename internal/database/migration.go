package database

import (
	"fmt"
	"github.com/mohammadgh1370/url-shortner/internal/model"
	"gorm.io/gorm"
)

var models = []interface{}{
	&model.User{},
	&model.Link{},
	&model.View{},
}

func RunAutoMigrations(DB *gorm.DB) {
	DB.AutoMigrate(models...)

	fmt.Println("Migrations complete")
}

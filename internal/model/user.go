package model

import (
	"time"
)

type User struct {
	Id        uint   `gorm:"primarykey" json:"-"`
	FirstName string `gorm:"type:varchar(191);"`
	LastName  string `gorm:"type:varchar(191);"`
	Username  string `gorm:"type:varchar(191); index:unique_username,unique"`
	Password  string `gorm:"type:varchar(191);" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

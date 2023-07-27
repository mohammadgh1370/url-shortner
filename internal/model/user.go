package model

import (
	"time"
)

type User struct {
	Id        uint   `gorm:"primarykey" json:"-"`
	FirstName string `gorm:"varchar(191)"`
	LastName  string `gorm:"varchar(191)"`
	Username  string `gorm:"varchar(191); index"`
	Password  string `gorm:"varchar(191)" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

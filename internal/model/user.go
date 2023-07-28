package model

import (
	"time"
)

type User struct {
	Id        uint   `gorm:"primarykey" json:"-"`
	FirstName string `gorm:"type:varchar(191);not null"`
	LastName  string `gorm:"type:varchar(191);not null"`
	Username  string `gorm:"type:varchar(191);not null;index:unique_username,unique"`
	Password  string `gorm:"type:varchar(191);not null" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

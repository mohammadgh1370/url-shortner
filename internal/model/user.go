package model

import (
	"time"
)

type User struct {
	Id        uint      `gorm:"primarykey" json:"-"`
	FirstName string    `gorm:"type:varchar(191);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(191);not null" json:"last_name"`
	Username  string    `gorm:"type:varchar(191);not null;index:unique_username,unique" json:"username"`
	Password  string    `gorm:"type:varchar(191);not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

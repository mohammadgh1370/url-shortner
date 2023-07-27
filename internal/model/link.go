package model

import (
	"gorm.io/gorm"
	"time"
)

type Link struct {
	Id        int    `gorm:"primarykey" json:"-"`
	User      User   `gorm:"ForeignKey:id;index:idx_user_id_url,unique"`
	Url       string `gorm:"text;index:idx_user_id_url,unique"`
	Link      string `gorm:"varchar(191)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

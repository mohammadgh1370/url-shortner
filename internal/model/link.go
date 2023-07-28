package model

import (
	"gorm.io/gorm"
	"time"
)

type Link struct {
	Id        uint   `gorm:"primarykey" json:"-"`
	UserId    uint   `gorm:"index:idx_user_id_url,unique"`
	User      User   `gorm:"foreignKey:UserId"`
	Url       string `gorm:"text;index:idx_user_id_url,unique"`
	Link      string `gorm:"type:varchar(191);"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

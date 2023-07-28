package model

import (
	"time"
)

type View struct {
	Id        uint `gorm:"primarykey" json:"-"`
	LinkId    uint
	Link      Link   `gorm:"ForeignKey:LinkId"`
	Ip        string `gorm:"type:varchar(191);"`
	Referrer  string `gorm:"type:varchar(191);"`
	UserAgent string `gorm:"text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

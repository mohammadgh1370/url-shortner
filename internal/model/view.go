package model

import (
	"time"
)

type View struct {
	Id        uint   `gorm:"primarykey" json:"-"`
	LinkId    uint   `gorm:"not null;"`
	Link      Link   `gorm:"ForeignKey:LinkId"`
	Ip        string `gorm:"type:varchar(191);not null;"`
	Referrer  string `gorm:"type:varchar(191);"`
	UserAgent string `gorm:"text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

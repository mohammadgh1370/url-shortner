package model

import (
	"time"
)

type View struct {
	Id        int    `gorm:"primarykey" json:"-"`
	Link      Link   `gorm:"ForeignKey:id"`
	Ip        string `gorm:"varchar(191)"`
	Referrer  string `gorm:"varchar(191)"`
	UserAgent string `gorm:"text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

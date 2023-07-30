package model

import (
	"time"
)

type View struct {
	Id        uint      `gorm:"primarykey" json:"-"`
	LinkId    uint      `gorm:"not null;" json:"link_id"`
	Link      Link      `gorm:"ForeignKey:LinkId" json:"-"`
	Ip        string    `gorm:"type:varchar(191);not null;" json:"ip"`
	Referer   string    `gorm:"type:varchar(191);" json:"referer"`
	UserAgent string    `gorm:"text" json:"user_agent"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

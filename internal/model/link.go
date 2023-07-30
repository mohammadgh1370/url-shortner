package model

import (
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
	"time"
)

type Link struct {
	Id        uint           `gorm:"primarykey" json:"id"`
	UserId    uint           `gorm:"not null;" json:"user_id"`
	User      User           `gorm:"foreignKey:UserId;" json:"-"`
	Url       string         `gorm:"not null;text;" json:"url"`
	Hash      string         `gorm:"not null;type:varchar(191);" json:"hash"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (link *Link) BeforeCreate(db *gorm.DB) error {
	model := new(Link)
	var hash string
	for {
		hash, _ = shortid.Generate()
		db.Model(&link).First(&model, Link{Hash: hash})
		if hash != model.Hash {
			break
		}
	}

	db.Set("hash", hash)
	link.Hash = hash
	return nil
}

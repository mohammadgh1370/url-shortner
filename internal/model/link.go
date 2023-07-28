package model

import (
	"github.com/teris-io/shortid"
	"gorm.io/gorm"
	"time"
)

type Link struct {
	Id        uint   `gorm:"primarykey" json:"-"`
	UserId    uint   `gorm:"not null;index:idx_user_id_url,unique"`
	User      User   `gorm:"foreignKey:UserId" json:"-"`
	Url       string `gorm:"not null;text;index:idx_user_id_url,unique"`
	Hash      string `gorm:"not null;type:varchar(191);index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
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

package mysql

import (
	"github.com/mohammadgh1370/url-shortner/internal/repository"
	"gorm.io/gorm"
)

type linkRepo struct {
	BaseRepo
}

func NewMysqlLinkRepo(db *gorm.DB) repository.ILinkRepo {
	return &linkRepo{
		BaseRepo{DB: db},
	}
}

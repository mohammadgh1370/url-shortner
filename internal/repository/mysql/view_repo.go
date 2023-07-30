package mysql

import (
	"github.com/mohammadgh1370/url-shortner/internal/repository"
	"gorm.io/gorm"
)

type viewRepo struct {
	BaseRepo
}

func NewMysqlViewRepo(db *gorm.DB) repository.IViewRepo {
	return &viewRepo{
		BaseRepo{DB: db},
	}
}

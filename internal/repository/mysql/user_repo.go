package mysql

import (
	"github.com/mohammadgh1370/url-shortner/internal/repository"
	"gorm.io/gorm"
)

type userRepo struct {
	BaseRepo
}

func NewMysqlUserRepo(db *gorm.DB) repository.IUserRepo {
	return &userRepo{
		BaseRepo{DB: db},
	}
}

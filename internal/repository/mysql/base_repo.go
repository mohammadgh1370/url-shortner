package mysql

import (
	"fmt"
	"gorm.io/gorm"
)

type BaseRepo struct {
	DB *gorm.DB
}

func (r *BaseRepo) First(model interface{}, condition interface{}) error {
	fmt.Println(model, condition)
	return r.DB.First(model, condition).Error
}

func (r *BaseRepo) Find(model interface{}, condition interface{}) error {
	return r.DB.Find(model, condition).Error
}

func (r *BaseRepo) Create(model interface{}) error {
	return r.DB.Create(model).Error
}

func (r *BaseRepo) FirstOrCreate(model interface{}, condition interface{}) error {
	return r.DB.FirstOrCreate(model, condition).Error
}

func (r *BaseRepo) UpdateOrCreate(model interface{}, condition interface{}) error {
	return r.DB.Where(condition).FirstOrCreate(model).Error
}

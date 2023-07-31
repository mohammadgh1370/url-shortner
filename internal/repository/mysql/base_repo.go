package mysql

import (
	"gorm.io/gorm"
)

type BaseRepo struct {
	DB *gorm.DB
}

func (r *BaseRepo) First(model interface{}, condition interface{}) error {
	return r.DB.First(model, condition).Error
}

func (r *BaseRepo) Find(model interface{}, condition interface{}) error {
	return r.DB.Order("id desc").Find(model, condition).Error
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

func (r *BaseRepo) Count(model interface{}, condition interface{}, count *int64) error {
	return r.DB.Model(&model).Where(condition).Count(count).Error
}

func (r *BaseRepo) Delete(model interface{}, condition interface{}) error {
	return r.DB.Where(condition).Delete(&model).Error
}

package repository

type IBaseRepo interface {
	First(model interface{}, condition interface{}) error
	Find(model interface{}, condition interface{}) error
	Create(model interface{}) error
	FirstOrCreate(model interface{}, condition interface{}) error
	UpdateOrCreate(model interface{}, condition interface{}) error
	Count(model interface{}, condition interface{}, count *int64) error
	Delete(model interface{}, condition interface{}) error
}

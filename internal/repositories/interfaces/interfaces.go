package interfaces

import "gorm.io/gorm"

type GORMRepositoryInterface interface {
	Raw(sql string, values ...interface{}) (tx *gorm.DB)
	Scan(dest interface{}) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Model(value interface{}) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	Updates(values interface{}) (tx *gorm.DB)
	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
	Exec(sql string, values ...interface{}) (tx *gorm.DB)
	Joins(query string, args ...interface{}) (tx *gorm.DB)
}

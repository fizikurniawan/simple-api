package repository

import (
	"simple-api/api/model"
	db "simple-api/database"

	"gorm.io/gorm"
)

type Food interface {
	FindAll(limit, offset int) ([]model.Food, error)
	CreateFood(data model.Food) (model.Food, error)
}

type food struct {
	db *gorm.DB
}

func NewFood() *food {
	db := db.DbManager()
	return &food{db}
}

func (r *food) FindAll(limit, offset int) ([]model.Food, error) {
	var datas []model.Food
	err := r.db.Model(&model.Food{}).Limit(limit).Offset(offset).Find(&datas).Error
	if err != nil {
		return datas, err
	}

	return datas, nil
}

func (r *food) CreateFood(data model.Food) (model.Food, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

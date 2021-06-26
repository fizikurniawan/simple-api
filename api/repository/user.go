package repository

import (
	"simple-api/api/model"
	db "simple-api/database"

	"gorm.io/gorm"
)

type User interface {
	FindAll() ([]model.User, error)
	FindByEmail(email string) (model.User, error)
	Save(data model.User) (model.User, error)
	UpdateUser(data model.User, userId int) (model.User, error)
}

type user struct {
	db *gorm.DB
}

func NewUser() *user {
	db := db.DbManager()
	return &user{db}
}

func (r *user) FindAll() ([]model.User, error) {
	var datas []model.User
	err := r.db.Find(&datas).Error
	if err != nil {
		return datas, err
	}
	return datas, nil
}

func (r *user) FindByEmail(email string) (model.User, error) {
	var data model.User
	err := r.db.Where("email = ?", email).First(&data).Error
	if err != nil {
		return data, err
	}

	return data, nil
}

func (r *user) Save(data model.User) (model.User, error) {
	err := r.db.Create(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *user) UpdateUser(data model.User, userId int) (model.User, error) {
	err := r.db.Where("ID = ?", userId).Save(&data).Error

	if err != nil {
		return data, err
	}

	return data, nil
}

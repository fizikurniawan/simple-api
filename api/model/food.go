package model

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	ID               uint           `gorm:"primaryKey;" json:"id"`
	DisplayName      string         `json:"name"`
	ShortDescription string         `json:"short_description"`
	Description      string         `json:"description"`
	MainPicture      uint           `json:"main_picture"`
	Price            float64        `json:"price"`
	CreatedAt        time.Time      `json:"created_at"`
	CreatedBy        string         `json:"created_by"`
	ModifiedAt       time.Time      `json:"modified_at"`
	ModifiedBy       string         `json:"modified_by"`
	DeletedAt        gorm.DeletedAt `json:"deleted_at"`
	DeletedBy        string         `json:"deleted_by"`
}

type FoodMedia struct {
	gorm.Model
	ID     uint `gorm:"primaryKey;" json:"id"`
	FileID uint `json:"file_id"`
	Order  int  `json:"order"`
	FoodID uint `json:"food_id"`
}

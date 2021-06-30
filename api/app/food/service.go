package food

import (
	"fmt"
	"simple-api/api/model"
	"simple-api/api/repository"
	res "simple-api/libs/util/response"
)

type Service interface {
	CreateFood(dto *FoodDTO) (model.Food, error)
	GetFoods(dto *model.PaginationDTO) ([]FoodDTOResponse, error)
}

type service struct {
	foodRespository repository.Food
}

func NewService() *service {
	foodRespository := repository.NewFood()
	return &service{foodRespository}
}

func (s *service) CreateFood(dto *FoodDTO) (model.Food, error) {
	var food model.Food

	food.DisplayName = dto.Name
	food.Description = dto.Description

	food, err := s.foodRespository.CreateFood(food)
	if err != nil {
		return food, res.BuildError(res.ErrUnprocessableEntity, err)
	}

	return food, nil
}

func (s *service) GetFoods(dto *model.PaginationDTO) ([]FoodDTOResponse, error) {
	var foods []FoodDTOResponse
	var limit, offset int

	if dto.PageSize == 0 {
		limit = 5
	} else {
		limit = dto.PageSize
	}

	if dto.Page <= 1 {
		offset = 0
	} else {
		offset = (dto.Page * limit) - limit
	}

	fmt.Println(limit, offset)
	results, err := s.foodRespository.FindAll(limit, offset)

	if err != nil {
		return foods, res.BuildError(res.ErrUnprocessableEntity, err)
	}

	for _, val := range results {
		food := FoodDTOResponse{
			val.DisplayName,
			val.Description,
		}
		foods = append(foods, food)

	}

	return foods, nil
}

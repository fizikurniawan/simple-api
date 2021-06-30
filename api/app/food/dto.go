package food

type FoodDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type FoodDTOResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

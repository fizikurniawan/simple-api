package account

import (
	"simple-api/api/repository"
	res "simple-api/libs/util/response"
)

type Service interface {
	GetProfile(email string) (ProfileResponseDTO, error)
	UpdateProfile(dto *UpdateProfileDTO, email string) (ProfileResponseDTO, error)
}

type service struct {
	userRepository repository.User
}

func NewService() *service {
	userRepository := repository.NewUser()
	return &service{userRepository}
}

func (s *service) GetProfile(email string) (ProfileResponseDTO, error) {
	var resDto ProfileResponseDTO
	user, err := s.userRepository.FindByEmail(email)

	if err != nil {
		return resDto, res.BuildError(res.ErrBadRequest, err)
	}

	resDto.Email = user.Email
	resDto.FirstName = user.FirstName
	resDto.LastName = user.LastName
	resDto.Name = user.Name
	resDto.Phone = user.Phone

	return resDto, nil
}

func (s *service) UpdateProfile(dto *UpdateProfileDTO, email string) (ProfileResponseDTO, error) {
	var resDto ProfileResponseDTO

	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return resDto, res.BuildError(res.ErrBadRequest, err)
	}

	user.Name = dto.FirstName + " " + dto.LastName
	user.FirstName = dto.FirstName
	user.LastName = dto.LastName
	user.Phone = dto.Phone

	user, err = s.userRepository.UpdateUser(user, int(user.ID))
	if err != nil {
		return resDto, res.BuildError(res.ErrUnprocessableEntity, err)
	}

	return resDto, nil

}

package authentication

import (
	"simple-api/api/model"
	"simple-api/api/repository"
	res "simple-api/libs/util/response"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(dto *RegisterDTO) (model.User, error)
	Login(dto *LoginDTO) (LoginResponseDTO, error)
}

type service struct {
	userRepository repository.User
}

func NewService() *service {
	userRepository := repository.NewUser()
	return &service{userRepository}
}

func (s *service) Register(dto *RegisterDTO) (model.User, error) {
	user, err := s.userRepository.FindByEmail(dto.Email)
	if err == nil {
		return user, res.BuildError(res.ErrDuplicate, err)
	}

	user.Name = dto.Name
	user.Email = dto.Email
	user.PasswordHash = dto.Password

	user.HashPassword()
	newUser, err := s.userRepository.Save(user)
	if err != nil {
		return newUser, res.BuildError(res.ErrUnprocessableEntity, err)
	}

	return newUser, nil
}

func (s *service) Login(dto *LoginDTO) (LoginResponseDTO, error) {
	var resDto LoginResponseDTO
	user, err := s.userRepository.FindByEmail(dto.Email)

	if err != nil {
		return resDto, res.BuildError(res.ErrUnauthorized, err)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(dto.Password)) != nil {
		return resDto, res.BuildError(res.ErrUnauthorized, err)
	}

	token, err := user.GenerateToken()
	if err != nil {
		return resDto, res.BuildError(res.ErrServerError, err)
	}

	resDto.Token = token

	return resDto, nil
}

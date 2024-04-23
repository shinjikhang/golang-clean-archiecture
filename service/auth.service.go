package service

import (
	"clean-architecture/dto/auth"
	"clean-architecture/model"
	"clean-architecture/repository"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(userDto dto.Register) error
	VerifyCredential(email string, passsword string) (bool, uint64)
}

type authService struct {
	authRepo repository.AuthRepo
}

func NewAuthService(authRepo repository.AuthRepo) *authService {
	return &authService{authRepo: authRepo}
}

func (service *authService) Register(userDto dto.Register) error {
	result, _ := service.authRepo.FindByEmail(userDto.Email)
	if result.Error == nil && result.RowsAffected > 0 {
		return errors.New("User already exists")
	}
	userModel := model.UserCreate{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
		Bio:      userDto.Bio,
	}
	if err := service.authRepo.Register(userModel); err != nil {
		return err
	}
	return nil
}

func (service *authService) VerifyCredential(email string, passsword string) (bool, uint64) {
	result, user := service.authRepo.FindByEmail(email)
	if result.Error == nil && user.UserId != 0 {
		return comparePassword([]byte(user.Password), []byte(passsword)), uint64(user.UserId)
	}
	return false, 0
}

func comparePassword(hashedPass []byte, plainPass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPass, plainPass)
	return err == nil
}

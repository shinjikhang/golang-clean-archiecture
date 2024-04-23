package service

import (
	"clean-architecture/dto/auth"
	"clean-architecture/model"
	"clean-architecture/repository"
	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

type AuthService interface {
	Register(userDto dto.Register) (*gorm.DB, *model.UserCreate)
	VerifyCredential(email string, passsword string) (bool, uint64)
}

type authService struct {
	authRepo repository.AuthRepo
}

func NewAuthService(authRepo repository.AuthRepo) *authService {
	return &authService{authRepo: authRepo}
}

func (service *authService) Register(userDto dto.Register) (*gorm.DB, *model.UserCreate) {
	userModel := model.UserCreate{
		Name:     userDto.Name,
		Email:    userDto.Email,
		Password: userDto.Password,
		Bio:      userDto.Bio,
	}
	return service.authRepo.Register(userModel)
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

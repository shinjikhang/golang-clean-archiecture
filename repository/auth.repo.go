package repository

import (
	"clean-architecture/helper"
	"clean-architecture/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type AuthRepo interface {
	Register(user model.UserCreate) error
	FindByEmail(email string) (*gorm.DB, model.User)
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *authRepo {
	return &authRepo{db: db}
}

func (repo *authRepo) Register(user model.UserCreate) error {
	activationCode := strconv.Itoa(rand.Intn(999999))
	emailSplit := strings.Split(user.Email, "@")
	user.UserName = emailSplit[0] + helper.GenerateRandomString(4)
	user.Password = hashAndSalt([]byte(user.Password))
	user.ActivationCode = activationCode
	user.ActivationCodeExpired = time.Now()
	user.Role = model.USER_ROLE

	if err := repo.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *authRepo) FindByEmail(email string) (*gorm.DB, model.User) {
	user := model.User{}
	userResult := repo.db.Where("email = ?", email).Take(&user)
	return userResult, user
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		panic("Failed to hash password")
	}
	return string(hash)
}

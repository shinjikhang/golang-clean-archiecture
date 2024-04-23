package controller

import (
	"clean-architecture/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	service service.IUserService
}

func NewUser(service service.IUserService) *User {
	return &User{
		service: service,
	}
}

func (h *User) GetUsers(c *gin.Context) {
	user, err := h.service.GetUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Trả về thông tin của bản ghi
	c.JSON(200, user)
}

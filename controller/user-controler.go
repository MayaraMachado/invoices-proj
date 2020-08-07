package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayaramachado/invoice-api/entity"
	"github.com/mayaramachado/invoice-api/service"
)

type UserController interface {
	Save(c *gin.Context)
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (ctrl *userController) Save(c *gin.Context) {
	var user entity.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	_, err = ctrl.service.Save(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Não foi possível criar usuário."})
		return
	}
	c.Status(http.StatusNoContent)
	return
}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mayaramachado/invoice-api/entity"
	"github.com/mayaramachado/invoice-api/service"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	userService service.UserService
	jWtService  service.JWTService
}

func NewLoginController(userService service.UserService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		userService: userService,
		jWtService:  jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var user entity.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		return ""
	}

	isAuthenticated := controller.userService.Login(user.Email, user.Password)
	if isAuthenticated {
		return controller.jWtService.GenerateToken(user.Email, true)
	}
	return ""
}

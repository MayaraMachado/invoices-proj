package main

import (
	"database/sql"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mayaramachado/invoice-api/controller"
	"github.com/mayaramachado/invoice-api/db"
	"github.com/mayaramachado/invoice-api/middlewares"
	"github.com/mayaramachado/invoice-api/repository"
	"github.com/mayaramachado/invoice-api/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	dbConnection *sql.DB = db.NewDB()

	invoiceRepository repository.InvoiceRepository = repository.NewInvoiceRepository(dbConnection)
	userRepository    repository.UserRepository    = repository.NewUserRepository(dbConnection)

	userService    service.UserService    = service.NewUserService(userRepository)
	invoiceService service.InvoiceService = service.NewInvoiceService(invoiceRepository)
	loginService   service.LoginService   = service.NewLoginService()
	jwtService     service.JWTService     = service.NewJWTService()

	userController    controller.UserController    = controller.NewUserController(userService)
	invoiceController controller.InvoiceController = controller.NewInvoiceController(invoiceService)
	loginController   controller.LoginController   = controller.NewLoginController(userService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	defer dbConnection.Close()
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Login Endpoint: Authentication + Token creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.Status(http.StatusUnauthorized)
		}
	})

	userRoutes := server.Group("/user")
	{
		userRoutes.POST("/register", userController.Save)
	}

	// JWT Authorization Middleware applies to "/api" only.
	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{

		apiRoutes.GET("/invoices", invoiceController.FindAll)

		apiRoutes.POST("/invoices", invoiceController.Save)

		apiRoutes.PUT("/invoices/:id", invoiceController.Update)

		apiRoutes.DELETE("/invoices/:id", invoiceController.Delete)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	server.Run(":" + port)
}

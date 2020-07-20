package main

import (
	"os"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mayaramachado/invoice-api/service"
	"github.com/mayaramachado/invoice-api/controller"
	"github.com/mayaramachado/invoice-api/middlewares"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	jwtService   service.JWTService   = service.NewJWTService()
	invoiceService service.InvoiceService = service.New()
	invoiceController controller.InvoiceController =  controller.New(invoiceService)
	loginService service.LoginService = service.NewLoginService()
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main(){
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())

	server.GET("/ping", func(c *gin.Context){
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
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	// JWT Authorization Middleware applies to "/api" only.
	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{

		apiRoutes.GET("/invoices", func(c *gin.Context){
			c.JSON(200, invoiceController.FindAll())
		})

		apiRoutes.POST("/invoices", func(c *gin.Context){
			err := invoiceController.Save(c)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{"message" : "Invoice input valid!"})
			}
		})
	}

	// port := os.Getenv("PORT")
	port:= "3000"
	server.Run(":" + port)
}
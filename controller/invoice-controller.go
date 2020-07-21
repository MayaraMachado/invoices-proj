package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mayaramachado/invoice-api/entity"
	"github.com/mayaramachado/invoice-api/service"
	"gopkg.in/go-playground/validator.v9"
)

type InvoiceController interface {
	FindAll() []entity.Invoice
	Save(c *gin.Context) error
}

type controller struct {
	service service.InvoiceService 
}

var validate *validator.Validate

func New(service service.InvoiceService) InvoiceController {
	return &controller {
		service: service,
	}
}

func (ctrl *controller) FindAll() []entity.Invoice{
	return ctrl.service.FindAll()
}

func (ctrl *controller) Save (c *gin.Context) error {
	var invoice entity.Invoice
	err := c.ShouldBindJSON(&invoice)
	if err != nil {
		return err
	}
	
	ctrl.service.Save(invoice)
	return nil
}
package controller

import (
	"strconv"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mayaramachado/invoice-api/entity"
	"github.com/mayaramachado/invoice-api/service"
	"gopkg.in/go-playground/validator.v9"
)

type InvoiceController interface {
	FindAll() []entity.Invoice
	Save(c *gin.Context) error
	Update(c *gin.Context) error
	Delete(c *gin.Context) error
}

type controller struct {
	service service.InvoiceService 
}

var validate *validator.Validate

func NewInvoiceController(service service.InvoiceService) InvoiceController {
	return &controller {
		service: service,
	}
}

func (ctrl *controller) FindAll() []entity.Invoice{
	// Permitir filtros por mês, ano e documento.


	invoices_list, err := ctrl.service.FindAll()
	if err != nil{
		panic(err)
	}

	return invoices_list
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

func (ctrl *controller) Update (c *gin.Context) error {
	var invoice entity.Invoice
	err := c.ShouldBindJSON(&invoice)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	
	invoice.Id = id
		
	ctrl.service.Update(invoice)
	return nil
}

func (ctrl *controller) Delete (c *gin.Context) error {
	var invoice entity.Invoice
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	invoice.Id = id
	
	ctrl.service.Delete(invoice)
	return nil
}
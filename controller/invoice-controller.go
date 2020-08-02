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
	FindAll(c *gin.Context) []entity.Invoice
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

func (ctrl *controller) FindAll(c *gin.Context) []entity.Invoice{
	// Permitir filtros por mês, ano e documento.
	query_params := c.Request.URL.Query()
	offset := 0
	limite := 10
	mes := 0
	ano := 0
	
	offset_string:= query_params.Get("offset")
	if offset_string != "" {
		offset_int, err := strconv.Atoi(offset_string)
		if err != nil {
			panic(err)
		}
		offset = offset_int
	}

	limite_string:= query_params.Get("limit")
	if limite_string != "" {
		limite_int, err := strconv.Atoi(limite_string)
		if err != nil {
			panic(err)
		}
		limite = limite_int
		fmt.Println("entrou: ", limite)
	}

	mes_string := query_params.Get("mes")
	if mes_string != "" {
		mes_int, err := strconv.Atoi(mes_string)
		if err != nil {
			panic(err)
		}
		mes = mes_int
	}

	ano_string := query_params.Get("ano")
	if ano_string != "" {
		ano_int, err := strconv.Atoi(ano_string)
		if err != nil {
			panic(err)
		}
		ano = ano_int
	}

	documento := query_params.Get("documento")
	order := query_params.Get("order")

	invoices_list, err := ctrl.service.FindAll(offset, limite, mes, ano, documento, order)
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
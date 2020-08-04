package controller

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/mayaramachado/invoice-api/entity"
	"github.com/mayaramachado/invoice-api/service"
	"gopkg.in/go-playground/validator.v9"
)

type InvoiceController interface {
	FindAll(c *gin.Context)
	Save(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
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

func (ctrl *controller) FindAll(c *gin.Context){
	// Permitir filtros por mês, ano e documento.
	query_params := c.Request.URL.Query()
	offset := 0
	limite := 10
	mes := 0
	ano := 0
	
	offset_string:= query_params.Get("offset")
	if offset_string != "" {
		offset_int, err := strconv.Atoi(offset_string)
		if err != nil || offset_int < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Queryparam inválido. Offset deve ser um valor inteiro positivo."})
        	return
		}
		offset = offset_int
	}

	limite_string:= query_params.Get("limit")
	if limite_string != "" {
		limite_int, err := strconv.Atoi(limite_string)
		if err != nil || limite_int <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Queryparam inválido. Limit deve ser um valor inteiro positivo maior que zero."})
			return
		}
		limite = limite_int
	}

	mes_string := query_params.Get("mes")
	if mes_string != "" {
		mes_int, err := strconv.Atoi(mes_string)
		if err != nil || 0 > mes_int && mes >= 12 {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Queryparam inválido. Mês deve ser um  valor inteiro positivo entre 1 e 12."})
			return
		}
		mes = mes_int
	}

	ano_string := query_params.Get("ano")
	if ano_string != "" {
		ano_int, err := strconv.Atoi(ano_string)
		if err != nil || 0 > ano_int {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Queryparam inválido. Ano deve ser um  valor inteiro positivo maior que zero."})
			return
		}
		ano = ano_int
	}

	documento := query_params.Get("documento")
	order := query_params.Get("order")

	invoices_list, err := ctrl.service.FindAll(offset, limite, mes, ano, documento, order)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Parâmetros inválidos."})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"offset" : offset, "max_por_pagina" : limite, "Data": invoices_list})
	return 
}

func (ctrl *controller) Save (c *gin.Context) {
	var invoice entity.Invoice
	err := c.ShouldBindJSON(&invoice)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	
	ctrl.service.Save(invoice)
	c.JSON(http.StatusCreated, invoice)
	return 
}

func (ctrl *controller) Update (c *gin.Context) {
	var invoice entity.Invoice
	err := c.ShouldBindJSON(&invoice)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return	}

	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Invoice não encontrado."})
		return
	}
	
	invoice.Id = id
		
	ctrl.service.Update(invoice)
	c.Status(http.StatusNoContent)
	return
}

func (ctrl *controller) Delete (c *gin.Context) {
	var invoice entity.Invoice
	id, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Invoice não encontrado."})
		return
	}
	invoice.Id = id
	
	ctrl.service.Delete(invoice)
	c.Status(http.StatusNoContent)
	return
}
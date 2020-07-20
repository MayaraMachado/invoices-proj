package service

import (
	"github.com/mayaramachado/invoice-api/entity"
)

type InvoiceService interface {
	Save(entity.Invoice) entity.Invoice
	FindAll() []entity.Invoice
}

type invoiceService struct {
	invoices []entity.Invoice
}

func New() InvoiceService {
	return &invoiceService{}
}

func (service *invoiceService) Save(invoice entity.Invoice) entity.Invoice {
	service.invoices = append(service.invoices, invoice)
	return invoice
}

func (service *invoiceService) FindAll() []entity.Invoice {
	return service.invoices
}

func (service *invoiceService) PartialUpdate(pk int32) []entity.Invoice {
	return service.invoices
}
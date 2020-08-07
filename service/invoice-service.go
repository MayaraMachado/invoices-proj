package service

import (
	"github.com/mayaramachado/invoice-api/entity"
	"github.com/mayaramachado/invoice-api/repository"
)

type InvoiceService interface {
	Save(invoice entity.Invoice) (entity.Invoice, error)
	Update(invoice entity.Invoice) (entity.Invoice, error)
	Delete(invoice entity.Invoice) (entity.Invoice, error)
	FindAll(offset int, limite int, mes int, ano int, documento string, order string) ([]entity.Invoice, error)
}

type invoiceService struct {
	invoiceRepository repository.InvoiceRepository
}

func NewInvoiceService(repo repository.InvoiceRepository) InvoiceService {
	return &invoiceService{
		invoiceRepository: repo,
	}
}

func (service *invoiceService) Save(invoice entity.Invoice) (entity.Invoice, error) {
	return service.invoiceRepository.Save(invoice)
}

func (service *invoiceService) FindAll(offset int, limite int, mes int, ano int, documento string, order string) ([]entity.Invoice, error) {
	return service.invoiceRepository.FindAll(offset, limite, mes, ano, documento, order)
}

func (service *invoiceService) Update(invoice entity.Invoice) (entity.Invoice, error) {
	_, err := service.invoiceRepository.GetByID(invoice.Id)
	if err != nil {
		return invoice, err
	}

	return service.invoiceRepository.Update(invoice)
}

func (service *invoiceService) Delete(invoice entity.Invoice) (entity.Invoice, error) {

	// verifica se o invoice existe:
	invoice_to_delete, err := service.invoiceRepository.GetByID(invoice.Id)
	if err != nil {
		return invoice, err
	}

	return service.invoiceRepository.Delete(invoice_to_delete)
}

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

// type invoiceService struct {
// 	db *sql.DB
// }

func New() InvoiceService {
	return &invoiceService{}
}

// func New(db *sql.DB) invoiceService {
// 	return invoiceService{
// 		db: db,
// 	}
// }

func (service *invoiceService) Save(invoice entity.Invoice) entity.Invoice {
	service.invoices = append(service.invoices, invoice)
	return invoice
}

func (service *invoiceService) FindAll() []entity.Invoice {
	return service.invoices
}

// func (service *invoiceService) FindAll() []entity.Invoice {
// 	invoices := []entity.Invoice{}
// 	stmt := fmt.Sprintf("SELECT * FROM invoices")
// 	result, err := service.db.Query(stmt)
// 	if err != nil {
// 		return invoices
// 	}

// 	defer result.Close()

// 	for result.Next() {
// 		invoice := entity.Invoice{}
// 		err := result.Scan(&invoice.Id, &invoice.ReferenceMonth, &invoice.ReferenceYear, &invoice.Document, &invoice.Description, &invoice.Amount, &invoice.IsActive, &invoice.CreatedAt, &invoice.DeactivatedAt)
// 		if err != nil {
// 			return invoices
// 		}
// 		invoices = append(invoices, invoice)
// 	}
// 	return invoices
// }

// func (service *invoiceService) PartialUpdate(pk int32) []entity.Invoice {
// 	return service.invoices
// }
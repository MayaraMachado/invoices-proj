package repository

import (
	"fmt"
	"github.com/mayaramachado/invoice-api/entity"
	"database/sql"
	"time"
	"reflect"
	"strconv"
)

type InvoiceRepository interface {
	Save(invoice entity.Invoice) (entity.Invoice, error)
	Update(invoice entity.Invoice) (entity.Invoice, error)
	Delete(invoice entity.Invoice) (entity.Invoice, error)
	GetByID(invoiceId uint64) (entity.Invoice, error)
	FindAll(offset int, limite int, mes int, ano int, documento string, order string) ([]entity.Invoice, error)
	CloseDbConnection()
}

type invoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(dbConnection *sql.DB) InvoiceRepository {
	return &invoiceRepository{
			db: dbConnection,
	}
}

func (repository *invoiceRepository) CloseDbConnection() {
	err := repository.db.Close()
	if err != nil{
			panic("Failed to close database!")
		}
}

func (repository *invoiceRepository) Save(invoice entity.Invoice) (entity.Invoice, error) {
	newInvoice := entity.Invoice{}
	active := 1
	query_string := `INSERT INTO invoices (reference_month, reference_year, document, description, amount, is_active, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;`
	result := repository.db.QueryRow(query_string, invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Document, invoice.Description, invoice.Amount, active, time.Now())
	
	err := result.Scan(&newInvoice.Id, &newInvoice.ReferenceMonth, &newInvoice.ReferenceYear, &newInvoice.Document, &newInvoice.Description, &newInvoice.Amount, &newInvoice.IsActive, &newInvoice.CreatedAt, &newInvoice.DeactivatedAt)
	if err != nil {
		return newInvoice, err
	}
	return newInvoice, nil
}

func (repository *invoiceRepository) Update(invoice entity.Invoice) (entity.Invoice, error) {
	updatedInvoice := entity.Invoice{}
	condition_str := ""
	
	invoice_reflect := reflect.ValueOf(invoice)
	type_of_invoice := invoice_reflect.Type()

	for i:= 0; i< invoice_reflect.NumField(); i++{
		field_name := type_of_invoice.Field(i).Tag.Get("json")
		field_value := invoice_reflect.Field(i).Interface()
		
		// Condicional com o id do invoice (melhorar)
		if i == 0 {
			condition_str = "WHERE id=" + strconv.FormatUint(invoice.Id, 10) + ";"
		} else {

			if !invoice_reflect.Field(i).IsZero() {
				query_string := fmt.Sprintf("UPDATE invoices SET %s=$1 " + condition_str, field_name)
				fmt.Printf(query_string)
				_, err := repository.db.Query(query_string, field_value)

				if err != nil {
					return updatedInvoice, err
				}
			}
		}
	}

	return updatedInvoice, nil
}

func (repository *invoiceRepository) Delete(invoice entity.Invoice) (entity.Invoice, error) {
	deletedInvoice := entity.Invoice{}
	deactivated := 0

	query_string := `UPDATE invoices SET is_active=$1, deactivated_at=$2 WHERE id=$3 RETURNING *;`
	result := repository.db.QueryRow(query_string, deactivated, time.Now(), invoice.Id)
	
	err := result.Scan(&deletedInvoice.Id, &deletedInvoice.ReferenceMonth, &deletedInvoice.ReferenceYear, &deletedInvoice.Document, &deletedInvoice.Description, &deletedInvoice.Amount, &deletedInvoice.IsActive, &deletedInvoice.CreatedAt, &deletedInvoice.DeactivatedAt)
	if err != nil {
		return deletedInvoice, err
	}
	return deletedInvoice, nil
}

func (repository *invoiceRepository) FindAll(offset int, limite int, mes int, ano int, documento string, order string) ([]entity.Invoice, error) {
		invoices := []entity.Invoice{}
		condition_str := ""
		query_string := fmt.Sprintf("SELECT * FROM invoices")
		
		//  OFFSET %s LIMIT %s;", order, strconv.Itoa(offset), strconv.Itoa(limite))
		
		// Montando cláusula WHERE
		if mes != 0 {
			condition_str =  condition_str + " reference_month=" + strconv.Itoa(mes)
		}

		if ano != 0 {
			if condition_str != ""{
				condition_str =  condition_str + " AND "
			}
			condition_str =  condition_str + " reference_year=" + strconv.Itoa(ano)
		}

		if documento != "" {
			if condition_str != ""{
				condition_str =  condition_str + " AND "
			}
			condition_str =  condition_str + " document='" + documento +"'"
		}

		if condition_str != "" {
			query_string = query_string + " WHERE " + condition_str
		}

		// Montando Cláusula ORDER BY
		if order != "" {
			query_string = query_string + " ORDER BY " + order + " ASC "
		}

		// Montando Query String
		query_string = query_string + " OFFSET " +  strconv.Itoa(offset) + " LIMIT "+ strconv.Itoa(limite) + ";"
		fmt.Println(query_string)
		result, err := repository.db.Query(query_string)
		if err != nil {
			return invoices, err
		}

		for result.Next() {
			invoice := entity.Invoice{}
			err := result.Scan(&invoice.Id, &invoice.ReferenceMonth, &invoice.ReferenceYear, &invoice.Document, &invoice.Description, &invoice.Amount, &invoice.IsActive, &invoice.CreatedAt, &invoice.DeactivatedAt)
			if err != nil {
				return invoices, err
			}
			invoices = append(invoices, invoice)
		}
		return invoices, nil
}

func (repository *invoiceRepository) GetByID(invoiceId uint64) (entity.Invoice, error){
	invoice := entity.Invoice{}
	query_string := `SELECT * FROM invoices WHERE id=$1 and is_active=1;`
	result := repository.db.QueryRow(query_string, invoiceId)
	err := result.Scan(&invoice.Id, &invoice.ReferenceMonth, &invoice.ReferenceYear, &invoice.Document, &invoice.Description, &invoice.Amount, &invoice.IsActive, &invoice.CreatedAt, &invoice.DeactivatedAt)
	if err != nil {
		return invoice, err
	}
	return invoice, nil
}

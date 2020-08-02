package entity

import (
	"time"
	"database/sql"
)

type Invoice struct {
	Id				uint64				`json:"id"`
	CreatedAt		time.Time			`json:"created_at" time_format:"2006-01-02" time_utc:"1"`
	ReferenceMonth	int					`json:"reference_month" binding:"required" `
	ReferenceYear	int					`json:"reference_year" binding:"required"`
	Document		string				`json:"document" binding:"required" `
	Description		string				`json:"description" binding:"required" `
	Amount			float64				`json:"amount" binding:"required" `
	IsActive		int       			`json:"is_active"`
	DeactivatedAt	sql.NullTime		`json:"-"`
}


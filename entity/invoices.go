package entity

import (
	"time"
)

type Invoice struct {
	Id				*int64		`json:"id"`
	CreatedAt		time.Time	`json:"createdAt" time_format:"2006-01-02" time_utc:"1"`
	ReferenceMonth	int			`json:"referenceMonth" binding:"required" `
	ReferenceYear	int			`json:"referenceYear" binding:"required"`
	Document		string		`json:"document" binding:"required" `
	Description		string		`json:"description" binding:"required" `
	Amount			float64		`json:"amount" binding:"required" `
	IsActive		int       	`json:"-"`
	DeactivatedAt	time.Time	`json:"-"`
}


package entity

import (
	"time"
	"database/sql"
)

type User struct {
	Id				uint64		 `json:"id"`
	Email 			string  	 `json:"email"`
	Password 		string  	 `json:"password"`
	CreatedAt		time.Time	 `json:"-" time_format:"2006-01-02" time_utc:"1"`
	IsActive		int       	 `json:"is_active"`
	DeactivatedAt	sql.NullTime `json:"-"`
}

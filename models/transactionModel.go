package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Name            string `gorm:"not null"`
	Description     *string
	Amount          float64 `gorm:"not null"`
	TransactionType string  `gorm:"not null"` // credit / debit
}

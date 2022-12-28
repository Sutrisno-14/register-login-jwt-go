package entity

import "time"

type Currecy struct {
	ID   int64   `gorm:"autoIncrement;primaryKey;column:id_currency"`
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`
}
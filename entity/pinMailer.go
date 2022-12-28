package entity

import "time"

type PinMailer struct {
	ID 			    int64  `gorm:"autoIncrement;primaryKey;column:id_pin"`
	VerifikasiID    int64	`gorm:"column:id_verifikasi;not null"`
	UserEmail		string
	Pin				int64   `gorm:"column:pin;not null"`
	UserStatus      bool 
	CreatedAt       time.Time `gorm:"not null"`
	UpdatedAt       time.Time `gorm:"not null"`
}
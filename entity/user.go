package entity

import "time"

type User struct {
	ID 		 		uint64  `gorm:"autoIncrement;primaryKey" json:"id_user"`
	Email    		string  `gorm:"uniqueIndex;type:varchar(255);not null" json:"email"`
	Password 		string  `gorm:"->;<-;not null" json:"-"`
	Token    		string  `gorm:"-" json:"token,omitempty"`
	Status          bool    `gorm:"not null"`
	CreatedAt       time.Time `gorm:"not null"`
	UpdatedAt       time.Time `gorm:"not null"`
}
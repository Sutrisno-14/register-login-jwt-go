package entity

import "time"

type Wallet struct {
	ID     		uint64    `gorm:"autoIncrement;primaryKey;column:id_wallet"`
	CurrencyID  uint64  
	UserID      uint64
	Amount      int64     `gorm:"column:amount"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

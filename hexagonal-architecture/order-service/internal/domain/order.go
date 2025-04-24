package domain

import "time"

type Order struct {
	ID          string    `gorm:"primaryKey;type:uuid"`
	CustomerID  string    `gorm:"type:uuid;not null"`
	TotalAmount float64   `gorm:"not null"`
	Status      string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

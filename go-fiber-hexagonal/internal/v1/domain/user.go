package domain

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
}

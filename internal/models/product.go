package models

type Product struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"unique;not null"`
	BrandID uint   `gorm:"not null;constraint:OnDelete:CASCADE"` //Do I need an index here?
}

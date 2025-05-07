package models

type Product struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `gorm:"unique;not null" json:"name"`
	BrandID uint   `gorm:"not null;constraint:OnDelete:CASCADE" json:"brandId"` //Do I need an index here?
}

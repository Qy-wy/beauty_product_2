package repositories

import (
	"github.com/Qy-wy/beauty_product.git/internal/models"
	"gorm.io/gorm"
)

type BrandRepository struct {
	db *gorm.DB
}

func NewBrandRepository(db *gorm.DB) *BrandRepository {
	return &BrandRepository{db: db}
}

func (r *BrandRepository) Create(brand *models.Brand) error {
	return r.db.Create(brand).Error
}

func (r *BrandRepository) GetAll() ([]models.Brand, error) {
	var brands []models.Brand
	err := r.db.Find(&brands).Error
	return brands, err
}

func (r *BrandRepository) GetByID(id uint) (*models.Brand, error) {
	var brand models.Brand
	err := r.db.First(&brand, id).Error
	return &brand, err
}

func (r *BrandRepository) Update(brand *models.Brand) error {
	return r.db.Save(brand).Error //which is better Save or Update?
}

func (r *BrandRepository) Delete(id uint) error {
	return r.db.Delete(&models.Brand{}, id).Error
}

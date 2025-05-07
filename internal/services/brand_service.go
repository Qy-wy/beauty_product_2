package services

import (
	"github.com/Qy-wy/beauty_product.git/internal/models"
	"github.com/Qy-wy/beauty_product.git/internal/repositories"
)

type BrandService struct {
	repo *repositories.BrandRepository
}

func NewBrandService(repo *repositories.BrandRepository) *BrandService {
	return &BrandService{repo: repo}
}

func (s *BrandService) Create(brand *models.Brand) error {
	return s.repo.Create(brand)
}

func (s *BrandService) GetAll() ([]models.Brand, error) {
	return s.repo.GetAll()
}

func (s *BrandService) GetByID(id uint) (*models.Brand, error) {
	return s.repo.GetByID(id)
}

func (s *BrandService) Update(brand *models.Brand) error {
	return s.repo.Update(brand)
}

func (s *BrandService) Delete(id uint) error {
	return s.repo.Delete(id)
}

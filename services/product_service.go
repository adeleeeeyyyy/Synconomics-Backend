package services

import (
	"Synconomics/models"
	"Synconomics/pkg/helpers"
	"Synconomics/repositories"
)

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo}
}

func (s *productService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *productService) GetAllProducts() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *productService) GetProductById(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) UpdateProduct(product *models.Product) error {
	// 1. Get existing product to compare image URL
	existing, err := s.repo.FindByID(product.ID)
	if err == nil && existing.ImageURL != "" && existing.ImageURL != product.ImageURL {
		// Image URL has changed, delete the old file
		// Note: We ignore errors from DeleteFile as it's not critical if the old file remains
		// but we should pass the actual path.
		// Usually ImageURL is something like "./public/uploads/filename.jpg"
		helpers.DeleteFile(existing.ImageURL)
	}

	return s.repo.Update(product)
}

func (s *productService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}

func (s *productService) GetProductsByBusinessId(businessID uint) ([]models.Product, error) {
	return s.repo.FindByBusinessID(businessID)
}
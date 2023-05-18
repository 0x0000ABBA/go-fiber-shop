package services

import "fiber-shop/internal/data/models"

func GetAllProducts() ([]models.Product, error) {
	return make([]models.Product, 0), nil // TODO implement me
}

func GetProductById(id string) (models.Product, error) {
	return models.Product{}, nil // TODO implement me
}

func AddProduct(p *models.Product) error {
	return nil // TODO implement me
}

package repositories

import (
	productRepositories "backend/repositories/product"
	userRepositories "backend/repositories/user"
	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

type IRepositoryRegistry interface {
	GetUser() userRepositories.IUserRepository
	GetProduct() productRepositories.IProductRepository
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &Registry{db: db}
}

func (r *Registry) GetUser() userRepositories.IUserRepository {
	return userRepositories.NewUserRepository(r.db)
}

func (r *Registry) GetProduct() productRepositories.IProductRepository {
	return productRepositories.NewProductRepository(r.db)
}

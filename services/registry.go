package services

import (
	"backend/repositories"
	productService "backend/services/product"
	userService "backend/services/user"
)

type Registry struct {
	repository repositories.IRepositoryRegistry
}

type IServiceRegistry interface {
	GetUser() userService.IUserService
	GetProduct() productService.IProductService
}

func NewServiceRegistry(repository repositories.IRepositoryRegistry) IServiceRegistry {
	return &Registry{
		repository: repository,
	}
}

func (r *Registry) GetUser() userService.IUserService {
	return userService.NewUserService(r.repository)
}

func (r *Registry) GetProduct() productService.IProductService {
	return productService.NewProductService(r.repository)
}

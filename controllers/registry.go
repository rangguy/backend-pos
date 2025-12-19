package controllers

import (
	productController "backend/controllers/product"
	userControllers "backend/controllers/user"
	"backend/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	GetUserController() userControllers.IUserController
	GetProductController() productController.IProductController
}

func NewControllerRegistry(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{
		service: service,
	}
}

func (r *Registry) GetUserController() userControllers.IUserController {
	return userControllers.NewUserController(r.service)
}

func (r *Registry) GetProductController() productController.IProductController {
	return productController.NewProductController(r.service)
}

package controllers

import (
	controllers "backend/controllers/user"
	"backend/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	GetUserController() controllers.IUserController
}

func NewRegistry(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{
		service: service,
	}
}

func (r *Registry) GetUserController() controllers.IUserController {
	return controllers.NewUserController(r.service)
}

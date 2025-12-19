package routes

import (
	"backend/controllers"
	productRoutes "backend/routes/product"
	userRoutes "backend/routes/user"
	"github.com/gofiber/fiber/v2"
)

type Registry struct {
	controller controllers.IControllerRegistry
	group      fiber.Router
}

type IRouterRegistry interface {
	Serve()
}

func NewRouteRegistry(controller controllers.IControllerRegistry, group fiber.Router) IRouterRegistry {
	return &Registry{
		controller: controller,
		group:      group,
	}
}

func (r *Registry) Serve() {
	r.userRoute().Run()
	r.productRoute().Run()
}

func (r *Registry) userRoute() userRoutes.IUserRoute {
	return userRoutes.NewUserRoute(r.controller, r.group)
}

func (r *Registry) productRoute() productRoutes.IProductRoute {
	return productRoutes.NewProductRoute(r.controller, r.group)
}

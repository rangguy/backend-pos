package routes

import (
	"backend/controllers"
	routes "backend/routes/user"
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
}

func (r *Registry) userRoute() routes.IUserRoute {
	return routes.NewUserRoute(r.controller, r.group)
}

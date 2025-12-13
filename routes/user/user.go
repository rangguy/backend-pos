package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

type UserRoute struct {
	controller controllers.IControllerRegistry
	group      fiber.Router
}

type IUserRoute interface {
	Run()
}

func NewUserRoute(controller controllers.IControllerRegistry, group fiber.Router) IUserRoute {
	return &UserRoute{
		controller: controller,
		group:      group,
	}
}

func (r *UserRoute) Run() {
	group := r.group.Group("/auth")
	group.Get("/user", middlewares.Authenticate(), r.controller.GetUserController().GetUserLogin)
	group.Get("/:uuid", middlewares.Authenticate(), r.controller.GetUserController().GetUserByUUID)
	group.Post("/login", r.controller.GetUserController().Login)
	group.Post("/register", r.controller.GetUserController().Register)
	group.Put("/:uuid", middlewares.Authenticate(), r.controller.GetUserController().Update)
}

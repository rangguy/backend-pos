package routes

import (
	"backend/constants"
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
	group.Post("/register",
		middlewares.CheckRole(
			[]string{constants.OwnerString},
		),
		r.controller.GetUserController().Register)
	group.Post("/login", r.controller.GetUserController().Login)
	group.Put("/:uuid", middlewares.Authenticate(), r.controller.GetUserController().Update)
}

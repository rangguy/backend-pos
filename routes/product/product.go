package routes

import (
	"backend/controllers"
	"backend/middlewares"
	"github.com/gofiber/fiber/v2"
)

type ProductRoute struct {
	controller controllers.IControllerRegistry
	group      fiber.Router
}

type IProductRoute interface {
	Run()
}

func NewProductRoute(controller controllers.IControllerRegistry, group fiber.Router) IProductRoute {
	return &ProductRoute{
		controller: controller,
		group:      group,
	}
}

func (r *ProductRoute) Run() {
	group := r.group.Group("/products")
	group.Get("", middlewares.Authenticate(), r.controller.GetProductController().GetAllWithoutPagination)
	group.Get("/pagination", middlewares.Authenticate(), r.controller.GetProductController().GetAllWithPagination)
	group.Get("/:uuid", middlewares.Authenticate(), r.controller.GetProductController().GetByUUID)
	group.Get("/code/:code", middlewares.Authenticate(), r.controller.GetProductController().GetByCode)

	group.Post("", middlewares.Authenticate(), r.controller.GetProductController().Create)
	group.Put("/:uuid", middlewares.Authenticate(), r.controller.GetProductController().Update)
	group.Delete("/:uuid", middlewares.Authenticate(), r.controller.GetProductController().Delete)
}

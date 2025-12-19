package cmd

import (
	"backend/common/response"
	"backend/config"
	"backend/constants"
	"backend/controllers"
	"backend/database/seeders"
	"backend/domain/models"
	"backend/middlewares"
	"backend/repositories"
	"backend/routes"
	"backend/services"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

var command = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		config.Init()
		db, err := config.InitDatabase()
		if err != nil {
			panic(err)
		}

		location, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			panic(err)
		}
		time.Local = location

		err = db.AutoMigrate(
			&models.Role{},
			&models.User{},
			&models.Product{},
		)
		if err != nil {
			panic(err)
		}

		seeders.NewSeederRegistry(db).Run()
		repository := repositories.NewRepositoryRegistry(db)
		service := services.NewServiceRegistry(repository)
		controller := controllers.NewControllerRegistry(service)

		app := fiber.New(fiber.Config{
			ErrorHandler: middlewares.HandlePanic(),
		})

		app.Use(func(c *fiber.Ctx) error {
			c.Set("Access-Control-Allow-Origin", "*")
			c.Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, PATCH")
			c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-service-name, x-api-key, x-request-at")

			if c.Method() == "OPTIONS" {
				return c.SendStatus(fiber.StatusNoContent)
			}

			return c.Next()
		})

		app.Use(middlewares.RateLimiter(
			config.Config.RateLimiterMaxRequest,
			time.Duration(config.Config.RateLimiterTimeSecond)*time.Second,
		))

		app.Get("/", func(c *fiber.Ctx) error {
			return c.Status(http.StatusOK).JSON(response.Response{
				Status:  constants.Success,
				Message: "Welcome to Backend POS",
			})
		})

		group := app.Group("/api/v1")
		route := routes.NewRouteRegistry(controller, group)
		route.Serve()

		app.Use(func(c *fiber.Ctx) error {
			return c.Status(http.StatusNotFound).JSON(response.Response{
				Status:  constants.Error,
				Message: fmt.Sprintf("Path %s", http.StatusText(http.StatusNotFound)),
			})
		})

		port := fmt.Sprintf(":%d", config.Config.Port)
		err = app.Listen(port)
		if err != nil {
			return
		}
	},
}

func Run() {
	err := command.Execute()
	if err != nil {
		panic(err)
	}
}

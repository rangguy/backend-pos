package cmd

import (
	"backend/config"
	"backend/domain/models"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
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
		)
		if err != nil {
			panic(err)
		}

	},
}

func Run() {
	err := command.Execute()
}

package main

import (
	"go-fiber-hexagonal/database"
	"go-fiber-hexagonal/internal/v1/adapter/http"
	"go-fiber-hexagonal/internal/v1/adapter/http/handler"
	"go-fiber-hexagonal/internal/v1/adapter/persistence"
	"go-fiber-hexagonal/internal/v1/domain"
	"go-fiber-hexagonal/internal/v1/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain.User{}) // optional

	repo := persistence.NewUserRepository(db)
	uc := usecase.NewUserUseCase(repo)
	h := handler.NewUserHandler(uc)

	app := fiber.New()
	http.SetupRoutes(app, h)

	log.Fatal(app.Listen(":3000"))
}

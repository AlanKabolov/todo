package main

import (
	"log"
	"todo-api-v2/database"
	"todo-api-v2/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Инициализация базы данных
	err := database.InitDB("postgres://postgres:1@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// руты
	app.Post("/tasks", func(c *fiber.Ctx) error {
		return handlers.CreateTask(c, database.Pool)
	})

	app.Get("/tasks", func(c *fiber.Ctx) error {
		return handlers.GetTasks(c, database.Pool)
	})

	app.Put("/tasks/:id", func(c *fiber.Ctx) error {
		return handlers.UpdateTask(c, database.Pool)
	})

	app.Delete("/tasks/:id", func(c *fiber.Ctx) error {
		return handlers.DeleteTask(c, database.Pool)
	})

	// Запуск сервера
	log.Fatal(app.Listen(":3000"))
}
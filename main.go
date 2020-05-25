package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/test", func(c *fiber.Ctx) {
		c.JSON("test")
	})

}

func main() {
	app := fiber.New()

	db, err := bolt.Open("store.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to store opened")
	defer db.Close()

	setupRoutes(app)

	app.Static("/", "./public")

	app.Listen(8000)

}

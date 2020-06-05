package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/gofiber/fiber"
	"github.com/mikkybang/url-shortener/link"
	"github.com/mikkybang/url-shortener/store"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/create", link.CreateUrl)

	app.Get("/:url", link.RedirectUrl)

	api.Get("/test", func(c *fiber.Ctx) {
		c.JSON("test")
	})

}

func setupStorage() {
	var err error

	db, err := bolt.Open("store.db", 0600, nil)

	store.Db = db

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to store opened")
}

func main() {
	app := fiber.New()

	app.Settings.CaseSensitive = true

	setupStorage()

	defer store.Db.Close()

	app.Static("/", "./public")

	setupRoutes(app)

	app.Listen(8000)

}

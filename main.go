package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/gofiber/fiber"
)



func main() {
	app := fiber.New()

	db, err := bolt.Open("store.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to store opened")
	defer db.Close()

	app.Static("/", "./public")

	app.Listen(8000)

}

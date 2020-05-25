package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/gofiber/fiber"
	"github.com/mikkybang/url-shortener/store"
)

func initStore() {

	store.storeConn, err := bolt.Open("store.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to store opened")
}

func main() {
	app := fiber.New()

	initStore()
	defer store.storeConn.Close()

	app.Static("/", "./public")

	app.Listen(8000)

}

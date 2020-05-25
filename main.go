package main


import {
	"log",
	"github.com/gofiber/giber",
	"github.com/boltdb/bolt",
	"fmt"
}



func initStore(){
	db, err := bolt.Open("store.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to store opened")
}

func main() {
	app := fiber.New()

	initStore()
	
	defer db.Close()

}
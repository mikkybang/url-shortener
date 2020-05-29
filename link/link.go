package link

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/gofiber/fiber"
)

type Link struct {
	OriginLink  string `json: "title"`
	 string `json: "author"`
	Rating int    `json: "rating"`
}

func createUrl(c *fiber.Ctx) {
	
}

func RedirectUrl(c *fiber.Ctx) {
	
}




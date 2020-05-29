package link

import (
	"encoding/json"
	"log"
	"time"

	"github.com/boltdb/bolt"
	"github.com/bradialabs/shortid"
	"github.com/gofiber/fiber"
	"github.com/mikkybang/url-shortener/store"
)

type Link struct {
	Url       string    `json: "Url"`
	hash      string    `json: "hash"`
	createdAt time.Time `json:"created_at" bson:"created_at"`
}

func createUrl(c *fiber.Ctx) {
	store := store.Db
	link := &Link
	if err := c.BodyParser(link); err != nil {
		c.Status(503).Send(err)
		return
	}
	link.hash = s.Generate()
	link.createdAt = time.Now()

	newLink, err := json.Marshal(link)

	if err != nil {
		c.Status(503).Send(err)
		return
	}

	err := store.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Link"))
		err := b.Put(newLink)
		return err
	})

	if err != null {
		c.Status(503).Send(err)
		return
	}

}

func RedirectUrl(c *fiber.Ctx) {

}

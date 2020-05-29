package link

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
	"github.com/bradialabs/shortid"
	"github.com/gofiber/fiber"
	"github.com/mikkybang/url-shortener/store"
)

type Link struct {
	Url       string    `json: "url"`
	hash      string    `json: "hash"`
	createdAt time.Time `json:"created_at" bson:"created_at"`
}

func CreateUrl(c *fiber.Ctx) {
	store := store.Db
	link := new(Link)
	if err := c.BodyParser(link); err != nil {
		c.Status(503).Send(err)
		return
	}
	s := shortid.New()
	link.hash = s.Generate()
	link.createdAt = time.Now()

	newLink, err := json.Marshal(link)

	if err != nil {
		c.Status(503).Send(err)
		return
	}

	error := store.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Link"))
		err := b.Put([]byte(link.Url), newLink)
		return err
	})

	if err != null {
		c.Status(503).Send(err)
		return
	}
	c.Status(200).JSON(newLink)
}

func RedirectUrl(c *fiber.Ctx) {

}

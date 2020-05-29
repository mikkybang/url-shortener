package link

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/bradialabs/shortid"
	"github.com/gofiber/fiber"
	"github.com/mikkybang/url-shortener/store"
)

type Link struct {
	Url  string `json: "url"`
	Hash string `json: "hash"`
}

var buk = []byte("link")

func CreateUrl(c *fiber.Ctx) {
	store := store.Db

	link := new(Link)

	if err := c.BodyParser(link); err != nil {
		c.Status(503).Send(err)
		return
	}
	s := shortid.New()
	link.Hash = s.Generate()

	key := []byte(link.Hash)
	value := []byte(link.Url)

	fmt.Println(key, value)

	err := store.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(buk)
		if err != nil {
			c.Status(503).Send(err)
		}
		return b.Put(key, value)
	})

	if err != nil {
		c.Status(500).Send(err)
	}

	c.Status(200).JSON(link)
}

func RedirectUrl(c *fiber.Ctx) {
	store := store.Db
	
	

	err := store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(buk)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", buk)
		}

		val := bucket.Get(key)
		url := string(val))
		
		if url != nil {
			return c.Redirect(url)
		}

	})

	if err != nil {
		c.Status(500).Send(err)
	}

}

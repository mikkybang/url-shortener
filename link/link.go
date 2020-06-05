package link

import (
	"bytes"
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

	fmt.Println(link.Hash, link.Url)

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

	urlkey := c.Params("url")

	key := []byte(urlkey)

	store.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(buk)
		if bucket == nil {
			return fmt.Errorf("Bucket %q not found!", buk)
		}

		url := ""

		cr := bucket.Cursor()

		for k, v := cr.First(); k != nil; k, v = cr.Next() {
			if bytes.Equal(key, k) {
				url = string(v)
				break
			}
		}

		c.Redirect("http://" + url)

		return nil
	})
}

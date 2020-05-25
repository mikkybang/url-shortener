package store

import (
	"github.com/boltdb/bolt"
)

var (
	storeConn *bolt.DB
)
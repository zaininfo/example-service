package pkg

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

// DB defines a database object
type DB interface {
	Ping() (string, error)
}

type db struct {
	client *redis.Client
}

// NewDB returns a new instance of database object
func NewDB() DB {
	return &db{
		client: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", DBHost, DBPort),
		}),
	}
}

// Ping returns the availability status of database
func (d *db) Ping() (string, error) {
	return d.client.Ping().Result()
}

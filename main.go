package main

import (
	"time"

	"github.com/nawazish-github/stock-management/client"
	"github.com/nawazish-github/stock-management/models"
)

func main() {
	c := client.Client
	c.Add(models.Item{
		ID:       0,
		Name:     "Ball",
		Category: "Toy",
		Expiry:   time.Now(),
	})
}

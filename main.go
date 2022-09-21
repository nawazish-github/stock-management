package main

import (
	"time"

	"github.com/nawazish-github/stock-management/controller"
	"github.com/nawazish-github/stock-management/models"
)

func main() {
	ctrlr := controller.NewControllerImpl()
	ctrlr.Add(models.Item{
		ID:       0,
		Name:     "Ball",
		Category: "Toy",
		Expiry:   time.Now(),
	})
}
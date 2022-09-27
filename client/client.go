package client

import (
	"github.com/nawazish-github/stock-management/controller"
	"github.com/nawazish-github/stock-management/models"
)

var Client = controller.NewControllerImpl()

func Add(item models.Item) {
	Client.Add(item)
}

func Remove(item models.Item) {
	Client.Remove(item)
}
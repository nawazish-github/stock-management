package handlers

import "github.com/nawazish-github/stock-management/models"

type EventHandler interface {
	Handle(models.Item)
}
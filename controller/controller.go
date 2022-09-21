package controller

import "github.com/nawazish-github/stock-management/models"

type Controller interface {
	Add(item models.Item)
	Remove(item models.Item)
}

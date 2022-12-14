package persistence

import "github.com/nawazish-github/stock-management/models"

type Repository interface {
	Add(item models.Item)
	Find(item models.Item) models.ItemStats
	Remove(item models.Item)
}

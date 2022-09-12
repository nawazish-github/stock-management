package database

import "github.com/nawazish-github/stock-management/persistence/models"


type Repository interface {
	Add(item models.Item)
	Find(item models.Item) models.ItemStats
	Remove(item models.Item)
}

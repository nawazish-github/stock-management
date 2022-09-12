package in_mem

import (
	"github.com/nawazish-github/stock-management/persistence/models"
)

type InMemImpl struct {
	// store is a map from item category to the list of items.
	// Example item category is ball, mapped list would be cricket ball, football etc.
	store map[string][]models.Item
}

func (i *InMemImpl) Add(item models.Item) {
	items, ok := i.store[item.Category]
	if ok {
		items = append(items, item)
		i.store[item.Category] = items
	} else {
		i.store[item.Category] = []models.Item{item}
	}
}

func (i *InMemImpl) Find(item models.Item) models.ItemStats {
	items := i.store[item.Category]
	return models.ItemStats{
		Items:       items,
		Count:       len(items),
		Description: "blah for now",
	}
}

func (i *InMemImpl) Remove(item models.Item) {
}

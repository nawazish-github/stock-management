package in_mem

import (
	"github.com/nawazish-github/stock-management/persistence/models"
)

type InMemImpl struct {
	// store is a map from item category to it's count.
	store map[string]int
}

func (i *InMemImpl) Add(item models.Item) {
	count, ok := i.store[item.Category]
	if ok {
		i.store[item.Category] = 1
	} else {
		count++
		i.store[item.Category] = count
	}
}

func (i *InMemImpl) Find(item models.Item) models.ItemStats {
	items := i.store[item.Category]
	return models.ItemStats{
		Item:        item.Category,
		Count:       items,
		Description: "blah for now",
	}
}

func (i *InMemImpl) Remove(item models.Item) {
	count, ok := i.store[item.Category]
	if ok {
		count--
		i.store[item.Category] = count
	}
}

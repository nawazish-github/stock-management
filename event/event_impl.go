package event

import (
	"github.com/nawazish-github/stock-management/models"
	"github.com/nawazish-github/stock-management/persistence"
)

type EventImpl struct {
	repo          persistence.Repository
	eventHandlers []EventHandler
}

// TO Trigger an event, we get the stats for the item
// And we check if the item is below level, if yes,
// then the handlers are triggered to trigger the event.
func (e *EventImpl) Trigger(item models.Item) {
	itemStats := e.repo.Find(item) // TODO: concurrent sync issues
	if isItemBelowLevel(itemStats) {
		for _, eventHandler := range e.eventHandlers {
			eventHandler.Trigger(item) // TODO: error handling
		}
	}
}

func isItemBelowLevel(itemStats models.ItemStats) bool {
	return true // TODO
}

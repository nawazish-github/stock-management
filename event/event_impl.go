package event

import (
	"github.com/nawazish-github/stock-management/event/handlers"
	"github.com/nawazish-github/stock-management/models"
	"github.com/nawazish-github/stock-management/persistence"
	"github.com/nawazish-github/stock-management/persistence/in_mem"
)

type EventImpl struct {
	repo          persistence.Repository
	eventHandlers []handlers.EventHandler
}

func NewEventImpl() Event {
	return &EventImpl{
		repo:          in_mem.NewInMemImpl(),
		eventHandlers: []handlers.EventHandler{handlers.NewConsoleEventHandlerImpl()},
	}
}

// TO Trigger an event, we get the stats for the item
// And we check if the item is below level, if yes,
// then the handlers are triggered to trigger the event.
func (e *EventImpl) Trigger(item models.Item) {
	for _, eventHandler := range e.eventHandlers {
		eventHandler.Handle(models.Event{
			Item: item,
		}) // TODO: error handling
	}
}

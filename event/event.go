package event

import (
	"github.com/nawazish-github/stock-management/event/handlers"
	"github.com/nawazish-github/stock-management/models"
)

type Event interface {
	Trigger(models.Item)
	AddHandlers(...handlers.EventHandler)
}

// EventHandler implementors would know exactly
// how to send the event to the recepient.

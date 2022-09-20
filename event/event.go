package event

import "github.com/nawazish-github/stock-management/models"

type Event interface {
	Trigger(models.Item)
}

// EventHandler implementors would know exactly 
// how to send the event to the recepient. 
type EventHandler interface {
	Trigger(models.Item)
}

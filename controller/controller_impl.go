package controller

import (
	"github.com/nawazish-github/stock-management/event"
	"github.com/nawazish-github/stock-management/models"
	"github.com/nawazish-github/stock-management/persistence"
	"github.com/nawazish-github/stock-management/persistence/in_mem"
)

type ControllerImpl struct {
	repo  persistence.Repository
	event event.Event
}

func NewControllerImpl() Controller {
	return &ControllerImpl{
		repo:  in_mem.NewInMemImpl(),
		event: event.NewEventImpl(),
	}
}

// TO Add an item to the system, we add it to the system
// and we trigger an event
func (cntrlr *ControllerImpl) Add(item models.Item) {
	cntrlr.repo.Add(item)
	itemStats := cntrlr.repo.Find(item)
	if trigger := isItemBelowMinReq(itemStats); trigger {
		cntrlr.event.Trigger(item)
	}
}

func (cntrlr *ControllerImpl) Remove(item models.Item) {
	cntrlr.repo.Remove(item) // TODO: Remove might return the next available stock stats but this would break CQ pattern
	itemStats := cntrlr.repo.Find(item)
	if trigger := isItemBelowMinReq(itemStats); trigger {
		cntrlr.event.Trigger(item)
	}
}

func isItemBelowMinReq(is models.ItemStats) bool {
	return is.Count < is.MinReq
}

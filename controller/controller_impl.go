package controller

import (
	"github.com/nawazish-github/stock-management/event"
	"github.com/nawazish-github/stock-management/models"
	"github.com/nawazish-github/stock-management/persistence"
)

type ControllerImpl struct {
	repo persistence.Repository
	event event.Event
}
// TO Add an item to the system, we add it to the system
// and we trigger an event
func (cntrlr *ControllerImpl) Add(item models.Item) {
	cntrlr.repo.Add(item)
	cntrlr.event.Trigger(item)	
}

func (cntrlr *ControllerImpl) Remove(item models.Item) {
	cntrlr.repo.Remove(item)
	cntrlr.event.Trigger(item)	
}

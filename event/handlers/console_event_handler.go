package handlers

import (
	"fmt"

	"github.com/nawazish-github/stock-management/models"
)

type ConsoleEventHandlerImpl struct{}

func NewConsoleEventHandlerImpl() EventHandler {
	return &ConsoleEventHandlerImpl{}
}

func (email *ConsoleEventHandlerImpl) Handle(event models.Event) {
	fmt.Printf("Console Event triggered: %v\n", event.Item)
}

package handlers

import "github.com/nawazish-github/stock-management/models"

type EmailEventHandlerImpl struct{}

func (email *EmailEventHandlerImpl) Handle(item models.Item) {
	// Send email
}
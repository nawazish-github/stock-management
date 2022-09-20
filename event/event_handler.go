package event

import "github.com/nawazish-github/stock-management/models"

type EmailHandlerImpl struct{}

func (email *EmailHandlerImpl) Trigger(item models.Item) {
	// Send email
}
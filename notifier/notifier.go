package notifier

import "github.com/panda2xx/alertmanager-webhook/model"

// Notifier is the interface for sending notifications.
type Notifier interface {
	Notify(payload model.WebhookPayload) error
}

package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/panda2xx/alertmanager-webhook/model"
	"github.com/panda2xx/alertmanager-webhook/notifier"
)

// AlertmanagerHandler handles webhooks from Alertmanager.
func AlertmanagerHandler(notifier notifier.Notifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		// Log the raw request body for debugging
		log.Printf("Received Alertmanager webhook: %s", string(body))

		var payload model.WebhookPayload
		if err := json.Unmarshal(body, &payload); err != nil {
			log.Printf("Error unmarshalling webhook payload: %v", err)
			http.Error(w, "Error unmarshalling webhook payload", http.StatusBadRequest)
			return
		}

		if err := notifier.Notify(payload); err != nil {
			log.Printf("Error sending notification: %v", err)
			http.Error(w, "Error sending notification", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Notification sent successfully"))
	}
}

package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/panda2xx/alertmanager-webhook/model"
)

// TelegramNotifier sends notifications to Telegram.
type TelegramNotifier struct {
	BotToken string
	ChatID   string
}

// NewTelegramNotifier creates a new TelegramNotifier.
func NewTelegramNotifier(botToken, chatID string) *TelegramNotifier {
	return &TelegramNotifier{
		BotToken: botToken,
		ChatID:   chatID,
	}
}

// Notify sends a notification to Telegram.
func (t *TelegramNotifier) Notify(payload model.WebhookPayload) error {
	telegramPayload := t.createTelegramPayload(payload)

	payloadBytes, err := json.Marshal(telegramPayload)
	if err != nil {
		return fmt.Errorf("error marshalling telegram payload: %w", err)
	}

	log.Printf("Sending Telegram payload: %s", string(payloadBytes))

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", t.BotToken)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error sending telegram notification: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("telegram notification failed with status code %d: %s", resp.StatusCode, string(body))
	}

	log.Println("Telegram notification sent successfully")
	return nil
}

func (t *TelegramNotifier) createTelegramPayload(payload model.WebhookPayload) model.TelegramSendMessagePayload {
	var emoji string
	if payload.Status == "resolved" {
		emoji = "✅"
	} else {
		emoji = "❌"
	}

	var message strings.Builder
	message.WriteString(fmt.Sprintf("*%s [%s] %s*\n\n", emoji, strings.ToUpper(payload.Status), payload.CommonLabels["alertname"]))

	for _, alert := range payload.Alerts {
		if len(alert.Labels) > 0 {
			message.WriteString("*Labels*:\n")
			for key, value := range alert.Labels {
				message.WriteString(fmt.Sprintf("- `%s`: `%s`\n", key, value))
			}
		}

		if summary, ok := alert.Annotations["summary"]; ok {
			message.WriteString(fmt.Sprintf("*Summary*: %s\n", summary))
		}
		if desc, ok := alert.Annotations["description"]; ok {
			message.WriteString(fmt.Sprintf("*Description*: %s\n\n", desc))
		}
	}

	return model.TelegramSendMessagePayload{
		ChatID:    t.ChatID,
		Text:      message.String(),
		ParseMode: "Markdown",
	}
}

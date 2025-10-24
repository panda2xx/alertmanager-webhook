package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/panda2xx/alertmanager-webhook/model"
)

// DiscordNotifier sends notifications to Discord.
type DiscordNotifier struct {
	WebhookURL string
	Username   string
	AvatarURL  string
}

// NewDiscordNotifier creates a new DiscordNotifier.
func NewDiscordNotifier(webhookURL, username, avatarURL string) *DiscordNotifier {
	return &DiscordNotifier{
		WebhookURL: webhookURL,
		Username:   username,
		AvatarURL:  avatarURL,
	}
}

// Notify sends a notification to Discord.
func (d *DiscordNotifier) Notify(payload model.WebhookPayload) error {
	discordPayload := d.createDiscordPayload(payload)

	payloadBytes, err := json.Marshal(discordPayload)
	if err != nil {
		return fmt.Errorf("error marshalling discord payload: %w", err)
	}

	log.Printf("Sending Discord payload: %s", string(payloadBytes))

	resp, err := http.Post(d.WebhookURL, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error sending discord notification: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("discord notification failed with status code %d: %s", resp.StatusCode, string(body))
	}

	log.Println("Discord notification sent successfully")
	return nil
}

func (d *DiscordNotifier) createDiscordPayload(payload model.WebhookPayload) model.DiscordWebhookPayload {
	var color int
	var emoji string
	if payload.Status == "resolved" {
		color = 2263842 // green
		emoji = "✅"
	} else {
		color = 16711680 // red
		emoji = "❌"
	}

	var description strings.Builder
	for _, alert := range payload.Alerts {
		description.WriteString(fmt.Sprintf("**Status**: %s\n", alert.Status))

		/*
			if len(alert.Labels) > 0 {
				description.WriteString("**Labels**:\n")
				for key, value := range alert.Labels {
					description.WriteString(fmt.Sprintf("- `%s`: `%s`\n", key, value))
				}
			}
		*/

		if summary, ok := alert.Annotations["summary"]; ok {
			description.WriteString(fmt.Sprintf("**Summary**: %s\n", summary))
		}
		if desc, ok := alert.Annotations["description"]; ok {
			description.WriteString(fmt.Sprintf("**Description**: %s\n\n", desc))
		}
	}

	embed := model.Embed{
		Title:       fmt.Sprintf("%s [%s] %s", emoji, strings.ToUpper(payload.Status), payload.CommonLabels["alertname"]),
		Description: description.String(),
		// URL:         payload.ExternalURL,
		Color:     color,
		Timestamp: time.Now().Format(time.RFC3339),
		Footer: &model.Footer{
			Text: "Alertmanager Webhook",
		},
	}

	return model.DiscordWebhookPayload{
		Username:  d.Username,
		AvatarURL: d.AvatarURL,
		Embeds:    []model.Embed{embed},
	}
}

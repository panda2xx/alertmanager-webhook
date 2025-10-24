package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/panda2xx/alertmanager-webhook/handler"
	"github.com/panda2xx/alertmanager-webhook/notifier"
)

func fatal(v ...interface{}) {
	fmt.Fprintln(os.Stderr, v...)
	os.Exit(1)
}

func fatalf(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", v...)
	os.Exit(1)
}

func main() {
	if os.Getenv("ENABLE_LOGGING") != "true" {
		log.SetOutput(io.Discard)
	}
	// General settings
	listenAddress := os.Getenv("LISTEN_ADDRESS")
	if listenAddress == "" {
		listenAddress = "0.0.0.0:9095"
	}

	// Select notifier based on environment variable
	notifierType := os.Getenv("NOTIFIER_TYPE")
	if notifierType == "" {
		notifierType = "discord" // Default to discord
	}

	var n notifier.Notifier

	switch strings.ToLower(notifierType) {
	case "discord":
		webhookURL := os.Getenv("DISCORD_WEBHOOK")
		if webhookURL == "" {
			fatal("NOTIFIER_TYPE is 'discord', but DISCORD_WEBHOOK environment variable is not set")
		}
		username := os.Getenv("DISCORD_USERNAME")
		avatarURL := os.Getenv("DISCORD_AVATAR_URL")
		n = notifier.NewDiscordNotifier(webhookURL, username, avatarURL)
	case "telegram":
		botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
		chatID := os.Getenv("TELEGRAM_CHAT_ID")
		if botToken == "" || chatID == "" {
			fatal("NOTIFIER_TYPE is 'telegram', but TELEGRAM_BOT_TOKEN or TELEGRAM_CHAT_ID is not set")
		}
		n = notifier.NewTelegramNotifier(botToken, chatID)
	default:
		fatalf("Unsupported NOTIFIER_TYPE: %s", notifierType)
	}

	if n == nil {
		fatal("Notifier could not be initialized")
	}

	http.HandleFunc("/", handler.AlertmanagerHandler(n))

	log.Printf("Starting server with '%s' notifier, listening on %s", notifierType, listenAddress)
	if err := http.ListenAndServe(listenAddress, nil); err != nil {
		fatalf("Error starting server: %v", err)
	}
}

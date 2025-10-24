package model

// TelegramSendMessagePayload is the payload for sending a message to Telegram.
type TelegramSendMessagePayload struct {
	ChatID    string `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

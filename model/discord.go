package model

// DiscordWebhookPayload is the payload for a Discord webhook.
type DiscordWebhookPayload struct {
	Content   string  `json:"content,omitempty"`
	Username  string  `json:"username,omitempty"`
	AvatarURL string  `json:"avatar_url,omitempty"`
	Embeds    []Embed `json:"embeds,omitempty"`
}

// Embed is a Discord embed object.
type Embed struct {
	Title       string       `json:"title,omitempty"`
	Description string       `json:"description,omitempty"`
	URL         string       `json:"url,omitempty"`
	Timestamp   string       `json:"timestamp,omitempty"`
	Color       int          `json:"color,omitempty"`
	Footer      *Footer      `json:"footer,omitempty"`
	Image       *Image       `json:"image,omitempty"`
	Thumbnail   *Thumbnail   `json:"thumbnail,omitempty"`
	Author      *Author      `json:"author,omitempty"`
	Fields      []EmbedField `json:"fields,omitempty"`
}

// Footer is a Discord embed footer.
type Footer struct {
	Text    string `json:"text,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// Image is a Discord embed image.
type Image struct {
	URL string `json:"url,omitempty"`
}

// Thumbnail is a Discord embed thumbnail.
type Thumbnail struct {
	URL string `json:"url,omitempty"`
}

// Author is a Discord embed author.
type Author struct {
	Name    string `json:"name,omitempty"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

// EmbedField is a Discord embed field.
type EmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}

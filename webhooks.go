package shopify

import "time"

// Webhook is a shopify webhook
type Webhook struct {
	// ID is the ID of the webhook.
	ID int64
	// Address is the destination to which the webhook subscription should send the POST request when an event occurs.
	Address string
	// Topic is the event that triggers the webhook.
	Topic string
	// CreatedAt is the date and time when the webhook subscription was created.
	CreatedAt time.Time
	// UpdatedAt is the date and time (ISO 8601 format) when the webhook subscription was last updated.
	UpdatedAt time.Time
}

// Webhooks is a collection of webhook
type Webhooks []Webhook

// WebhookRepository maintains the webhooks of a shop.
type WebhookRepository interface {
	// List retrieves a list of webhooks
	List() (Webhooks, error)
	// Create creates a new webhook
	Create(webhook Webhook) (Webhook, error)
}

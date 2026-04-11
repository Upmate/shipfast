package billing

import "time"

type WebhookEvent struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Payload   string    `json:"payload"`
	Processed bool      `json:"processed"`
	CreatedAt time.Time `json:"created_at"`
}

func NewWebhookEvent(eventType, payload string) *WebhookEvent {
	return &WebhookEvent{
		ID:        generateID(),
		Type:      eventType,
		Payload:   payload,
		Processed: false,
		CreatedAt: time.Now(),
	}
}

func (e *WebhookEvent) MarkProcessed() {
	e.Processed = true
}

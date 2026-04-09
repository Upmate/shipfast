package billing

import "time"

type AuditEntry struct {
	ID        string    `json:"id"`
	Action    string    `json:"action"`
	EntityID  string    `json:"entity_id"`
	UserID    string    `json:"user_id"`
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"created_at"`
}

func LogAudit(action, entityID, userID, details string) *AuditEntry {
	return &AuditEntry{
		ID:        generateID(),
		Action:    action,
		EntityID:  entityID,
		UserID:    userID,
		Details:   details,
		CreatedAt: time.Now(),
	}
}

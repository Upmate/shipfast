package billing

import "time"

type Discount struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Percent   int       `json:"percent"`
	MaxUses   int       `json:"max_uses"`
	UsedCount int       `json:"used_count"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (d *Discount) IsValid() bool {
	if d.UsedCount >= d.MaxUses {
		return false
	}
	if time.Now().After(d.ExpiresAt) {
		return false
	}
	return true
}

func (d *Discount) Apply(amount int64) int64 {
	discount := amount * int64(d.Percent) / 100
	return amount - discount
}

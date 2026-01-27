package model

import (
	"errors"
	"time"
)

const (
	BookingStatusConfirmed = "confirmed"
	BookingStatusPending   = "pending"
	BookingStatusCancel    = "cancelled"
)

type Booking struct {
	ID         int       `json:"id" db:"id"`
	UserID     int       `json:"user_id" db:"user_id"`
	ResourceID int       `json:"resource_id" db:"resource_id"`
	Status     string    `json:"status" db:"status"`
	BookedAt   time.Time `json:"booked_at" db:"booked_at"`
}

func (b *Booking) Validate() error {
	if b.UserID <= 0 {
		return errors.New("некорректный User_ID")
	}
	if b.ResourceID <= 0 {
		return errors.New("некорректный resource_ID")
	}

	if b.Status == "" {
		return errors.New("статус бронирование обязателен")
	}

	validStatus := map[string]bool{
		BookingStatusConfirmed: true,
		BookingStatusPending:   true,
		BookingStatusCancel:    true,
	}

	if !validStatus[b.Status] {
		return errors.New("статус должен быть:pending, confirmed, cancelled")
	}
	return nil
}

package model

import (
	"errors"
	"time"
)

const (
	ResourceTypeTable       = "table"
	ResourceTypeDoctor      = "doctor"
	ResourceTypeMeetingRoom = "meeting_room"
)

type Resource struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Type        string    `json:"type" db:"type"`
	IsAvailable bool      `json:"is_available" db:"is_available"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

func (r *Resource) Validate() error {
	if r.Name == "" {
		return errors.New("название ресурса обязательно")
	}

	if r.Type == "" {
		return errors.New("тип ресурса обязателен")
	}

	validTypes := map[string]bool{
		ResourceTypeMeetingRoom: true,
		ResourceTypeDoctor:      true,
		ResourceTypeTable:       true,
	}

	if !validTypes[r.Type] {
		return errors.New("тип должен быть: doctor, table, meeting_room")
	}

	return nil
}

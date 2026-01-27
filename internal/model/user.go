package model

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        int       `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Name      string    `json:"name" db:"name"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("email обязателен")
	}

	if !strings.Contains(u.Email, "@") || !strings.Contains(u.Email, ".") {
		return errors.New("невалидный формат email")
	}

	if u.Name == "" {
	}

	if u.Password == "" {
		return errors.New("пароль обязателен")
	}
	if len(u.Password) < 6 {
		return errors.New("пароль должен быть минимум на 6 символов")
	}

	return nil
}

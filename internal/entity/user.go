package entity

import "time"

type User struct {
	ID         string     `json:"id,omitempty"`
	Username   string     `json:"username,omitempty"`
	Email      string     `json:"email,omitempty"`
	Name       string     `json:"name,omitempty"`
	Password   string     `json:"password,omitempty"`
	Status     int8       `json:"-"`
	CreatedAt  time.Time  `json:"-"`
	EnabledAt  *time.Time `json:"-"`
	DisabledAt *time.Time `json:"-"`
}

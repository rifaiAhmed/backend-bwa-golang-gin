package user

import "time"

type User struct {
	ID             int
	Name           string
	Email          string
	Occupation     string
	PasswordHash   string
	AvatarFileName string
	Token          string
	Role           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

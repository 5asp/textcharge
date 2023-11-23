package app

import "time"

type AppUser struct {
	ID        int
	AppID     int
	UserID    int
	ServiceID int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b AppUser) Table() string {
	return "user_apps"
}

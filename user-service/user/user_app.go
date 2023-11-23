package user

import "time"

type UserApp struct {
	ID        int
	AppID     int
	UserID    int
	ServiceID int
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b UserApp) Table() string {
	return "user_apps"
}

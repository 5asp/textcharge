package user

import "time"

type UserApp struct {
	ID        int
	AppID     int
	UserID    int
	ServiceID int
	CreatedAt time.Time
	UpdatedAt time.Time
}

package user

import "time"

type UserApp struct {
	ID        int
	AppID     int
	UserID    int
	Quota     int
	ServiceID int
	CreatedAt time.Time
	UpdatedAt time.Time
}

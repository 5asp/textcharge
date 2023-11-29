package user

import "time"

type UserService struct {
	ID        int
	UserID    int
	ServiceID int
	Quota     int
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b UserService) Table() string {
	return "user_services"
}

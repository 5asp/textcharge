package user

import "time"

type User struct {
	ID           int
	Account      string
	Password     string
	Status       int
	UserApps     []UserApp     `ref:"id" fk:"user_id"`
	UserServices []UserService `ref:"id" fk:"user_id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (b User) Table() string {
	return "users"
}

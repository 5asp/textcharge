package app

import "time"

type App struct {
	ID        int
	AppID     int
	Secret    string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

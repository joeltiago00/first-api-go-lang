package resources

import "time"

type UserResource struct {
	Data UserData `json:"data"`
}

type UserData struct {
	ID        int       `json:"ID"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
	DeletedAt time.Time `json:"DeletedAt"`
}

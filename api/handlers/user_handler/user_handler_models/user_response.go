package user_handler_models

import "time"

// UserResponse - API model for responding with user data.
type UserResponse struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package user_handler_models

// CreateUserRequest - API model for creating a new user.
type CreateUserRequest struct {
	Username  string `json:"username" validate:"required,min=3,max=50"`
	FirstName string `json:"first_name" validate:"max=50"`
	LastName  string `json:"last_name" validate:"max=50"`
}

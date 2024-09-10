package handler_models

// UpdateUserRequest - API model for updating an existing user.
type UpdateUserRequest struct {
	Username  string `json:"username" validate:"required,min=3,max=50"`
	FirstName string `json:"first_name" validate:"max=50"`
	LastName  string `json:"last_name" validate:"max=50"`
}

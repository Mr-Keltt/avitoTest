package user_models

type UserUpdateModel struct {
	ID        int    `json:"id" validate:"required"`
	Username  string `json:"username" validate:"required,min=3,max=50"`
	FirstName string `json:"first_name" validate:"max=50"`
	LastName  string `json:"last_name" validate:"max=50"`
}

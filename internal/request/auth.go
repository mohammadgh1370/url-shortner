package request

type UserSignUpRequest struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Username  string `json:"username" validate:"required,min=3,max=32"`
	Password  string `json:"password" validate:"required,min=5"`
}

type UserSignInRequest struct {
	Username string `json:"username" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=5"`
}

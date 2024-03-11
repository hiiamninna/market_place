package collections

type InputUserRegister struct {
	ID       string
	Name     string `json:"name" validate:"required;min:5;max:15"`
	Username string `json:"username" validate:"required;min:5;max:50"`
	Password string `json:"password" validate:"required;min:5;max:15"`
}

type UserLoginInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

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

type ProductInput struct {
	ID             string
	Name           string   `json:"name"`
	Price          int      `json:"price"`
	ImageUrl       string   `json:"imageUrl"`
	Stock          int      `json:"stock"`
	Condition      string   `json:"condition"`
	Tags           []string `json:"tags"`
	IsPurchaseable bool     `json:"isPurchaseable"`
}

type ProductStockInput struct {
	ID    string
	Stock int `json:"stock"`
}

package collections

type UserRegisterAndLogin struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	AccessToken string `json:"accessToken"`
}

type User struct {
	ID       string
	Name     string
	Username string
	Password string
}

type Product struct {
	ID             string
	Name           string
	Price          int
	ImageUrl       string
	Stock          int
	Condition      string
	Tags           []string
	IsPurchaseable bool
}

package collections

type UserRegisterAndLogin struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	AccessToken string `json:"accessToken"`
}

type User struct {
	ID       string `json:"userId"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"accessToken"`
}

type Product struct {
	ID             string
	Name           string   `json:"name"`
	Price          int      `json:"price"`
	ImageUrl       string   `json:"imageUrl"`
	Stock          int      `json:"stock"`
	Condition      string   `json:"condition"`
	Tags           []string `json:"tags"`
	IsPurchaseable bool     `json:"isPurchaseable"`
}

type FileUpload struct {
	ImageUrl string `json:"imageUrl"`
}

type BankAccount struct {
	ID                string `json:"bankAccountId"`
	BankName          string `json:"bankName"`
	BankAccountName   string `json:"bankAccountName"`
	BankAccountNumber string `json:"bankAccountNumber"`
}

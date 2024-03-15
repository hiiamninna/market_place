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
	ID             string   `json:"productId"`
	UserID         string   `json:"-"`
	Name           string   `json:"name"`
	Price          int      `json:"price"`
	ImageUrl       string   `json:"imageUrl"`
	Stock          int      `json:"stock"`
	Condition      string   `json:"condition"`
	Tags           []string `json:"tags"`
	IsPurchaseable bool     `json:"isPurchaseable"`
	PurchaseCount  int      `json:"purchaseCount"`
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

type Seller struct {
	Name             string        `json:"name"`
	ProductSoldTotal int           `json:"productSoldTotal"`
	BankAccounts     []BankAccount `json:"bankAccounts"`
}

type ProductDetail struct {
	Product Product `json:"product"`
	Seller  Seller  `json:"seller"`
}

type ProductList struct {
	ProductID      string   `json:"productId"`
	Name           string   `json:"name"`
	Price          int      `json:"price"`
	ImageURL       string   `json:"imageUrl"`
	Stock          int      `json:"stock"`
	Condition      string   `json:"condition"`
	Tags           []string `json:"tags"`
	IsPurchaseable bool     `json:"isPurchaseable"`
	PurchaseCount  int      `json:"purchaseCount"`
}

type Meta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

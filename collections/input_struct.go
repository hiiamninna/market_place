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
	UserID         string
	Name           string   `json:"name" validate:"required;min:5;max:60"`
	Price          int      `json:"price" validate:"min:0"`
	ImageUrl       string   `json:"imageUrl" validate:"url"`
	Stock          int      `json:"stock" validate:"min:0"`
	Condition      string   `json:"condition" validate:"required;enum:condition"`
	Tags           []string `json:"tags"`
	IsPurchaseable bool     `json:"isPurchaseable"`
}

type ProductStockInput struct {
	ID    string
	Stock int `json:"stock" validate:"min:0"`
}

type BankAccountInput struct {
	ID                string
	UserID            string
	BankName          string `json:"bankName" validate:"required;min:5;max:15"`
	BankAccountName   string `json:"bankAccountName" validate:"required;min:5;max:15"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required;min:5;max:15"`
}

type PaymentInput struct {
	UserID        string
	ProductID     string
	BankAccountID string `json:"bankAccountId" validate:"required"`
	PaymentProof  string `json:"paymentProofImageUrl" validate:"required;url"`
	Quantity      int    `json:"quantity" validate:"min:1"`
	TotalPayment  int
}

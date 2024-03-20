package collections

type InputUserRegister struct {
	ID       string
	Name     string `json:"name" validate:"required,min=5,max=15"`
	Username string `json:"username" validate:"required,min=5,max=50"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

type UserLoginInput struct {
	Username string `json:"username" validate:"required,min=5,max=15"`
	Password string `json:"password" validate:"required,min=5,max=15"`
}

type ProductInput struct {
	ID             string
	UserID         string
	Name           *string   `json:"name" validate:"required,min=5,max=60"`
	Price          *int      `json:"price" validate:"required,min=0"`
	ImageUrl       *string   `json:"imageUrl" validate:"required,url"`
	Stock          *int      `json:"stock" validate:"required,min=0"`
	Condition      *string   `json:"condition" validate:"required"`
	Tags           *[]string `json:"tags" validate:"required"`
	IsPurchaseable *bool     `json:"isPurchaseable" validate:"required"`
}

type ProductPageInput struct {
	UserID         string
	UserOnly       bool     `query:"userOnly"`
	Limit          int      `query:"limit"`
	Offset         int      `query:"offset"`
	Tags           []string `query:"tags"`
	Condition      string   `query:"condition"`
	ShowEmptyStock bool     `query:"showEmptyStock"`
	MaxPrice       int      `query:"maxPrice"`
	MinPrice       int      `query:"minPrice"`
	SortBy         string   `query:"sortBy"`
	OrderBy        string   `query:"orderBy"`
	Search         string   `query:"search"`
}

type ProductStockInput struct {
	ID    string
	Stock int `json:"stock" validate:"required,min=0"`
}

type BankAccountInput struct {
	ID                string
	UserID            string
	BankName          string `json:"bankName" validate:"required,min=5,max=15"`
	BankAccountName   string `json:"bankAccountName" validate:"required,min=5,max=15"`
	BankAccountNumber string `json:"bankAccountNumber" validate:"required,min=5,max=15"`
}

type PaymentInput struct {
	UserID        string
	ProductID     string
	BankAccountID string `json:"bankAccountId" validate:"required"`
	PaymentProof  string `json:"paymentProofImageUrl" validate:"required,url"` //TODO : must a url
	Quantity      int    `json:"quantity" validate:"required,min=1"`
	TotalPayment  int
}

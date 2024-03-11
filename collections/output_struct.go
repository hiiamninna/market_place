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

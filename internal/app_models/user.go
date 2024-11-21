package appmodels

type User struct {
	Id       int64  `json:"-" db:"id"`
	Username string `json:"username"`
	SignInInput
}

type SignInInput struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

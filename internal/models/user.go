package models

type User struct {
	Number   string `json:"email,omitempty" db:"phone_number"`
	Password string `json:"password,omitempty" db:"password"`
}

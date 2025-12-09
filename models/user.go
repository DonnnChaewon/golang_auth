package models

type User struct {
	ID int `gorm:"primaryKey"`
	Name, Email, Password string
}

type Register struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type MyProfile struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}
package entity

type User struct {
	Username string `json:"username"`
	Password string `json:"password"` // Hashed password in storage
}

package models

type NewUser struct {
	UserID   string `json:"userid"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Token  string `json:"token"`
	Status int    `json:"status"`
}

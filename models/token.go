package models

type Token struct {
	Role        string `json:"role"`
	UserId      string `json:"userId"`
	TokenString string `json:"token"`
}

package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(userID, role string) (string, error) {
	var secretkey = "e2922901-f374-4283-a1ad-0e3c6d06011f"
	var mySigningKey = []byte(secretkey)

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userID"] = userID
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

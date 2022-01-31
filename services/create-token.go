package services

import (
	"go-simple-crud-webservice/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(userID, role string) (*models.TokenDetails, error) {
	var secretkey = os.Getenv("ACCESS_SECRET")
	var mySigningKey = []byte(secretkey)
	var err error

	td := &models.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 30).Unix()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["userID"] = userID
	atClaims["role"] = role
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(mySigningKey))
	if err != nil {
		return nil, err
	}

	return td, nil
}

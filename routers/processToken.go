package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pabloelisseo/twitt3r/bd"
	"github.com/pabloelisseo/twitt3r/models"
)

var Email string
var UserID string

func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("incredible_key")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token")
	}

	token = strings.TrimSpace(splitToken[1])

	jwtToken, err := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, found, _ := bd.CheckUserExists(claims.Email)
		if found {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, found, UserID, nil
	}
	if !jwtToken.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}

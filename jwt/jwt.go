package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pabloelisseo/twitt3r/models"
)

func GenerateJWT(t models.User) (string, error) {
	privateKey := []byte("incredible_key")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastName":  t.LastName,
		"birthdate": t.Birthdate,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.Website,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(privateKey)
	return tokenStr, err
}

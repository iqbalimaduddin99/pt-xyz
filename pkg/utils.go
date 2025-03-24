package pkg

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var secretKey = []byte(os.Getenv("SECRET_KEY_TOKEN"))

type Claims struct {
	ID       uuid.UUID    	`json:"id"`
	UserName string 		`json:"user_name"`
	IsAdmin  bool 		 `json:"is_admin"`
	jwt.StandardClaims
}

func GenerateJWT(id uuid.UUID, userName string, isAdmin bool) (string, error) {
	claims := Claims{
		ID:       id,
		UserName: userName,
		IsAdmin:     isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("Error generating token: %v", err)
	}
	return tokenString, nil
}


func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

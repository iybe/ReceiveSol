package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id string, durationSeconds int, keySecret string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(durationSeconds * int(time.Second)))
	claims := jwt.MapClaims{
		"id":  id,
		"exp": expirationTime.Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(keySecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(token, id, keySecret string) error {
	tokenActual, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return []byte(keySecret), nil
	})
	if err != nil || !tokenActual.Valid {
		return fmt.Errorf("token invalido")
	}
	if !tokenActual.Valid {
		return fmt.Errorf("token expirado")
	}

	claims, ok := tokenActual.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("ocorreu um erro ao verificar o token")
	}
	idActual := claims["id"].(string)
	if idActual != id {
		return fmt.Errorf("token invalido")
	}

	return nil
}

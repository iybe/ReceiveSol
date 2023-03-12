package services

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	// Parâmetros de entrada
	id := "abc123"
	durationSeconds := 3600
	secret := "my-secret"

	// Chamada da função a ser testada
	token, err := CreateToken(id, durationSeconds, secret)

	// Verificação dos resultados
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)

	claims := parsedToken.Claims.(jwt.MapClaims)
	assert.Equal(t, id, claims["id"])
	assert.InDelta(t, time.Now().Add(time.Duration(durationSeconds)).Unix(), claims["exp"], 1.0)
}

func TestVerifyToken(t *testing.T) {
	token, err := CreateToken("123", 60, "chave-secreta")
	if err != nil {
		t.Errorf("Erro ao criar token: %v", err)
	}

	err = VerifyToken(token, "123", "chave-secreta")
	if err != nil {
		t.Errorf("Erro ao verificar token: %v", err)
	}
}

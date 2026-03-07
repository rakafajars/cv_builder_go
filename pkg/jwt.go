package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, secretKey string) (string, error) {

	// isi data token secara terstruktur
	claims := JWTCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "cv-builder-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// tanda tanggal token dengan secret key
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateToken(tokenString string, secretKey string) (*JWTCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTCustomClaims)
	if !ok || !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return claims, nil
}

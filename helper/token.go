package helper

import (
	"Trainify/model"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Email  string `json:"email"`
	UserId uint   `json:"userId"`
	jwt.RegisteredClaims
}

var secret string = "secret"

// GenerateToken generates a JWT for a given user
func GenerateToken(user model.User) (string, error) {
	claims := CustomClaims{
		Email:  user.Email,
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * 20)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Println("Error in token signing.")
		return "", err
	}

	return t, nil
}

// ValidateToken validates a given JWT and extracts the claims
func ValidateToken(clientToken string) (*CustomClaims, error) {
    token, err := jwt.ParseWithClaims(clientToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        // Validate the signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(secret), nil
    })

    if err != nil {
        if errors.Is(err, jwt.ErrTokenMalformed) {
            return nil, errors.New("malformed token")
        } else if errors.Is(err, jwt.ErrTokenExpired) {
            return nil, errors.New("token has expired")
        } else if errors.Is(err, jwt.ErrTokenNotValidYet) {
            return nil, errors.New("token not valid yet")
        }
        return nil, err
    }

    claims, ok := token.Claims.(*CustomClaims)
    if !ok || !token.Valid {
        return nil, errors.New("invalid token claims")
    }

    return claims, nil
}

func RefreshToken(claims *CustomClaims) (string, error) {
    // Create a new token with extended expiration
    newClaims := CustomClaims{
        Email:  claims.Email,
        UserId: claims.UserId,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 24)),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
    return token.SignedString([]byte(secret))
}

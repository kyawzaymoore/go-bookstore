package common

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/kzm/go-bookstore/api/config"
)

func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Validate the signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        // Use the secret key from the config
        secretKey := []byte(config.AppConfig.App.SecretKey)
        return secretKey, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    } else {
        return nil, fmt.Errorf("invalid token")
    }
}
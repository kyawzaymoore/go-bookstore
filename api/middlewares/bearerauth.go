package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/kzm/go-bookstore/api/common"
	"github.com/kzm/go-bookstore/api/config"
	"github.com/kzm/go-bookstore/api/repository"
)

// BearerTokenMiddleware checks for a valid Bearer token and uses the UserRepository.
func BearerTokenMiddleware(userRepo repository.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Missing authorization header")
				return
			}

			if !strings.HasPrefix(tokenString, "Bearer ") {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Invalid authorization header format")
				return
			}

			tokenString = tokenString[len("Bearer "):]

			err := verifyToken(tokenString)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "Invalid token")
				return
			}

			claims, err := common.GetClaimsFromToken(tokenString)
			if err != nil {
				fmt.Printf("Error parsing token: %v", err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// Access specific claim values
			userid := claims["userid"].(string)
			uuid := claims["uuid"].(string)

			// Use the UserRepository to check the session
			sessionResult := userRepo.CheckSession(userid, uuid)

			if !sessionResult {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			// If token is valid and session check passed, proceed to the next handler
			next.ServeHTTP(w, r)
		})
	}
}

func verifyToken(tokenString string) error {
	secretKey := []byte(config.AppConfig.App.SecretKey)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

package handler

import (
	"context"
	"net/http"
	"strings"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("secret-key")

func BaseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Отсутствует заголовок Authorization", http.StatusUnauthorized)
			return
		}
		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			http.Error(w, "Неверный формат заголовка Authorization", http.StatusUnauthorized)
			return
		}
		tokenString := authParts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return secretKey, nil
		})
		if err != nil {
			http.Error(w, "Неверный токен", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Неверный токен", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Неверный формат токена", http.StatusUnauthorized)
			return
		}

		id, ok := claims["id"].(string)
		if !ok {
			http.Error(w, "Отсутствует идентификатор пользователя в токене", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "id", id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

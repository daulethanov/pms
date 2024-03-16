package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"errors"
)

var secretKey = []byte("secret-key")

func JWTGenearate(id, email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() 

	tokenString, err := token.SignedString(secretKey) 
	if err != nil {
		return "", err
	}
	return tokenString, nil 
}


func JWTGenerateRefreshToken(id,email string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
    claims["email"] = email
    claims["exp"] = time.Now().Add(time.Hour * 24 * 30).Unix() 

    refreshTokenString, err := token.SignedString(secretKey) 
    if err != nil {
        return "", err
    }

    return refreshTokenString, nil 
}


func RefreshAccessToken(refreshToken string) (string, error) {
    token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
        return secretKey, nil
    })
    if err != nil {
        return "", err
    }
    
    if !token.Valid || token.Method != jwt.SigningMethodHS256 {
        return "", errors.New("неверный refresh token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", errors.New("неверный формат токена")
    }
    email, ok := claims["email"].(string)
	id, ok := claims["id"].(string)
	if !ok {
        return "", errors.New("неверный формат токена")
    }

    newAccessToken, err := JWTGenearate(id,email)
    if err != nil {
        return "", err
    }

    return newAccessToken, nil
}
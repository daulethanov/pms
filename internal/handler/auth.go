package handler

import (
	"encoding/json"
	"net/http"
	"todo/internal/model/schema"
	"todo/internal/service"
	"todo/pkg/jwt"
)

type AuthHandlerInterface interface {
	SignUp(w http.ResponseWriter, r *http.Request)
	SignIn(w http.ResponseWriter, r *http.Request)
	NewJwtToken(w http.ResponseWriter, r *http.Request)
}

type AuthHandler struct {
	service service.AuthServiceInterface
}

func NewAuthHandler(service service.AuthServiceInterface) AuthHandlerInterface {
	return &AuthHandler{
		service: service,
	}
}

func (a *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	body := schema.CreateUserSchema{}
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
		return
	}
	if err = a.service.EmailValidate(body.Email); err != nil {
		http.Error(w, "Ошибка валидации Email", http.StatusBadRequest)
		return
	}
	id, err := a.service.CreateUser(body.Email, body.Password); 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	access, err := jwt.JWTGenearate(id, body.Email)
	if err != nil {
		http.Error(w, "Ошибка генерации jwt", http.StatusBadRequest)
		return
	}
	refresh, err := jwt.JWTGenerateRefreshToken(id, body.Email)
	if err != nil {
		http.Error(w, "Ошибка генерации jwt", http.StatusBadRequest)
		return
	}

	response := map[string]string{
		"access_token": access,
		"refresh_token": refresh,

	}
	responseData, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Ошибка сериализации данных", http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "application/json")
    
    w.WriteHeader(http.StatusOK)
    
    _, err = w.Write(responseData)
    if err != nil {
        http.Error(w, "Ошибка отправки данных", http.StatusInternalServerError)
        return
    }
}


func (a *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	body := schema.LoginUserSchema{}
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
		return
	}
	if err = a.service.EmailValidate(body.Email); err != nil {
		http.Error(w, "Ошибка валидации Email", http.StatusBadRequest)
		return
	}
	user_id, err := a.service.LoginUser(body.Email, body.Password)
	if err != nil{
		http.Error(w, "Не верный логин или пароль", http.StatusUnauthorized)
	}
	
	access, err := jwt.JWTGenearate(user_id, body.Email)
	if err != nil {
		http.Error(w, "Ошибка генерации jwt", http.StatusBadRequest)
		return
	}
	refresh, err := jwt.JWTGenerateRefreshToken(user_id, body.Email)
	if err != nil {
		http.Error(w, "Ошибка генерации jwt", http.StatusBadRequest)
		return
	}
	response := map[string]string{
		"access_token": access,
		"refresh_token": refresh,
	}
	responseData, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Ошибка сериализации данных", http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "application/json")
    
    w.WriteHeader(http.StatusOK)
    
    _, err = w.Write(responseData)
    if err != nil {
        http.Error(w, "Ошибка отправки данных", http.StatusInternalServerError)
        return
    }
}


func (a *AuthHandler) NewJwtToken(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	body := schema.NewTokenSchema{} 
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
		return
	}
	access, err := jwt.RefreshAccessToken(body.RefreshToken)
	if err != nil{
		http.Error(w, "Токен устарел", http.StatusUnauthorized)
        return
	}
	response := map[string]string{
		"access_token": access,
	}
	responseData, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Ошибка сериализации данных", http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "application/json")
    
    w.WriteHeader(http.StatusOK)
    
    _, err = w.Write(responseData)
    if err != nil {
        http.Error(w, "Ошибка отправки данных", http.StatusInternalServerError)
        return
    }
}

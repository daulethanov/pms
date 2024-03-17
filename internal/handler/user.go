package handler

import (
	"encoding/json"
	"net/http"
	"todo/internal/service"
)

type ProfileHandlerInterface interface{
	ViewProfile(w http.ResponseWriter, r *http.Request)
	EditImageProfile(w http.ResponseWriter, r *http.Request)
}

type ProfileHandler struct{
	servie service.AuthServiceInterface
}

func NewProfileHandler(servie service.AuthServiceInterface) ProfileHandlerInterface{
	return &ProfileHandler{
		servie: servie,
	}
}

func (a *ProfileHandler)ViewProfile(w http.ResponseWriter, r *http.Request){
	userID := r.Context().Value("id").(string)
	user, err := a.servie.ViewProfile(userID)
	if err != nil{ 
		http.Error(w, "Пользователь не найден", http.StatusUnauthorized)

	}
	response := map[string]any{
		"id": userID,
		"email":  user.Email,
		"username": user.Username,
		"created_at": user.CretedAt,
	}
	responseData, err := json.Marshal(response)
	if err != nil {
        http.Error(w, "Ошибка сериализации данных", http.StatusInternalServerError)
        return
    }
	w.WriteHeader(http.StatusOK)
	w.Write(responseData)
}

func (a *ProfileHandler)EditImageProfile(w http.ResponseWriter, r *http.Request){
	err := r.ParseMultipartForm(20 << 30) 
	if err != nil {
        http.Error(w, "Ошибка парсинга формы", http.StatusInternalServerError)
        return
    }
	
}
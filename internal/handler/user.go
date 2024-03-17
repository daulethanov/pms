package handler

import (
	"encoding/json"
	"net/http"
	"todo/internal/service"
	"log"
)

type ProfileHandlerInterface interface {
	ViewProfile(w http.ResponseWriter, r *http.Request)
	EditImageProfile(w http.ResponseWriter, r *http.Request)
}

type ProfileHandler struct {
	servie service.AuthServiceInterface
}

func NewProfileHandler(servie service.AuthServiceInterface) ProfileHandlerInterface {
	return &ProfileHandler{
		servie: servie,
	}
}

func (a *ProfileHandler) ViewProfile(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("id").(string)
	user, err := a.servie.ViewProfile(userID)
	if err != nil {
		http.Error(w, "Пользователь не найден", http.StatusUnauthorized)

	}
	response := map[string]any{
		"id":         userID,
		"email":      user.Email,
		"username":   user.Username,
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

func (a *ProfileHandler) EditImageProfile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(20 << 30)
	userID := r.Context().Value("id").(string)
	if err != nil {
		http.Error(w, "Пользователь не найден", http.StatusUnauthorized)

	}
	if err != nil {
		http.Error(w, "Ошибка парсинга формы", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Пользователь не найден", http.StatusUnauthorized)
	}
	defer file.Close()
	if err = a.servie.EditImageProfile(userID, file, handler.Size); err != nil {
		log.Fatalln(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	w.WriteHeader(http.StatusOK)
}

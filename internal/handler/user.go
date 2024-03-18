package handler

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
	"todo/internal/model"
	"todo/internal/model/schema"
	"todo/internal/service"
	mail "todo/pkg/smtp"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ProfileHandlerInterface interface {
	ViewProfile(w http.ResponseWriter, r *http.Request)
	EditImageProfile(w http.ResponseWriter, r *http.Request)
	EditPassword(w http.ResponseWriter, r *http.Request)
	EditPasswordCode(w http.ResponseWriter, r *http.Request)
	EditPasswordConfirm(w http.ResponseWriter, r *http.Request)
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
	err := r.ParseMultipartForm(20 << 20)
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


func generateEditPasswordCode()int{
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(900000) + 100000
	return randomNumber

}

var messageCodeEditPassword map[int]string
var urlEditPassword map[string]string

func (p *ProfileHandler)EditPassword(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	body := schema.EditPasswordSchema{}

	err := decoder.Decode(&body)
	if err != nil{ 
		http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
		return
	}
	if err := model.EmailValidate(body.Email); err != nil {
		http.Error(w, "Ошибка валидации Email", http.StatusBadRequest)
		return
	}
	
    code := generateEditPasswordCode()
    if err := p.servie.EditPassword(body.Email); err != nil{
		http.Error(w, "Ошибка email", http.StatusBadRequest)

	}
    if err = mail.SendMessageEditPassword(body.Email, code); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	messageCodeEditPassword[code] = body.Email

	
	w.WriteHeader(http.StatusOK)
} 

func (a *ProfileHandler) EditPasswordCode(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	body := schema.EditPasswordCodeSchema{}

	err := decoder.Decode(&body)
	if err != nil{ 
		http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
		return
	}
	
	email, ok := messageCodeEditPassword[body.Code]
	urlConfirmEditPassword := uuid.New().String()
	urlEditPassword[urlConfirmEditPassword] = email
    if !ok {
        http.Error(w, "Неверный код подтверждения", http.StatusBadRequest)
        return
    }
	delete(messageCodeEditPassword, body.Code)
	

	w.WriteHeader(http.StatusOK)

}

func (a *ProfileHandler) EditPasswordConfirm(w http.ResponseWriter, r *http.Request){
	
	vars := mux.Vars(r)
	slug := vars["slug"]
	decoder := json.NewDecoder(r.Body)
	body := schema.EditPasswordConfirmSchema{}

	err := decoder.Decode(&body)
	if err != nil{ 
		http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
		return
	}
	email, ok := urlEditPassword[slug]
	if !ok {
        http.Error(w, "Неверный код подтверждения", http.StatusBadRequest)
        return
    }
	if err := a.servie.ConfirmEditPassword(email, body.Password);err != nil{
		http.Error(w, "не удалось изменит пароль", http.StatusBadRequest)
        return
	}
	w.WriteHeader(http.StatusOK)
}
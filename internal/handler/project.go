package handler

import (
	"encoding/json"
	"net/http"
	"todo/internal/model/schema"
	"todo/internal/service"
)

type ProjectHandlerInterface interface{
	CreateProjectRoom(w http.ResponseWriter, r *http.Request)
}

type ProjectHandler struct{
	service service.ProjectServiceInterface
}

func NewProjectHandler(service service.ProjectServiceInterface)ProjectHandlerInterface{
	return &ProjectHandler{
		service: service,
	}
}

func (p *ProjectHandler) CreateProjectRoom(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	body := schema.CreateProjectSchema{}
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
		return
	}
	if err := p.service.CreateProject(&body);err != nil{
		http.Error(w, "Ошибка при создание проекта", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}


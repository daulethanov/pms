package handler

import (
	"encoding/json"
	"net/http"
	"todo/internal/model/schema"
	"todo/internal/service"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type ProjectHandlerInterface interface {
	TaskEdit(w http.ResponseWriter, r *http.Request)
	CreateProjectRoom(w http.ResponseWriter, r *http.Request)
	DetailProjectView(w http.ResponseWriter, r *http.Request)
}

type ProjectHandler struct {
	service service.ProjectServiceInterface
}

func NewProjectHandler(service service.ProjectServiceInterface) ProjectHandlerInterface {
	return &ProjectHandler{
		service: service,
	}
}

func (p *ProjectHandler) CreateProjectRoom(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	userID := r.Context().Value("id").(string)
	body := schema.CreateProjectSchema{}
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
		return
	}
	if err := p.service.CreateProject(&body, userID); err != nil {
		http.Error(w, "Ошибка при создание проекта", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (p *ProjectHandler) DetailProjectView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	projectID := vars["id"]
	project, err := p.service.DetailProject(projectID)
	if err != nil {
		http.Error(w, "Ошибка при просмотре проекта", http.StatusBadRequest)
		return
	}
	response := map[string]string{
		"name": project.Name,
		// "general": project.General,
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

func (p *ProjectHandler) TaskEdit(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Ошибка отправки данных", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

}

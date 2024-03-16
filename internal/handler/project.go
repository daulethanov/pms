package handler

import (
	"encoding/json"
	"net/http"
	"todo/internal/model/schema"
)

type ProjectHandlerInterface interface{
	CreateProjectRoom(w http.ResponseWriter, r *http.Request)
}

type ProjectHandler struct{

}

func (p *ProjectHandler) CreateProjectRoom(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	body := schema.CreateProjectSchema{}
	err := decoder.Decode(&body)
	if err != nil {
		http.Error(w, "Ошибка разбора JSON", http.StatusBadRequest)
		return
	}
	
}


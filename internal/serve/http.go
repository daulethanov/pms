package serve

import (
	"net/http"
	"todo/internal/handler"
	"todo/internal/service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)



func HttpServer(userCollection, projectCollection *mongo.Collection) error {
	router := mux.NewRouter()
	authService := service.NewAuthService(userCollection)
	authHandler := handler.NewAuthHandler(authService)
	
	projectService := service.NewProjectService(projectCollection)
	projectHandler := handler.NewProjectHandler(projectService)

	api := router.PathPrefix("/api").Subrouter()
	{
		auth := api.PathPrefix("/auth").Subrouter()
		{
			auth.HandleFunc("/sign-up", authHandler.SignUp).Methods("POST")
			auth.HandleFunc("/sign-in", authHandler.SignIn).Methods("POST")
			auth.HandleFunc("/refresh", authHandler.NewJwtToken).Methods("POST")
		}

		project := api.PathPrefix("/project").Subrouter()
		{
			project.HandleFunc("/create", projectHandler.CreateProjectRoom).Methods("POST")
		}
	}

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		return err
	}
	return nil
}

package serve

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"todo/internal/handler"
	"todo/internal/service"
)

func HttpServer(userCollection, projectCollection *mongo.Collection) error {
	router := mux.NewRouter()
	authService := service.NewAuthService(userCollection)
	authHandler := handler.NewAuthHandler(authService)

	projectService := service.NewProjectService(projectCollection)
	projectHandler := handler.NewProjectHandler(projectService)


	roomHandler := handler.NewRoomHandler()

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
			project.Handle("/create", handler.BaseMiddleware(
				http.HandlerFunc(projectHandler.CreateProjectRoom))).Methods("POST")
			project.Handle("/detail/{id}", handler.BaseMiddleware(
				http.HandlerFunc(projectHandler.DetailProjectView))).Methods("GET")
		}
		room := api.PathPrefix("/room").Subrouter()
		{
			room.HandleFunc("", roomHandler.RoomMessage)
		}
	}

	err := http.ListenAndServe(":8000", router)
	if err != nil {
		return err
	}
	return nil
}

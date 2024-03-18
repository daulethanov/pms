package serve

import (
	"net/http"
	"todo/internal/handler"
	"todo/internal/service"
	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/mongo"
)

func HttpServer(userCollection, projectCollection, taskLevelCollection *mongo.Collection, minio *minio.Client) error {
	router := mux.NewRouter()
	authService := service.NewAuthService(userCollection, minio)
	authHandler := handler.NewAuthHandler(authService)

	projectService := service.NewProjectService(projectCollection, taskLevelCollection)
	projectHandler := handler.NewProjectHandler(projectService)


	profileHandler := handler.NewProfileHandler(authService)
	roomHandler := handler.NewRoomHandler()

	api := router.PathPrefix("/api").Subrouter()
	{
		auth := api.PathPrefix("/auth").Subrouter()
		{
			auth.HandleFunc("/sign-up", authHandler.SignUp).Methods("POST")
			auth.HandleFunc("/sign-in", authHandler.SignIn).Methods("POST")
			auth.HandleFunc("/refresh", authHandler.NewJwtToken).Methods("POST")
			auth.HandleFunc("/edit/password", profileHandler.EditPassword).Methods("POST")
			auth.HandleFunc("/edit/password/code", profileHandler.EditPasswordCode).Methods("POST")
			auth.HandleFunc("/edit/password/confirm/{slug}", profileHandler.EditPasswordConfirm).Methods("POST")
		}
		{
			auth.Handle("/profile", handler.BaseMiddleware(
				http.HandlerFunc(profileHandler.ViewProfile))).Methods("GET")
			auth.Handle("/profile/edit/image", handler.BaseMiddleware(
				http.HandlerFunc(profileHandler.EditImageProfile))).Methods("POST")

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

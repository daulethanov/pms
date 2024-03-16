package serve

import (
	"net/http"
	"todo/internal/handler"
	"todo/internal/service"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)



func HttpServer(userCollection *mongo.Collection) error {
	router := mux.NewRouter()
	authService := service.NewAuthService(userCollection)
	authHandler := handler.NewAuthHandler(authService)

	api := router.PathPrefix("/api").Subrouter()
	{
		auth := api.PathPrefix("/auth").Subrouter()
		{
			auth.HandleFunc("/sign-up", authHandler.SignUp).Methods("POST")
			auth.HandleFunc("/sign-in", authHandler.SignIn).Methods("POST")
			auth.HandleFunc("/refresh", authHandler.NewJwtToken).Methods("POST")
		}
	}
	
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		return err
	}
	return nil
}

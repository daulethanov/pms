package main

import (
	"log"
	"todo/config"
	"todo/internal/serve"
	"todo/pkg/mongo"
)

func init(){
	err := config.GoYamlParse()
	if err != nil {
		log.Fatal("Error parse YAML", err)
	}
}


func main() {
	mongo, err := mongo.NewMongoClient()
	if err != nil{
		log.Fatal("Error connect mongo")
	}

	users_collection := mongo.Database("db").Collection("users")
	if err = serve.HttpServer(users_collection); err != nil {
		log.Fatal("Error start http server")
	}
	
}
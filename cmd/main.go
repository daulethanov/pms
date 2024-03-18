package main

import (
	"log"
	"todo/config"
	"todo/internal/serve"
	"todo/pkg/minio"
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
	minioClient, err := minio.NewMinioClient()
	if err != nil{
		log.Fatal("Error connect minio")
	}
	projects_collection := mongo.Database("db").Collection("projects")
	users_collection := mongo.Database("db").Collection("users")
	task_level_collection := mongo.Database("db").Collection("task_level")
	if err = serve.HttpServer(users_collection, projects_collection,task_level_collection, minioClient); err != nil {
		log.Fatal("Error start http server")
	}
	
}
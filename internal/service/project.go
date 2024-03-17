package service

import (
	"context"
	"time"
	"todo/internal/model"
	"todo/internal/model/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectServiceInterface interface {
	CreateProject(schema *schema.CreateProjectSchema, userID string) error
	DetailProject(id string)(*model.Project, error)
}


type ProjectService struct {
	projectCollection *mongo.Collection
}


func NewProjectService(projectCollection *mongo.Collection) ProjectServiceInterface {
	return &ProjectService{
		projectCollection: projectCollection,
	}
}


func (p *ProjectService) CreateProject(body *schema.CreateProjectSchema, userID string)error{
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	project := model.Project{
		ID: primitive.NewObjectID(),
		CretedAt: time.Now(),
		Name: body.Name,
		General: body.General,
		UserID: userObjectID,
	}
	_, err = p.projectCollection.InsertOne(context.TODO(), project)
	if err != nil{
		return err
	}
	return nil
}

func (p *ProjectService) DetailProject(id string)(*model.Project, error){
	var project model.Project
	projectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
    err = p.projectCollection.FindOne(context.TODO(), bson.M{"_id": projectID}).Decode(&project)
	if err != nil{
		return nil, err
	}
	return &project, nil
}

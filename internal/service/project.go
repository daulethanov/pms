package service

import (
	"context"
	"time"
	"todo/internal/model"
	"todo/internal/model/schema"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectServiceInterface interface {
	CreateProject(*schema.CreateProjectSchema) error
}


type ProjectService struct {
	projectCollection *mongo.Collection
}


func NewProjectService(projectCollection *mongo.Collection) ProjectServiceInterface {
	return &ProjectService{
		projectCollection: projectCollection,
	}
}


func (p *ProjectService) CreateProject(body *schema.CreateProjectSchema)error{
	project := model.Project{
		ID: primitive.NewObjectID(),
		CretedAt: time.Now(),
		Name: body.Name,
		General: body.General,
		
	}

	_, err := p.projectCollection.InsertOne(context.TODO(), project)
	if err != nil{
		return err
	}
	return nil
}

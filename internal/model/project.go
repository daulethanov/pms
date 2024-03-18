package model

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID            primitive.ObjectID   `bson:"_id"`
	UserID        primitive.ObjectID   `bson:"user_id"`
	RoomID        primitive.ObjectID   `bson:"room_id"`               
	Name          string               `bson:"name"`
	Category      string               `bson:"category"`
	Description   string               `bosn:"description"`
	General       bool                 `bson:"general"`
	Logo          string               `bson:"logo"`
	CretedAt      time.Time            `bson:"created_at"`
}



type TaskImportance struct{
	ID             primitive.ObjectID   `bson:"_id"`
	ProjectID      primitive.ObjectID   `bson:"project_id"`
	Level          string               `bson:"level"`
	TaskIDs        []primitive.ObjectID `bson:"task_ids"`
}


type Task struct {
	ID            primitive.ObjectID   `bson:"_id"`
	Title         string               `bson:"title"`
	Description   string               `bson:"description"`
	CretedAt      time.Time            `bson:"created_at"`
	UpdatedAt     time.Time            `bson:"updated_at"`
	LevelStage    string               `bson:"stage_level"`
	Image         []string             `bson:"image"`
	File          []string             `bson:"file"`
}


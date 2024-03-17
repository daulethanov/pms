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

type TaskLevel struct {
	ID            primitive.ObjectID   `bson:"_id"`
	TaskID        primitive.ObjectID   `bson:"task_id"`
	Name          string               `name`
}

type Task struct {
	ID            primitive.ObjectID   `bson:"_id"`
	ProjectID     primitive.ObjectID   `bson:"project_id"`
	Title         string               `bson:"title"`
	Description   string               `bson:description`
	CretedAt      time.Time            `bson:"created_at"`
	UpdatedAt     time.Time            `bson:"updated_at"`
}


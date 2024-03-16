package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID            primitive.ObjectID   `bson:"_id"`
	RoomID        primitive.ObjectID   `bson:"room_id"`               
	Name          string               `bson:"name"`
	General       bool                 `bson:"general"`
	CretedAt      time.Time            `bson:"created_at"`
}



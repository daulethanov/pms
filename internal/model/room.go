package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/google/uuid"
)

type Room struct {
	ID            primitive.ObjectID   `bson:"_id"`
	UsersID       []primitive.ObjectID `bson:"users"`
	InviteLink    string               `bson:"invite_link"`
}

func (r *Room) CreateInviteLink() string {
	uuid := uuid.New()
	return uuid.String()
}

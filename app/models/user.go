package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email"`
	Password string             `bson:"password,omitempty"`

	CreatedAt time.Time `bson:"createdAt"`
	DeletedAt time.Time `bson:"deletedAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}

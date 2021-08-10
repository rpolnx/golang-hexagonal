package user

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty" validate:"empty=false"`
	Age       int                `json:"age,omitempty" bson:"age,omitempty"`
	Birthday  string             `json:"birthday,omitempty" bson:"birthday,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
}

type UserDTO struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty" validate:"empty=false"`
	Age       int                `json:"age,omitempty" bson:"age,omitempty"`
	Birthday  string             `json:"birthday,omitempty" bson:"birthday,omitempty"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
}

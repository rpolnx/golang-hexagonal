package user

import "time"

type User struct {
	Name      string    `json:"name,omitempty" bson:"name,omitempty" validate:"empty=false"`
	Age       int       `json:"age,omitempty" bson:"age,omitempty"`
	Birthday  string    `json:"birthday,omitempty" bson:"birthday,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
}

type UserDTO struct {
	Name      string    `json:"name,omitempty" bson:"name,omitempty" validate:"empty=false"`
	Age       int       `json:"age,omitempty" bson:"age,omitempty"`
	Birthday  string    `json:"birthday,omitempty" bson:"birthday,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
}

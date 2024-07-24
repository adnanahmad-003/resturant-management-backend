package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       *string            `json:"name" validate:"required,min=2,max=100"`
	Price      *float64           `json:"price" validate:"required"`
	Food_image *string            `json:"food_image" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	Food_id    string             `json:"food_id"`
	Menu_id    *string            `json:"menu_id" validate:"required"`
}

//*string represent that if the Menu_id is not set it will be nil
//*string can be nil that is value not set yet it can be empty that means value is set to empty
//where as string can not be nil it can be only empty

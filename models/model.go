package models //name of the folder

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VideoLibrary struct {
	VideoID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Video   string             `json:"video,omitempty"`
	Watched bool               `json:"watech,omitempty"`
}

func main() {

}

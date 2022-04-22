package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Usuario modelo */
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre"`
	Apellidos       string             `bson:"apellidos" json:"apellidos"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password"`
	Avatar          string             `bson:"avatar" json:"avatar"`
	Banner          string             `bson:"banner" json:"banner"`
	Biografia       string             `bson:"biografia" json:"biografia"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion"`
	SitioWeb        string             `bson:"sitioweb" json:"sitioweb"`
}

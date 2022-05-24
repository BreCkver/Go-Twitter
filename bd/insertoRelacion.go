package bd

import (
	"context"
	"time"

	"github.com/BreCkver/Go-Twitter/models"
)

/*InsertoRelacion graba la relacion en la BD*/
func InsertoRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}

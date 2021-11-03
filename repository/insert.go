package repository

import (
	"context"
	"time"

	"github.com/nicolascancino/web-service-go/config"
	"github.com/nicolascancino/web-service-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertRegister(user models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := config.MongoConnection.Database("Name")
	col := db.Collection("usuarios")

	user.Password, _ = config.EncryptPassword(user.Password)

	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}

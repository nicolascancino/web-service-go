package repository

import (
	"context"
	"time"

	"github.com/nicolascancino/web-service-go/config"
	"github.com/nicolascancino/web-service-go/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckIfExistUser(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := config.MongoConnection.Database("Database Name")
	col := db.Collection("usuarios")

	condition := bson.M{"email": email}

	var result models.Usuario

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
